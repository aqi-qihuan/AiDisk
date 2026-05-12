package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aqi/AqiCloud-Agent/internal/config"
	"github.com/aqi/AqiCloud-Agent/internal/model"
	"github.com/aqi/AqiCloud-Agent/internal/util"
	"gorm.io/gorm"
)

type RecycleService struct{}

func NewRecycleService() *RecycleService { return &RecycleService{} }

func (s *RecycleService) ListRecycle(ctx context.Context, accountID int64) ([]*model.AccountFileDTO, error) {
	db := config.GetDB()

	var deletedFiles []model.AccountFile
	db.WithContext(ctx).
		Where("account_id = ? AND del = 1", accountID).
		Order("del_time DESC").Find(&deletedFiles)

	var parentIDs []int64
	for _, f := range deletedFiles {
		parentIDs = append(parentIDs, f.ParentID)
	}

	var parentDeletedCount int64
	if len(parentIDs) > 0 {
		db.Model(&model.AccountFile{}).
			Where("id IN ? AND account_id = ? AND del = 1", parentIDs, accountID).
			Count(&parentDeletedCount)
	}

	if parentDeletedCount == 0 {
		return toAccountFileDTOs(deletedFiles), nil
	}

	var result []model.AccountFile
	for _, f := range deletedFiles {
		if s.isAncestorDeleted(f, deletedFiles) {
			continue
		}
		result = append(result, f)
	}
	return toAccountFileDTOs(result), nil
}

func (s *RecycleService) isAncestorDeleted(file model.AccountFile, allDeleted []model.AccountFile) bool {
	parentID := file.ParentID
	for parentID != 0 {
		for _, d := range allDeleted {
			if d.ID == parentID {
				return true
			}
		}
		var parent model.AccountFile
		if err := config.GetDB().Model(&model.AccountFile{}).Where("id = ?", parentID).First(&parent).Error; err != nil {
			break
		}
		parentID = parent.ParentID
	}
	return false
}

func (s *RecycleService) PermanentDelete(ctx context.Context, req *model.RecycleDelReq) error {
	db := config.GetDB()

	for _, fileID := range req.FileIDs {
		s.deleteRecursive(ctx, db, fileID, req.AccountID)
	}
	return nil
}

func (s *RecycleService) deleteRecursive(ctx context.Context, db *gorm.DB, fileID int64, accountID int64) {
	var children []model.AccountFile
	db.WithContext(ctx).Where("parent_id = ? AND account_id = ?", fileID, accountID).Find(&children)
	for _, c := range children {
		s.deleteRecursive(ctx, db, c.ID, accountID)
	}

	var file model.AccountFile
	if err := db.WithContext(ctx).Where("id = ? AND account_id = ?", fileID, accountID).First(&file).Error; err != nil {
		return
	}

	if file.IsDir == 0 && file.FileID != nil {
		var fileRecord model.File
		db.WithContext(ctx).Where("id = ?", file.FileID).First(&fileRecord)
		if fileRecord.ID != 0 {
			db.WithContext(ctx).Delete(&fileRecord)
			getStoreEngine().Delete(ctx, fileRecord.ObjectKey)
		}
	}

	db.WithContext(ctx).Delete(&file)
}

func (s *RecycleService) Restore(ctx context.Context, req *model.RecycleRestoreReq) error {
	db := config.GetDB()

	for _, fileID := range req.FileIDs {
		s.restoreRecursive(ctx, db, fileID, req.AccountID)
	}
	return nil
}

func (s *RecycleService) restoreRecursive(ctx context.Context, db *gorm.DB, fileID int64, accountID int64) {
	var children []model.AccountFile
	db.WithContext(ctx).Where("parent_id = ? AND account_id = ?", fileID, accountID).Find(&children)
	for _, c := range children {
		s.restoreRecursive(ctx, db, c.ID, accountID)
	}

	var file model.AccountFile
	if err := db.WithContext(ctx).Where("id = ? AND account_id = ?", fileID, accountID).First(&file).Error; err != nil {
		return
	}

	newName := file.FileName
	var count int64
	db.Model(&model.AccountFile{}).
		Where("account_id = ? AND parent_id = ? AND file_name = ? AND id != ? AND del = 0",
			accountID, file.ParentID, newName, fileID).Count(&count)
	if count > 0 {
		ext := ""
		for i := len(newName) - 1; i >= 0; i-- {
			if newName[i] == '.' {
				ext = newName[i:]
				break
			}
		}
		if ext != "" {
			newName = newName[:len(newName)-len(ext)] + fmt.Sprintf("_%d", time.Now().UnixMilli()) + ext
		} else {
			newName = fmt.Sprintf("%s_%d", newName, time.Now().UnixMilli())
		}
	}

	db.WithContext(ctx).Model(&model.AccountFile{}).Where("id = ?", fileID).Updates(map[string]interface{}{
		"del":       0,
		"del_time":  nil,
		"file_name": newName,
	})
}

func (s *RecycleService) GetDeletedFileSize(ctx context.Context, accountID int64) int64 {
	var total int64
	config.GetDB().WithContext(ctx).
		Model(&model.AccountFile{}).
		Where("account_id = ? AND del = 1 AND is_dir = 0").
		Select("COALESCE(SUM(file_size), 0)").Scan(&total)
	return total
}

var _ = util.NextID
