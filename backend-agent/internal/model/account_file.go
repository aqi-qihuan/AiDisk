package model

import "time"

type AccountFile struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	AccountID  int64     `gorm:"column:account_id" json:"accountId"`
	IsDir      int       `gorm:"column:is_dir" json:"isDir"`
	ParentID   int64     `gorm:"column:parent_id" json:"parentId"`
	FileID     *int64    `gorm:"column:file_id" json:"fileId"`
	FileName   string    `gorm:"column:file_name" json:"fileName"`
	FileType   string    `gorm:"column:file_type" json:"fileType"`
	FileSuffix string    `gorm:"column:file_suffix" json:"fileSuffix"`
	FileSize   int64     `gorm:"column:file_size" json:"fileSize"`
	Del        bool      `gorm:"column:del" json:"del"`
	DelTime    *time.Time `gorm:"column:del_time" json:"delTime"`
	GmtCreate  time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

func (AccountFile) TableName() string { return "account_file" }

type AccountFileDTO struct {
	ID          int64  `json:"id"`
	AccountID   int64  `json:"accountId"`
	IsDir       int    `json:"isDir"`
	ParentID    int64  `json:"parentId"`
	FileID      *int64 `json:"fileId"`
	FileName    string `json:"fileName"`
	FileType    string `json:"fileType"`
	FileSuffix  string `json:"fileSuffix"`
	FileSize    int64  `json:"fileSize"`
	Del         bool   `json:"del"`
	DelTime     string `json:"delTime"`
	GmtCreate   string `json:"gmtCreate"`
	GmtModified string `json:"gmtModified"`
}

type File struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	AccountID  int64     `gorm:"column:account_id" json:"accountId"`
	FileName   string    `gorm:"column:file_name" json:"fileName"`
	FileSuffix string    `gorm:"column:file_suffix" json:"fileSuffix"`
	FileSize   int64     `gorm:"column:file_size" json:"fileSize"`
	ObjectKey  string    `gorm:"column:object_key" json:"objectKey"`
	Identifier string    `gorm:"column:identifier" json:"identifier"`
	Del        bool      `gorm:"column:del" json:"del"`
	GmtCreate  time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

func (File) TableName() string { return "file" }

type FileDownloadDTO struct {
	FileName     string `json:"fileName"`
	DownloadURL  string `json:"downloadUrl"`
}

type FolderTreeNodeDTO struct {
	ID       int64             `json:"id"`
	ParentID int64             `json:"parentId"`
	Label    string            `json:"label"`
	Children []*FolderTreeNodeDTO `json:"children"`
}
