package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/aqi/AqiCloud-Agent/internal/config"
	"github.com/aqi/AqiCloud-Agent/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	storeEngine     *StoreEngine
	storeEngineOnce sync.Once
)

func getStoreEngine() *StoreEngine {
	storeEngineOnce.Do(func() {
		storeEngine = NewStoreEngine()
	})
	return storeEngine
}

type StoreEngine struct {
	client        *s3.Client
	bucket        string
	avatarBucket  string
	presignClient *s3.PresignClient
}

func NewStoreEngine() *StoreEngine {
	cfg := config.GetConfig()
	creds := credentials.NewStaticCredentialsProvider(cfg.MinIOAccessKeyID, cfg.MinIOSecretAccessKey, "")

	// 内部端点用于后端与 MinIO 通信
	internalEndpoint := "http://" + cfg.MinIOEndpoint
	awsCfg := aws.Config{
		Region:       "us-east-1",
		BaseEndpoint: aws.String(internalEndpoint),
		Credentials:  creds,
	}
	// 强制使用 path-style 访问，避免 virtual-hosted-style 导致的 DNS 解析问题
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	presignClient := s3.NewPresignClient(client)

	return &StoreEngine{
		client:        client,
		bucket:        cfg.MinIOBucketName,
		avatarBucket:  cfg.MinIOAvatarBucket,
		presignClient: presignClient,
	}
}

func (e *StoreEngine) Upload(ctx context.Context, reader io.Reader, objectKey string, size int64) error {
	_, err := e.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(e.bucket),
		Key:           aws.String(objectKey),
		Body:          reader,
		ContentLength: aws.Int64(size),
	})
	return err
}

func (e *StoreEngine) UploadAvatar(ctx context.Context, reader io.Reader, filename string) (string, error) {
	ext := filepath.Ext(filename)
	now := time.Now()
	key := fmt.Sprintf("%d/%d/%d/%s%s", now.Year(), now.Month(), now.Day(),
		fmt.Sprintf("%d", time.Now().UnixNano()), ext)

	_, err := e.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(e.avatarBucket),
		Key:    aws.String(key),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}

	cfg := config.GetConfig()
	// 使用外部端点生成 URL（通过 Nginx 代理）
	externalEndpoint := cfg.MinIOExternalEndpoint
	if externalEndpoint == "" {
		externalEndpoint = "http://" + cfg.MinIOEndpoint
	}
	return fmt.Sprintf("%s/%s/%s", externalEndpoint, e.avatarBucket, key), nil
}

func (e *StoreEngine) Download(ctx context.Context, objectKey string) (io.ReadCloser, int64, error) {
	resp, err := e.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, 0, err
	}
	size := int64(0)
	if resp.ContentLength != nil {
		size = *resp.ContentLength
	}
	return resp.Body, size, nil
}

func (e *StoreEngine) Delete(ctx context.Context, objectKey string) error {
	_, err := e.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(objectKey),
	})
	return err
}

func (e *StoreEngine) GeneratePresignedGetURL(ctx context.Context, objectKey string, expires time.Duration) (string, error) {
	req, err := e.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(objectKey),
	}, func(o *s3.PresignOptions) {
		o.Expires = expires
	})
	if err != nil {
		return "", err
	}
	// 替换 URL 为外部端点（通过 Nginx 代理）
	cfg := config.GetConfig()
	if cfg.MinIOExternalEndpoint != "" {
		// 将内部端点 URL 替换为外部端点
		internalURL := fmt.Sprintf("http://%s/%s/", cfg.MinIOEndpoint, e.bucket)
		externalURL := fmt.Sprintf("%s/%s/", cfg.MinIOExternalEndpoint, e.bucket)
		return strings.Replace(req.URL, internalURL, externalURL, 1), nil
	}
	return req.URL, nil
}

func (e *StoreEngine) GeneratePresignedPutURL(ctx context.Context, objectKey string, expires time.Duration) (string, error) {
	req, err := e.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(objectKey),
	}, func(o *s3.PresignOptions) {
		o.Expires = expires
	})
	if err != nil {
		return "", err
	}
	// 替换 URL 为外部端点（通过 Nginx 代理）
	cfg := config.GetConfig()
	if cfg.MinIOExternalEndpoint != "" {
		internalURL := fmt.Sprintf("http://%s/%s/", cfg.MinIOEndpoint, e.bucket)
		externalURL := fmt.Sprintf("%s/%s/", cfg.MinIOExternalEndpoint, e.bucket)
		return strings.Replace(req.URL, internalURL, externalURL, 1), nil
	}
	return req.URL, nil
}

func (e *StoreEngine) PresignUploadPart(ctx context.Context, objectKey, uploadID string, partNumber int32, expires time.Duration) (string, error) {
	req, err := e.presignClient.PresignUploadPart(ctx, &s3.UploadPartInput{
		Bucket:     aws.String(e.bucket),
		Key:        aws.String(objectKey),
		UploadId:   aws.String(uploadID),
		PartNumber: aws.Int32(partNumber),
	}, func(o *s3.PresignOptions) {
		o.Expires = expires
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (e *StoreEngine) CreateMultipartUpload(ctx context.Context, objectKey string) (string, error) {
	resp, err := e.client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return "", err
	}
	return aws.ToString(resp.UploadId), nil
}

func (e *StoreEngine) CompleteMultipartUpload(ctx context.Context, objectKey, uploadID string, parts []types.CompletedPart) error {
	_, err := e.client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(e.bucket),
		Key:      aws.String(objectKey),
		UploadId: aws.String(uploadID),
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: parts,
		},
	})
	return err
}

func (e *StoreEngine) ListMultipartParts(ctx context.Context, objectKey, uploadID string) ([]model.PartSummary, error) {
	resp, err := e.client.ListParts(ctx, &s3.ListPartsInput{
		Bucket:   aws.String(e.bucket),
		Key:      aws.String(objectKey),
		UploadId: aws.String(uploadID),
	})
	if err != nil {
		return nil, err
	}
	parts := make([]model.PartSummary, 0, len(resp.Parts))
	for _, p := range resp.Parts {
		lm := ""
		if p.LastModified != nil {
			lm = p.LastModified.Format("2006-01-02 15:04:05")
		}
		parts = append(parts, model.PartSummary{
			PartNumber:   int(*p.PartNumber),
			ETag:         aws.ToString(p.ETag),
			Size:         *p.Size,
			LastModified: lm,
		})
	}
	return parts, nil
}

func (e *StoreEngine) GetBucket() string       { return e.bucket }
func (e *StoreEngine) GetAvatarBucket() string { return e.avatarBucket }

func (e *StoreEngine) DoesObjectExist(ctx context.Context, objectKey string) (bool, error) {
	_, err := e.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		if _, ok := errors.AsType[*types.NotFound](err); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
