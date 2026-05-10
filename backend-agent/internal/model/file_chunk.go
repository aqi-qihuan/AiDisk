package model

import "time"

type FileChunk struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Identifier  string    `gorm:"column:identifier" json:"identifier"`
	UploadID    string    `gorm:"column:upload_id" json:"uploadId"`
	FileName    string    `gorm:"column:file_name" json:"fileName"`
	BucketName  string    `gorm:"column:bucket_name" json:"bucketName"`
	ObjectKey   string    `gorm:"column:object_key" json:"objectKey"`
	TotalSize   int64     `gorm:"column:total_size" json:"totalSize"`
	ChunkSize   int64     `gorm:"column:chunk_size" json:"chunkSize"`
	ChunkNum    int       `gorm:"column:chunk_num" json:"chunkNum"`
	AccountID   int64     `gorm:"column:account_id" json:"accountId"`
	GmtCreate   time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

func (FileChunk) TableName() string { return "file_chunk" }

type FileChunkDTO struct {
	ID           int64          `json:"id"`
	Identifier   string         `json:"identifier"`
	UploadID     string         `json:"uploadId"`
	FileName     string         `json:"fileName"`
	BucketName   string         `json:"bucketName"`
	ObjectKey    string         `json:"objectKey"`
	TotalSize    int64          `json:"totalSize"`
	ChunkSize    int64          `json:"chunkSize"`
	ChunkNum     int            `json:"chunkNum"`
	AccountID    int64          `json:"accountId"`
	Finished     bool           `json:"finished"`
	ExitPartList []PartSummary  `json:"exitPartList"`
}

// PartSummary mirrors AWS SDK PartSummary / Java com.amazonaws.services.s3.model.PartSummary
type PartSummary struct {
	PartNumber   int    `json:"partNumber"`
	ETag         string `json:"eTag"`
	Size         int64  `json:"size"`
	LastModified string `json:"lastModified"`
}

type FileType struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FileTypeName string `gorm:"column:file_type_name" json:"fileTypeName"`
}

func (FileType) TableName() string { return "file_type" }

type FileSuffix struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FileSuffix string `gorm:"column:file_suffix" json:"fileSuffix"`
	FileTypeID int    `gorm:"column:file_type_id" json:"fileTypeId"`
}

func (FileSuffix) TableName() string { return "file_suffix" }
