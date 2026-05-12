package service

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/aqi/AqiCloud-Agent/internal/config"
	"github.com/aqi/AqiCloud-Agent/internal/model"
	"github.com/aqi/AqiCloud-Agent/internal/util"
	"gorm.io/gorm"
)

type FileService struct{}

func NewFileService() *FileService { return &FileService{} }

func (s *FileService) ListFiles(ctx context.Context, accountID, parentID int64) ([]*model.AccountFileDTO, error) {
	var files []model.AccountFile
	err := config.GetDB().WithContext(ctx).
		Where("account_id = ? AND parent_id = ? AND del = 0", accountID, parentID).
		Order("is_dir DESC, gmt_create DESC").
		Find(&files).Error
	if err != nil {
		return nil, err
	}
	return toAccountFileDTOs(files), nil
}

func (s *FileService) CreateFolder(ctx context.Context, req *model.FolderCreateReq) error {
	db := config.GetDB()

	var parent model.AccountFile
	if req.ParentID != 0 {
		if err := db.WithContext(ctx).Where("id = ? AND account_id = ? AND is_dir = 1 AND del = 0", req.ParentID, req.AccountID).First(&parent).Error; err != nil {
			return fmt.Errorf("父文件夹不存在")
		}
	}

	var count int64
	db.Model(&model.AccountFile{}).
		Where("account_id = ? AND parent_id = ? AND file_name = ? AND is_dir = 1 AND del = 0",
			req.AccountID, req.ParentID, req.FolderName).Count(&count)
	if count > 0 {
		return fmt.Errorf("文件夹名称重复")
	}

	folder := model.AccountFile{
		ID:        util.NextID(),
		AccountID: req.AccountID,
		IsDir:     1,
		ParentID:  req.ParentID,
		FileName:  req.FolderName,
		Del:       false,
	}
	now := time.Now()
	folder.GmtCreate = now
	folder.GmtModified = now
	return db.WithContext(ctx).Create(&folder).Error
}

func (s *FileService) RenameFile(ctx context.Context, req *model.FileUpdateReq) error {
	db := config.GetDB()
	var file model.AccountFile
	if err := db.WithContext(ctx).Where("id = ? AND account_id = ? AND del = 0", req.FileID, req.AccountID).First(&file).Error; err != nil {
		return fmt.Errorf("文件不存在")
	}
	if file.FileName == req.NewFileName {
		return fmt.Errorf("文件夹名称重复")
	}
	var count int64
	db.Model(&model.AccountFile{}).
		Where("account_id = ? AND parent_id = ? AND file_name = ? AND del = 0",
			req.AccountID, file.ParentID, req.NewFileName).Count(&count)
	if count > 0 {
		return fmt.Errorf("文件夹名称重复")
	}
	return db.WithContext(ctx).Model(&model.AccountFile{}).
		Where("id = ?", req.FileID).Update("file_name", req.NewFileName).Error
}

func (s *FileService) GetFolderTree(ctx context.Context, accountID int64) ([]*model.FolderTreeNodeDTO, error) {
	var folders []model.AccountFile
	err := config.GetDB().WithContext(ctx).
		Where("account_id = ? AND is_dir = 1 AND del = 0", accountID).
		Find(&folders).Error
	if err != nil {
		return nil, err
	}

	nodeMap := make(map[int64]*model.FolderTreeNodeDTO)
	var roots []*model.FolderTreeNodeDTO

	for _, f := range folders {
		node := &model.FolderTreeNodeDTO{
			ID:       f.ID,
			ParentID: f.ParentID,
			Label:    f.FileName,
			Children: make([]*model.FolderTreeNodeDTO, 0),
		}
		nodeMap[f.ID] = node
	}

	for _, node := range nodeMap {
		if node.ParentID == 0 {
			roots = append(roots, node)
		} else if parent, ok := nodeMap[node.ParentID]; ok {
			parent.Children = append(parent.Children, node)
		}
	}

	return roots, nil
}

func (s *FileService) UploadFile(ctx context.Context, accountID, parentID int64, fileName string, fileSize int64, file io.Reader) error {
	db := config.GetDB()

	var storage model.Storage
	if err := db.WithContext(ctx).Where("account_id = ?", accountID).First(&storage).Error; err != nil {
		return err
	}
	if storage.UsedSize+fileSize > storage.TotalSize {
		return fmt.Errorf("存储空间不足")
	}

	ext := strings.TrimPrefix(filepath.Ext(fileName), ".")
	fileType := util.DetectFileType(ext)
	objectKey := fmt.Sprintf("%d/%d/%d/%s.%s",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		util.NextID(), ext)

	if err := getStoreEngine().Upload(ctx, file, objectKey, fileSize); err != nil {
		return err
	}

	fileID := util.NextID()
	fileRecord := model.File{
		ID:          fileID,
		AccountID:   accountID,
		FileName:    fileName,
		FileSuffix:  ext,
		FileSize:    fileSize,
		ObjectKey:   objectKey,
		Identifier:  util.MD5(fileName + fmt.Sprintf("%d", time.Now().UnixNano())),
		Del:         false,
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
	}
	db.WithContext(ctx).Create(&fileRecord)

	accountFile := model.AccountFile{
		ID:          util.NextID(),
		AccountID:   accountID,
		IsDir:       0,
		ParentID:    parentID,
		FileID:      &fileID,
		FileName:    fileName,
		FileType:    fileType,
		FileSuffix:  ext,
		FileSize:    fileSize,
		Del:         false,
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
	}
	db.WithContext(ctx).Create(&accountFile)

	return db.WithContext(ctx).Model(&model.Storage{}).
		Where("account_id = ?", accountID).
		UpdateColumn("used_size", gorm.Expr("used_size + ?", fileSize)).Error
}

func (s *FileService) BatchMove(ctx context.Context, req *model.FileBatchReq) error {
	db := config.GetDB()
	return db.WithContext(ctx).Model(&model.AccountFile{}).
		Where("id IN ? AND account_id = ? AND del = 0", req.FileIDs, req.AccountID).
		Update("parent_id", req.TargetParentID).Error
}

func (s *FileService) BatchCopy(ctx context.Context, req *model.FileBatchReq) error {
	db := config.GetDB()

	var files []model.AccountFile
	db.WithContext(ctx).Where("id IN ? AND account_id = ? AND del = 0", req.FileIDs, req.AccountID).Find(&files)

	var totalSize int64
	for _, f := range files {
		totalSize += f.FileSize
	}

	var storage model.Storage
	db.WithContext(ctx).Where("account_id = ?", req.AccountID).First(&storage)
	if storage.UsedSize+totalSize > storage.TotalSize {
		return fmt.Errorf("存储空间不足")
	}

	for _, f := range files {
		newID := util.NextID()
		newFile := f
		newFile.ID = newID
		newFile.AccountID = req.AccountID
		newFile.ParentID = req.TargetParentID
		newFile.GmtCreate = time.Now()
		newFile.GmtModified = time.Now()

		if f.IsDir == 1 {
			db.WithContext(ctx).Create(&newFile)
			s.copyChildren(ctx, f.ID, newID, req.AccountID)
		} else {
			db.WithContext(ctx).Create(&newFile)
		}
	}

	return db.WithContext(ctx).Model(&model.Storage{}).
		Where("account_id = ?", req.AccountID).
		UpdateColumn("used_size", gorm.Expr("used_size + ?", totalSize)).Error
}

func (s *FileService) copyChildren(ctx context.Context, oldParentID, newParentID, accountID int64) {
	db := config.GetDB()
	var children []model.AccountFile
	db.WithContext(ctx).Where("parent_id = ? AND del = 0", oldParentID).Find(&children)
	for _, c := range children {
		newID := util.NextID()
		newChild := c
		newChild.ID = newID
		newChild.ParentID = newParentID
		newChild.GmtCreate = time.Now()
		db.WithContext(ctx).Create(&newChild)
		if c.IsDir == 1 {
			s.copyChildren(ctx, c.ID, newID, accountID)
		}
	}
}

func (s *FileService) BatchDelete(ctx context.Context, req *model.FileDelReq) error {
	db := config.GetDB()

	// 获取待删除的文件记录
	var files []model.AccountFile
	db.WithContext(ctx).Where("id IN ? AND account_id = ? AND del = 0", req.FileIDs, req.AccountID).Find(&files)

	// 递归获取所有文件（包括文件夹内的文件）
	allFiles := make([]model.AccountFile, 0, len(files))
	s.collectAllFiles(ctx, files, &allFiles)

	// 计算总文件大小（只统计文件，不统计文件夹）
	var totalSize int64
	for _, f := range allFiles {
		if f.IsDir == 0 {
			totalSize += f.FileSize
		}
	}

	// 更新存储空间
	if totalSize > 0 {
		db.WithContext(ctx).Model(&model.Storage{}).
			Where("account_id = ?", req.AccountID).
			UpdateColumn("used_size", gorm.Expr("used_size - ?", totalSize))
	}

	// 收集所有需要删除的ID
	allIDs := make([]int64, 0, len(allFiles))
	for _, f := range allFiles {
		allIDs = append(allIDs, f.ID)
	}

	now := time.Now()
	return db.WithContext(ctx).Model(&model.AccountFile{}).
		Where("id IN ? AND del = 0", allIDs).
		Updates(map[string]interface{}{
			"del":      1,
			"del_time": now,
		}).Error
}

// collectAllFiles 递归收集所有文件（包括文件夹内的子文件）
func (s *FileService) collectAllFiles(ctx context.Context, input []model.AccountFile, result *[]model.AccountFile) {
	db := config.GetDB()
	for _, f := range input {
		*result = append(*result, f)
		if f.IsDir == 1 {
			var children []model.AccountFile
			db.WithContext(ctx).Where("parent_id = ? AND del = 0", f.ID).Find(&children)
			if len(children) > 0 {
				s.collectAllFiles(ctx, children, result)
			}
		}
	}
}

func (s *FileService) SecondUpload(ctx context.Context, req *model.FileSecondUploadReq) (bool, error) {
	db := config.GetDB()

	var existing model.File
	err := db.WithContext(ctx).Where("identifier = ? AND del = 0", req.Identifier).First(&existing).Error
	if err == nil {
		// 检查空间是否足够
		var storage model.Storage
		if err := db.WithContext(ctx).Where("account_id = ?", req.AccountID).First(&storage).Error; err != nil {
			return false, err
		}
		if storage.UsedSize+existing.FileSize > storage.TotalSize {
			return false, fmt.Errorf("存储空间不足")
		}

		newName := req.FileName
		if ext := filepath.Ext(newName); ext != "" {
			nameWithoutExt := strings.TrimSuffix(newName, ext)
			newName = fmt.Sprintf("%s_%d%s", nameWithoutExt, time.Now().UnixMilli(), ext)
		} else {
			newName = fmt.Sprintf("%s_%d", newName, time.Now().UnixMilli())
		}

		accountFile := model.AccountFile{
			ID:          util.NextID(),
			AccountID:   req.AccountID,
			IsDir:       0,
			ParentID:    req.ParentID,
			FileID:      &existing.ID,
			FileName:    newName,
			FileType:    util.DetectFileType(existing.FileSuffix),
			FileSuffix:  existing.FileSuffix,
			FileSize:    existing.FileSize,
			Del:         false,
			GmtCreate:   time.Now(),
			GmtModified: time.Now(),
		}
		db.WithContext(ctx).Create(&accountFile)

		// 更新存储空间
		db.WithContext(ctx).Model(&model.Storage{}).
			Where("account_id = ?", req.AccountID).
			UpdateColumn("used_size", gorm.Expr("used_size + ?", existing.FileSize))

		return true, nil
	}

	return false, nil
}

func (s *FileService) SearchFiles(ctx context.Context, accountID int64, keyword string) ([]*model.AccountFileDTO, error) {
	var files []model.AccountFile
	err := config.GetDB().WithContext(ctx).
		Where("account_id = ? AND file_name LIKE ? AND del = 0 AND is_dir = 0",
			accountID, "%"+keyword+"%").
		Limit(30).Find(&files).Error
	if err != nil {
		return nil, err
	}
	return toAccountFileDTOs(files), nil
}

func (s *FileService) BatchDownloadURL(ctx context.Context, req *model.FileDownloadReq) ([]*model.FileDownloadDTO, error) {
	var files []model.AccountFile
	config.GetDB().WithContext(ctx).
		Where("id IN ? AND account_id = ? AND is_dir = 0 AND del = 0", req.FileIDs, req.AccountID).
		Find(&files)

	var results []*model.FileDownloadDTO
	for _, f := range files {
		var file model.File
		if err := config.GetDB().WithContext(ctx).Where("id = ?", f.FileID).First(&file).Error; err != nil {
			continue
		}
		url, err := getStoreEngine().GeneratePresignedGetURL(ctx, file.ObjectKey, 5*time.Minute)
		if err != nil {
			continue
		}
		results = append(results, &model.FileDownloadDTO{
			FileName:    f.FileName,
			DownloadURL: url,
		})
	}
	return results, nil
}

func (s *FileService) DownloadFile(ctx context.Context, fileID int64) (io.ReadCloser, string, int64, error) {
	db := config.GetDB()
	var af model.AccountFile
	if err := db.WithContext(ctx).Where("id = ? AND del = 0 AND is_dir = 0", fileID).First(&af).Error; err != nil {
		return nil, "", 0, fmt.Errorf("文件不存在")
	}

	if af.FileID == nil {
		return nil, "", 0, fmt.Errorf("文件存储记录不存在")
	}

	var file model.File
	if err := db.WithContext(ctx).Where("id = ?", *af.FileID).First(&file).Error; err != nil {
		return nil, "", 0, fmt.Errorf("文件存储记录不存在")
	}

	reader, size, err := getStoreEngine().Download(ctx, file.ObjectKey)
	return reader, af.FileName, size, err
}

// PreviewFile 获取文件预览信息（流 + 文件名 + 大小 + MIME类型）
func (s *FileService) PreviewFile(ctx context.Context, fileID int64) (io.ReadCloser, string, int64, string, error) {
	db := config.GetDB()
	var af model.AccountFile
	if err := db.WithContext(ctx).Where("id = ? AND del = 0 AND is_dir = 0", fileID).First(&af).Error; err != nil {
		return nil, "", 0, "", fmt.Errorf("文件不存在")
	}

	if af.FileID == nil {
		return nil, "", 0, "", fmt.Errorf("文件存储记录不存在")
	}

	var file model.File
	if err := db.WithContext(ctx).Where("id = ?", *af.FileID).First(&file).Error; err != nil {
		return nil, "", 0, "", fmt.Errorf("文件存储记录不存在")
	}

	reader, size, err := getStoreEngine().Download(ctx, file.ObjectKey)
	if err != nil {
		return nil, "", 0, "", err
	}
	mimeType := util.GetMIMEType(file.FileSuffix)
	return reader, af.FileName, size, mimeType, nil
}

// downloadEntry 预检后的可下载文件信息
type downloadEntry struct {
	objectKey  string
	fileName   string
	folderPath string
}

func (s *FileService) BatchDownloadZip(ctx context.Context, fileIDs []int64, accountID int64) (*bytes.Buffer, error) {
	db := config.GetDB()

	if len(fileIDs) == 0 {
		return nil, fmt.Errorf("缺少文件ID")
	}

	var items []model.AccountFile
	if err := db.WithContext(ctx).Where("id IN ? AND account_id = ? AND del = 0", fileIDs, accountID).Find(&items).Error; err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("文件不存在或无权访问 (查询IDs: %v, accountID: %d)", fileIDs, accountID)
	}

	// 预检：收集所有可下载的文件
	entries := make([]downloadEntry, 0)
	if err := s.collectDownloadableFiles(ctx, db, items, accountID, "", &entries); err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("选中的项中没有可下载的文件")
	}

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for _, e := range entries {
		reader, _, err := getStoreEngine().Download(ctx, e.objectKey)
		if err != nil {
			w.Close()
			return nil, fmt.Errorf("下载文件「%s」失败: %w", e.fileName, err)
		}
		fw, err := w.Create(e.folderPath + e.fileName)
		if err != nil {
			reader.Close()
			w.Close()
			return nil, fmt.Errorf("创建 ZIP 条目失败: %w", err)
		}
		_, err = io.Copy(fw, reader)
		reader.Close()
		if err != nil {
			w.Close()
			return nil, fmt.Errorf("写入 ZIP 条目失败: %w", err)
		}
	}
	w.Close()
	return buf, nil
}

// collectDownloadableFiles 递归收集所有可下载的文件，批量查询 file 记录
func (s *FileService) collectDownloadableFiles(ctx context.Context, db *gorm.DB, items []model.AccountFile, accountID int64, folderPath string, result *[]downloadEntry) error {
	// 收集所有非目录文件的 FileID，批量查询 file 记录
	fileIDPtrs := make([]*int64, 0)
	for _, item := range items {
		if item.IsDir == 0 && item.FileID != nil {
			fileIDPtrs = append(fileIDPtrs, item.FileID)
		}
	}

	fileMap := make(map[int64]*model.File)
	if len(fileIDPtrs) > 0 {
		var files []model.File
		if err := db.WithContext(ctx).Where("id IN ?", fileIDPtrs).Find(&files).Error; err != nil {
			return err
		}
		for i := range files {
			fileMap[files[i].ID] = &files[i]
		}
	}

	for _, item := range items {
		if item.IsDir == 1 {
			// 递归收集子文件，加上 account_id 过滤
			var children []model.AccountFile
			if err := db.WithContext(ctx).Where("parent_id = ? AND account_id = ? AND del = 0", item.ID, accountID).Find(&children).Error; err != nil {
				return err
			}
			if err := s.collectDownloadableFiles(ctx, db, children, accountID, folderPath+item.FileName+"/", result); err != nil {
				return err
			}
		} else if item.FileID != nil {
			fileRecord, ok := fileMap[*item.FileID]
			if !ok {
				continue
			}
			*result = append(*result, downloadEntry{
				objectKey:  fileRecord.ObjectKey,
				fileName:   item.FileName,
				folderPath: folderPath,
			})
		}
	}
	return nil
}

func toAccountFileDTO(files []model.AccountFile) []*model.AccountFileDTO {
	result := make([]*model.AccountFileDTO, 0, len(files))
	for _, f := range files {
		dto := &model.AccountFileDTO{
			ID:          f.ID,
			AccountID:   f.AccountID,
			IsDir:       f.IsDir,
			ParentID:    f.ParentID,
			FileID:      f.FileID,
			FileName:    f.FileName,
			FileType:    f.FileType,
			FileSuffix:  f.FileSuffix,
			FileSize:    f.FileSize,
			Del:         f.Del,
			GmtCreate:   f.GmtCreate.Format("2006-01-02 15:04:05"),
			GmtModified: f.GmtModified.Format("2006-01-02 15:04:05"),
		}
		if f.DelTime != nil {
			dto.DelTime = f.DelTime.Format("2006-01-02 15:04:05")
		}
		result = append(result, dto)
	}
	return result
}

func toAccountFileDTOs(files []model.AccountFile) []*model.AccountFileDTO {
	return toAccountFileDTO(files)
}
