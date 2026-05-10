package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/config"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/util"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"gorm.io/gorm"
)

type ChunkService struct{}

func NewChunkService() *ChunkService { return &ChunkService{} }

func (s *ChunkService) InitChunkTask(ctx context.Context, req *model.FileChunkInitTaskReq) (*model.FileChunkDTO, error) {
	db := config.GetDB()

	totalSize := int64(req.TotalSize)
	chunkSize := int64(req.ChunkSize)

	var storage model.Storage
	if err := db.WithContext(ctx).Where("account_id = ?", req.AccountID).First(&storage).Error; err != nil {
		return nil, err
	}
	if storage.UsedSize+totalSize > storage.TotalSize {
		return nil, fmt.Errorf("存储空间不足")
	}

	var existing model.FileChunk
	db.WithContext(ctx).Where("identifier = ? AND account_id = ?", req.Identifier, req.AccountID).First(&existing)
	if existing.ID != 0 && existing.UploadID != "" {
		parts, _ := getStoreEngine().ListMultipartParts(ctx, existing.ObjectKey, existing.UploadID)
		dto := toChunkDTO(&existing)
		dto.ExitPartList = parts
		return dto, nil
	}

	objectKey := fmt.Sprintf("%d/%d/%d/%s",
		time.Now().Year(), time.Now().Month(), time.Now().Day(), req.Identifier)

	uploadID, err := getStoreEngine().CreateMultipartUpload(ctx, objectKey)
	if err != nil {
		return nil, err
	}

	chunkNum := int((totalSize + chunkSize - 1) / chunkSize)

	chunk := model.FileChunk{
		ID:         util.NextID(),
		Identifier: req.Identifier,
		UploadID:   uploadID,
		FileName:   req.FileName,
		BucketName: getStoreEngine().GetBucket(),
		ObjectKey:  objectKey,
		TotalSize:  totalSize,
		ChunkSize:  chunkSize,
		ChunkNum:   chunkNum,
		AccountID:  req.AccountID,
		GmtCreate:  time.Now(),
		GmtModified: time.Now(),
	}
	db.WithContext(ctx).Create(&chunk)
	return toChunkDTO(&chunk), nil
}

func (s *ChunkService) GetPresignedURL(ctx context.Context, identifier string, accountID int64, partNumber int32) (string, error) {
	db := config.GetDB()
	var chunk model.FileChunk
	if err := db.WithContext(ctx).Where("identifier = ? AND account_id = ?", identifier, accountID).First(&chunk).Error; err != nil {
		return "", fmt.Errorf("分片任务不存在")
	}

	// Check if task is older than 24h
	if time.Since(chunk.GmtCreate) > 24*time.Hour {
		return "", fmt.Errorf("分片任务已过期")
	}

	return getStoreEngine().PresignUploadPart(ctx, chunk.ObjectKey, chunk.UploadID, partNumber, 1*time.Hour)
}

func (s *ChunkService) MergeChunks(ctx context.Context, req *model.FileChunkMergeReq) error {
	db := config.GetDB()
	var chunk model.FileChunk
	if err := db.WithContext(ctx).Where("identifier = ? AND account_id = ?", req.Identifier, req.AccountID).First(&chunk).Error; err != nil {
		return fmt.Errorf("分片任务不存在")
	}

	// 检查分片任务是否过期（24小时）
	if time.Since(chunk.GmtCreate) > 24*time.Hour {
		return fmt.Errorf("分片任务已过期")
	}

	// 获取已上传的分片列表
	parts, err := getStoreEngine().ListMultipartParts(ctx, chunk.ObjectKey, chunk.UploadID)
	if err != nil {
		return err
	}

	// 检查分片数量是否匹配
	if len(parts) < chunk.ChunkNum {
		return fmt.Errorf("分片数量不匹配，期望: %d, 实际: %d", chunk.ChunkNum, len(parts))
	}

	// 检查存储空间是否足够
	var storage model.Storage
	if err := db.WithContext(ctx).Where("account_id = ?", req.AccountID).First(&storage).Error; err != nil {
		return fmt.Errorf("用户存储信息不存在")
	}
	realSize := int64(0)
	for _, p := range parts {
		realSize += p.Size
	}
	if storage.UsedSize+realSize > storage.TotalSize {
		return fmt.Errorf("存储空间不足")
	}

	// 构建 CompletedPart 列表
	completedParts := make([]types.CompletedPart, 0, len(parts))
	for _, p := range parts {
		completedParts = append(completedParts, types.CompletedPart{
			ETag:       aws.String(p.ETag),
			PartNumber: aws.Int32(int32(p.PartNumber)),
		})
	}

	// 合并分片
	if err := getStoreEngine().CompleteMultipartUpload(ctx, chunk.ObjectKey, chunk.UploadID, completedParts); err != nil {
		return err
	}

	// 更新存储空间
	if err := db.WithContext(ctx).Model(&model.Storage{}).
		Where("account_id = ?", chunk.AccountID).
		UpdateColumn("used_size", gorm.Expr("used_size + ?", realSize)).Error; err != nil {
		return err
	}

	// 创建文件记录
	fileID := util.NextID()
	ext := ""
	if dot := chunk.FileName; dot != "" {
		if idx := len(dot) - 1; idx > 0 && dot[idx] == '.' {
			ext = dot[idx+1:]
		}
	}
	fileRecord := model.File{
		ID:          fileID,
		AccountID:   chunk.AccountID,
		FileName:    chunk.FileName,
		FileSuffix:  ext,
		FileSize:    chunk.TotalSize,
		ObjectKey:   chunk.ObjectKey,
		Identifier:  chunk.Identifier,
		Del:         false,
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
	}
	db.WithContext(ctx).Create(&fileRecord)

	accountFile := model.AccountFile{
		ID:         util.NextID(),
		AccountID:  chunk.AccountID,
		IsDir:      0,
		ParentID:   req.ParentID,
		FileID:     &fileID,
		FileName:   chunk.FileName,
		FileType:   util.DetectFileType(ext),
		FileSuffix: ext,
		FileSize:   chunk.TotalSize,
		Del:        false,
		GmtCreate:  time.Now(),
		GmtModified: time.Now(),
	}
	db.WithContext(ctx).Create(&accountFile)

	return db.WithContext(ctx).Delete(&chunk).Error
}

func (s *ChunkService) GetProgress(ctx context.Context, accountID int64, identifier string) (*model.FileChunkDTO, error) {
	db := config.GetDB()
	var chunk model.FileChunk
	if err := db.WithContext(ctx).Where("identifier = ? AND account_id = ?", identifier, accountID).First(&chunk).Error; err != nil {
		return nil, fmt.Errorf("分片任务不存在")
	}
	dto := toChunkDTO(&chunk)
	dto.Finished = false

	// 检查文件在存储端是否存在
	exist, _ := getStoreEngine().DoesObjectExist(ctx, chunk.ObjectKey)
	if !exist {
		// 获取已上传的分片列表
		parts, _ := getStoreEngine().ListMultipartParts(ctx, chunk.ObjectKey, chunk.UploadID)
		dto.ExitPartList = parts
		if len(parts) == chunk.ChunkNum {
			dto.Finished = true
		}
	}
	return dto, nil
}

func toChunkDTO(c *model.FileChunk) *model.FileChunkDTO {
	return &model.FileChunkDTO{
		ID:         c.ID,
		Identifier: c.Identifier,
		UploadID:   c.UploadID,
		FileName:   c.FileName,
		BucketName: c.BucketName,
		ObjectKey:  c.ObjectKey,
		TotalSize:  c.TotalSize,
		ChunkSize:  c.ChunkSize,
		ChunkNum:   c.ChunkNum,
		AccountID:  c.AccountID,
	}
}
