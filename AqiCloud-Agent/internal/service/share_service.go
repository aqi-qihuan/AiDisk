package service

import (
	"context"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"github.com/aqi/AqiCloud-Agent/internal/config"
	"github.com/aqi/AqiCloud-Agent/internal/model"
	"github.com/aqi/AqiCloud-Agent/internal/util"
	"gorm.io/gorm"
)

type ShareService struct{}

func NewShareService() *ShareService { return &ShareService{} }

func (s *ShareService) ListShare(ctx context.Context, accountID int64) ([]*model.ShareDTO, error) {
	db := config.GetDB()
	var shares []model.Share
	err := db.WithContext(ctx).Where("account_id = ?", accountID).
		Order("gmt_create DESC").Find(&shares).Error
	if err != nil {
		return nil, err
	}
	result := make([]*model.ShareDTO, 0, len(shares))
	for _, share := range shares {
		result = append(result, &model.ShareDTO{
			ID:           share.ID,
			ShareName:    share.ShareName,
			ShareType:    share.ShareType,
			ShareDayType: share.ShareDayType,
			ShareDay:     share.ShareDay,
			ShareEndTime: share.ShareEndTime.Format("2006-01-02 15:04:05"),
			ShareURL:     share.ShareURL,
			ShareCode:    share.ShareCode,
			ShareStatus:  share.ShareStatus,
			AccountID:    share.AccountID,
			GmtCreate:    share.GmtCreate.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}

func (s *ShareService) CreateShare(ctx context.Context, req *model.ShareCreateReq) (*model.ShareDTO, error) {
	db := config.GetDB()
	cfg := config.GetConfig()

	for _, fileID := range req.FileIDs {
		var count int64
		db.Model(&model.AccountFile{}).Where("id = ? AND account_id = ? AND del = 0", fileID, req.AccountID).Count(&count)
		if count == 0 {
			return nil, fmt.Errorf("文件不存在或不属于当前用户")
		}
	}

	shareID := util.NextID()
	var shareCode string
	var shareType = strings.ToUpper(req.ShareType)
	if shareType == "NEED_CODE" {
		shareCode = util.RandomString(6)
	}

	var shareEndTime time.Time
	var shareDay int
	switch req.ShareDayType {
	case 1:
		shareDay = 7
		shareEndTime = time.Now().AddDate(0, 0, 7)
	case 2:
		shareDay = 30
		shareEndTime = time.Now().AddDate(0, 0, 30)
	default:
		shareDay = 0
		shareEndTime = time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	}

	shareURL := fmt.Sprintf("%s/share/%d", cfg.FrontendBaseURL, shareID)

	share := model.Share{
		ID:           shareID,
		ShareName:    req.ShareName,
		ShareType:    shareType,
		ShareDayType: req.ShareDayType,
		ShareDay:     shareDay,
		ShareEndTime: shareEndTime,
		ShareURL:     shareURL,
		ShareCode:    shareCode,
		ShareStatus:  "USED",
		AccountID:    req.AccountID,
	}
	now := time.Now()
	share.GmtCreate = now
	share.GmtModified = now
	if err := db.WithContext(ctx).Create(&share).Error; err != nil {
		return nil, err
	}

	for _, fileID := range req.FileIDs {
		sf := model.ShareFile{
			ID:            util.NextID(),
			ShareID:       shareID,
			AccountFileID: fileID,
			AccountID:     req.AccountID,
		}
		sf.GmtCreate = now
		sf.GmtModified = now
		db.WithContext(ctx).Create(&sf)
	}

	return &model.ShareDTO{
		ID:           share.ID,
		ShareName:    share.ShareName,
		ShareType:    share.ShareType,
		ShareDayType: share.ShareDayType,
		ShareDay:     share.ShareDay,
		ShareEndTime: share.ShareEndTime.Format("2006-01-02 15:04:05"),
		ShareURL:     share.ShareURL,
		ShareCode:    share.ShareCode,
		ShareStatus:  share.ShareStatus,
		AccountID:    share.AccountID,
		GmtCreate:    share.GmtCreate.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *ShareService) CancelShare(ctx context.Context, req *model.ShareCancelReq) error {
	db := config.GetDB()

	var shares []model.Share
	db.WithContext(ctx).Where("id IN ? AND account_id = ?", req.ShareIDs, req.AccountID).Find(&shares)
	if len(shares) == 0 {
		return fmt.Errorf("分享不存在")
	}

	// 删除分享详情（share_file 关联记录）
	db.WithContext(ctx).Where("share_id IN ?", req.ShareIDs).Delete(&model.ShareFile{})

	// 删除分享记录
	return db.WithContext(ctx).Where("id IN ?", req.ShareIDs).Delete(&model.Share{}).Error
}

func (s *ShareService) VisitShare(ctx context.Context, shareID int64) (*model.ShareSimpleDTO, error) {
	db := config.GetDB()

	var share model.Share
	if err := db.WithContext(ctx).Where("id = ?", shareID).First(&share).Error; err != nil {
		return nil, fmt.Errorf("分享不存在")
	}

	if share.ShareStatus == "CANCELED" {
		return nil, fmt.Errorf("分享已取消")
	}
	if share.ShareStatus == "EXPIRED" || (share.ShareEndTime.Before(time.Now()) && share.ShareDayType != 0) {
		return nil, fmt.Errorf("分享已过期")
	}

	var owner model.Account
	db.WithContext(ctx).Where("id = ?", share.AccountID).First(&owner)

	var fileCount int64
	db.Model(&model.ShareFile{}).Where("share_id = ?", shareID).Count(&fileCount)

	dto := &model.ShareSimpleDTO{
		ID:           share.ID,
		ShareName:    share.ShareName,
		ShareType:    share.ShareType,
		ShareDayType: share.ShareDayType,
		ShareDay:     share.ShareDay,
		ShareEndTime: share.ShareEndTime.Format("2006-01-02 15:04:05"),
		ShareURL:     share.ShareURL,
		ShareAccountDTO: &model.ShareAccountDTO{
			ID:        owner.ID,
			Username:  owner.Username,
			AvatarURL: owner.AvatarURL,
		},
		FileCount: int(fileCount),
	}

	if strings.EqualFold(share.ShareType, "NO_CODE") {
		dto.ShareToken = util.GenerateShareToken(shareID)
	}

	return dto, nil
}

func (s *ShareService) CheckShareCode(ctx context.Context, req *model.ShareCheckReq) (string, error) {
	db := config.GetDB()

	var share model.Share
	if err := db.WithContext(ctx).Where("id = ? AND share_code = ? AND share_status = ?", req.ShareID, req.ShareCode, "USED").First(&share).Error; err != nil {
		return "", fmt.Errorf("提取码错误或分享不存在")
	}

	// 判断状态和是否过期
	if share.ShareStatus == "EXPIRED" || share.ShareEndTime.Before(time.Now()) {
		return "", fmt.Errorf("分享已过期")
	}

	return util.GenerateShareToken(share.ID), nil
}

func (s *ShareService) GetShareDetail(ctx context.Context, shareID int64) (*model.ShareDetailDTO, error) {
	db := config.GetDB()

	var share model.Share
	if err := db.WithContext(ctx).Where("id = ?", shareID).First(&share).Error; err != nil {
		return nil, fmt.Errorf("分享不存在")
	}

	if share.ShareStatus == "CANCELED" {
		return nil, fmt.Errorf("分享已取消")
	}
	if share.ShareStatus == "EXPIRED" || (share.ShareEndTime.Before(time.Now()) && share.ShareDayType != 0) {
		return nil, fmt.Errorf("分享已过期")
	}

	var owner model.Account
	db.WithContext(ctx).Where("id = ?", share.AccountID).First(&owner)

	var shareFiles []model.ShareFile
	db.WithContext(ctx).Where("share_id = ?", shareID).Find(&shareFiles)

	fileIDs := make([]int64, len(shareFiles))
	for i, sf := range shareFiles {
		fileIDs[i] = sf.AccountFileID
	}

	var accountFiles []model.AccountFile
	if len(fileIDs) > 0 {
		db.WithContext(ctx).Where("id IN ?", fileIDs).Find(&accountFiles)
	}

	return &model.ShareDetailDTO{
		ID:           share.ID,
		ShareName:    share.ShareName,
		ShareType:    share.ShareType,
		ShareDayType: share.ShareDayType,
		ShareDay:     share.ShareDay,
		ShareEndTime: share.ShareEndTime.Format("2006-01-02 15:04:05"),
		ShareURL:     share.ShareURL,
		ShareAccountDTO: &model.ShareAccountDTO{
			ID:        owner.ID,
			Username:  owner.Username,
			AvatarURL: owner.AvatarURL,
		},
		FileDTOList: toAccountFileDTOs(accountFiles),
	}, nil
}

func (s *ShareService) ListShareFiles(ctx context.Context, shareID int64, parentID int64) ([]*model.AccountFileDTO, error) {
	db := config.GetDB()

	var share model.Share
	if err := db.WithContext(ctx).Where("id = ?", shareID).First(&share).Error; err != nil {
		return nil, fmt.Errorf("分享不存在")
	}

	// 获取分享的文件列表
	var shareFileDOS []model.ShareFile
	db.WithContext(ctx).Where("share_id = ?", shareID).Find(&shareFileDOS)

	shareFileIDs := make([]int64, len(shareFileDOS))
	for i, sf := range shareFileDOS {
		shareFileIDs[i] = sf.AccountFileID
	}

	// 获取分享文件的完整列表（包括递归子文件）
	var shareAccountFiles []model.AccountFile
	if len(shareFileIDs) > 0 {
		db.WithContext(ctx).Where("id IN ? AND del = 0", shareFileIDs).Find(&shareAccountFiles)
	}

	// 递归收集所有子文件
	var allShareFiles []model.AccountFile
	s.findAllShareFilesWithRecur(ctx, db, shareAccountFiles, &allShareFiles, false)

	// 构建分享文件ID集合
	allShareFileIDSet := make(map[int64]bool, len(allShareFiles))
	for _, f := range allShareFiles {
		allShareFileIDSet[f.ID] = true
	}

	// 校验 parentId 是否在分享范围内（parentID 为 0 表示根目录，不需要校验）
	if parentID != 0 && !allShareFileIDSet[parentID] {
		return nil, fmt.Errorf("目标文件夹不在分享范围内")
	}

	// 按 parentId 分组，返回指定父文件夹下的子文件（确保只在分享范围内返回）
	var files []model.AccountFile
	if parentID == 0 {
		// 根目录：只返回分享的顶级文件
		db.WithContext(ctx).Where("id IN ? AND del = 0", shareFileIDs).Order("is_dir DESC").Find(&files)
	} else {
		// 非根目录：返回该文件夹下的子文件，但只保留在分享范围内的
		db.WithContext(ctx).Where("parent_id = ? AND del = 0", parentID).Order("is_dir DESC").Find(&files)
		filtered := make([]model.AccountFile, 0, len(files))
		for _, f := range files {
			if allShareFileIDSet[f.ID] {
				filtered = append(filtered, f)
			}
		}
		files = filtered
	}

	return toAccountFileDTOs(files), nil
}

func (s *ShareService) TransferToOwn(ctx context.Context, shareID int64, fileIDs []int64, targetParentID int64, accountID int64) error {
	db := config.GetDB()

	// 验证目标文件夹是否属于当前用户（targetParentID=0 表示转存到根目录）
	if targetParentID != 0 {
		var targetFolder model.AccountFile
		if err := db.WithContext(ctx).Where("id = ? AND account_id = ? AND is_dir = 1 AND del = 0", targetParentID, accountID).First(&targetFolder).Error; err != nil {
			return fmt.Errorf("目标文件夹不存在或不属于当前用户")
		}
	}

	// 获取分享中的文件列表
	var shareFileDOS []model.ShareFile
	db.WithContext(ctx).Where("share_id = ?", shareID).Find(&shareFileDOS)
	shareFileIDs := make([]int64, len(shareFileDOS))
	for i, sf := range shareFileDOS {
		shareFileIDs[i] = sf.AccountFileID
	}

	// 获取分享文件的完整列表（包括递归子文件）
	var shareAccountFiles []model.AccountFile
	if len(shareFileIDs) > 0 {
		db.WithContext(ctx).Where("id IN ? AND del = 0", shareFileIDs).Find(&shareAccountFiles)
	}
	var allShareFiles []model.AccountFile
	s.findAllShareFilesWithRecur(ctx, db, shareAccountFiles, &allShareFiles, false)

	// 构建分享文件ID集合
	allShareFileIDSet := make(map[int64]bool, len(allShareFiles))
	for _, f := range allShareFiles {
		allShareFileIDSet[f.ID] = true
	}

	// 验证要转存的文件是否都在分享范围内
	for _, fileID := range fileIDs {
		if !allShareFileIDSet[fileID] {
			return fmt.Errorf("文件不在分享范围内")
		}
	}

	// 获取要转存的文件实体
	var files []model.AccountFile
	db.WithContext(ctx).Where("id IN ? AND del = 0", fileIDs).Find(&files)

	// 递归生成所有需要转存的文件（包括子文件），类似 Java 的 findBatchCopyFileWithRecur
	batchCopyFiles := make([]model.AccountFile, 0)
	for _, f := range files {
		s.copyShareFileWithRecur(f, targetParentID, accountID, &batchCopyFiles)
	}

	// 计算总存储空间（只统计文件，不统计文件夹）
	var totalSize int64
	for _, f := range batchCopyFiles {
		if f.IsDir == 0 {
			totalSize += f.FileSize
		}
	}

	// 检查存储空间是否足够
	var storage model.Storage
	if err := db.WithContext(ctx).Where("account_id = ?", accountID).First(&storage).Error; err != nil {
		return err
	}
	if storage.UsedSize+totalSize > storage.TotalSize {
		return fmt.Errorf("存储空间不足")
	}

	// 批量插入新记录
	for i := range batchCopyFiles {
		db.WithContext(ctx).Create(&batchCopyFiles[i])
	}

	return db.WithContext(ctx).Model(&model.Storage{}).
		Where("account_id = ?", accountID).
		UpdateColumn("used_size", gorm.Expr("used_size + ?", totalSize)).Error
}

// copyShareFileWithRecur 递归复制文件及其子文件，生成新ID，处理文件名重复
func (s *ShareService) copyShareFileWithRecur(src model.AccountFile, targetParentID int64, accountID int64, result *[]model.AccountFile) {
	oldID := src.ID
	newID := util.NextID()

	src.ID = newID
	src.AccountID = accountID
	src.ParentID = targetParentID
	src.GmtCreate = time.Now()
	src.GmtModified = time.Now()

	// 处理文件名重复（对标 Java 的 processFileNameDuplicate）
	s.processFileNameDuplicate(&src)

	*result = append(*result, src)

	// 如果是文件夹，递归处理子文件
	if src.IsDir == 1 {
		db := config.GetDB()
		var children []model.AccountFile
		db.Where("parent_id = ? AND del = 0", oldID).Find(&children)
		for _, child := range children {
			s.copyShareFileWithRecur(child, newID, accountID, result)
		}
	}
}

// processFileNameDuplicate 处理转存时的文件名重复
func (s *ShareService) processFileNameDuplicate(f *model.AccountFile) {
	db := config.GetDB()
	var count int64
	db.Model(&model.AccountFile{}).
		Where("account_id = ? AND parent_id = ? AND is_dir = ? AND file_name = ? AND del = 0",
			f.AccountID, f.ParentID, f.IsDir, f.FileName).Count(&count)
	if count > 0 {
		if f.IsDir == 1 {
			// 文件夹重复：加时间戳
			f.FileName = fmt.Sprintf("%s_%d", f.FileName, time.Now().UnixMilli())
		} else {
			// 文件重复：在扩展名前加时间戳
			ext := filepath.Ext(f.FileName)
			nameWithoutExt := strings.TrimSuffix(f.FileName, ext)
			if ext != "" {
				f.FileName = fmt.Sprintf("%s_%d%s", nameWithoutExt, time.Now().UnixMilli(), ext)
			} else {
				f.FileName = fmt.Sprintf("%s_%d", f.FileName, time.Now().UnixMilli())
			}
		}
	}
}

// findAllShareFilesWithRecur 递归收集分享文件中的所有子文件
func (s *ShareService) findAllShareFilesWithRecur(ctx context.Context, db *gorm.DB, input []model.AccountFile, result *[]model.AccountFile, onlyFolder bool) {
	for _, f := range input {
		if f.IsDir == 1 {
			var children []model.AccountFile
			db.WithContext(ctx).Where("parent_id = ? AND del = 0", f.ID).Find(&children)
			s.findAllShareFilesWithRecur(ctx, db, children, result, onlyFolder)
		}
		if !onlyFolder || f.IsDir == 1 {
			*result = append(*result, f)
		}
	}
}

var _ = rand.Int
