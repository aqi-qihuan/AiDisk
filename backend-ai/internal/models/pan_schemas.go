package models

import "time"

// PanQueryRequest 网盘查询请求（对标 Python models/pan_schemas.py）
type PanQueryRequest struct {
	Query string `json:"query"`
}

// FileInfo 文件信息
type FileInfo struct {
	ID          int       `json:"id"`
	FileID      int       `json:"file_id"`
	FileName    string    `json:"file_name"`
	FileType    string    `json:"file_type"`
	FileSuffix  string    `json:"file_suffix"`
	FileSize    int64     `json:"file_size"`
	GmtCreate   time.Time `json:"gmt_create"`
	GmtModified time.Time `json:"gmt_modified"`
}

// StorageInfo 存储空间信息
type StorageInfo struct {
	UsedSize        int64   `json:"used_size"`
	TotalSize       int64   `json:"total_size"`
	UsedPercentage  float64 `json:"used_percentage"`
}

// FileStatistics 文件统计信息
type FileStatistics struct {
	TotalFiles  int              `json:"total_files"`
	TotalSize   int64            `json:"total_size"`
	FileTypes   map[string]int   `json:"file_types"`
	RecentFiles []FileInfo       `json:"recent_files"`
}
