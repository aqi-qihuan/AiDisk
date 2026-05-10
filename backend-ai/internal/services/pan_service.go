package services

import (
	"context"
	"sync"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/core"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/models"
	"gorm.io/gorm"
)

// PanService 网盘查询服务（对标 Python agent/pan_agent.py 的数据库查询部分）
type PanService struct {
	db *gorm.DB
}

var (
	panServiceOnce sync.Once
	panService     *PanService
)

func GetPanService() *PanService {
	panServiceOnce.Do(func() {
		panService = &PanService{db: core.GetDB()}
	})
	return panService
}

// QueryStorage 查询存储空间使用情况
func (s *PanService) QueryStorage(ctx context.Context, accountID int) (*models.StorageInfo, error) {
	var result struct {
		UsedSize       int64
		TotalSize      int64
		Percentage     float64
	}

	err := s.db.WithContext(ctx).Raw(
		"SELECT COALESCE(SUM(file_size), 0) as used_size, MAX(s.total_size) as total_size, ROUND(COALESCE(SUM(file_size), 0) * 100.0 / NULLIF(MAX(s.total_size), 0), 2) as percentage FROM account_file af CROSS JOIN (SELECT total_size FROM storage WHERE account_id = ? LIMIT 1) s WHERE af.account_id = ? AND af.del = 0",
		accountID, accountID,
	).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &models.StorageInfo{
		UsedSize:       result.UsedSize,
		TotalSize:      result.TotalSize,
		UsedPercentage: result.Percentage,
	}, nil
}

// QueryFiles 查询文件列表
func (s *PanService) QueryFiles(ctx context.Context, accountID int, suffix string, limit int) ([]models.FileInfo, error) {
	var files []models.FileInfo

	query := s.db.WithContext(ctx).Table("account_file").
		Select("id, file_id, file_name, file_type, file_suffix, file_size, gmt_create, gmt_modified").
		Where("account_id = ? AND del = 0", accountID)

	if suffix != "" {
		query = query.Where("file_suffix = ?", suffix)
	}

	query = query.Order("gmt_create DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&files).Error
	return files, err
}

// QueryFileCount 查询文件总数
func (s *PanService) QueryFileCount(ctx context.Context, accountID int) (int64, error) {
	var count int64
	err := s.db.WithContext(ctx).Table("account_file").
		Where("account_id = ? AND del = 0", accountID).Count(&count).Error
	return count, err
}

// QueryTotalSize 查询文件总大小
func (s *PanService) QueryTotalSize(ctx context.Context, accountID int) (int64, error) {
	var totalSize int64
	err := s.db.WithContext(ctx).Table("account_file").
		Where("account_id = ? AND del = 0", accountID).
		Select("COALESCE(SUM(file_size), 0)").Scan(&totalSize).Error
	return totalSize, err
}

// QueryFileTypes 查询各文件类型的数量分布
func (s *PanService) QueryFileTypes(ctx context.Context, accountID int) (map[string]int, error) {
	type FileTypeCount struct {
		FileSuffix string
		Count      int
	}

	var results []FileTypeCount
	err := s.db.WithContext(ctx).Table("account_file").
		Select("file_suffix, COUNT(*) as count").
		Where("account_id = ? AND del = 0", accountID).
		Group("file_suffix").Find(&results).Error
	if err != nil {
		return nil, err
	}

	types := make(map[string]int)
	for _, r := range results {
		types[r.FileSuffix] = r.Count
	}
	return types, nil
}
