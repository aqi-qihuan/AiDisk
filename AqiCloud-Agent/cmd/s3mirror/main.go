// s3mirror - S3 存储迁移工具：跨 endpoint 迁移 bucket 数据 + 设置匿名读策略
// 连接配置一律通过环境变量提供（勿写入代码/命令行/文档）：
//
//	SRC_ENDPOINT / SRC_ACCESS_KEY / SRC_SECRET_KEY
//	DST_ENDPOINT / DST_ACCESS_KEY / DST_SECRET_KEY
//
// 用法:
//
//	go run ./cmd/s3mirror -mode mirror -bucket <bucket>
//	go run ./cmd/s3mirror -mode policy -bucket <bucket>   # 设置匿名只读
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/ptr"
)

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func client(ctx context.Context, endpoint, ak, sk string) *s3.Client {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(ak, sk, "")),
	)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})
}

// ensureBucket 目标 bucket 不存在时创建
func ensureBucket(ctx context.Context, c *s3.Client, bucket string) {
	_, err := c.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(bucket)})
	if err == nil {
		return
	}
	_, err = c.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: aws.String(bucket)})
	if err != nil {
		log.Fatalf("create bucket %s: %v", bucket, err)
	}
	log.Printf("bucket %s created", bucket)
}

func doMirror(ctx context.Context, src, dst *s3.Client, bucket string) {
	ensureBucket(ctx, dst, bucket)
	var migrated, failed int
	var pager = s3.NewListObjectsV2Paginator(src, &s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	for pager.HasMorePages() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("list %s: %v", bucket, err)
		}
		for _, obj := range page.Contents {
			get, err := src.GetObject(ctx, &s3.GetObjectInput{
				Bucket: aws.String(bucket), Key: obj.Key,
			})
			if err != nil {
				log.Printf("[FAIL get] %s: %v", *obj.Key, err)
				failed++
				continue
			}
			// 读入内存使其可 seek（HTTP 明文 + SDK trailing checksum 要求）
			data, err := io.ReadAll(get.Body)
			ct := ptr.ToString(get.ContentType)
			get.Body.Close()
			if err != nil {
				log.Printf("[FAIL read] %s: %v", *obj.Key, err)
				failed++
				continue
			}
			_, err = dst.PutObject(ctx, &s3.PutObjectInput{
				Bucket:      aws.String(bucket),
				Key:         obj.Key,
				Body:        bytes.NewReader(data),
				ContentType: aws.String(ct),
			})
			if err != nil {
				log.Printf("[FAIL put] %s: %v", *obj.Key, err)
				failed++
				continue
			}
			migrated++
			fmt.Printf("[OK] %s (%d bytes)\n", *obj.Key, len(data))
		}
	}
	log.Printf("bucket %s: migrated=%d failed=%d", bucket, migrated, failed)
	if failed > 0 {
		os.Exit(1)
	}
}

func doPolicy(ctx context.Context, dst *s3.Client, bucket string) {
	ensureBucket(ctx, dst, bucket)
	pol := map[string]any{
		"Version": "2012-10-17",
		"Statement": []map[string]any{{
			"Effect":    "Allow",
			"Principal": map[string]any{"AWS": []string{"*"}},
			"Action":    []string{"s3:GetObject"},
			"Resource":  []string{fmt.Sprintf("arn:aws:s3:::%s/*", bucket)},
		}},
	}
	raw, _ := json.Marshal(pol)
	_, err := dst.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
		Bucket: aws.String(bucket), Policy: aws.String(string(raw)),
	})
	if err != nil {
		log.Fatalf("put policy: %v", err)
	}
	log.Printf("bucket %s: anonymous read policy set", bucket)
}

func main() {
	mode := flag.String("mode", "mirror", "mirror | policy")
	bucket := flag.String("bucket", "", "bucket 名（必填）")
	srcEndpoint := flag.String("src-endpoint", env("SRC_ENDPOINT", ""), "源 endpoint")
	srcAK := flag.String("src-ak", env("SRC_ACCESS_KEY", ""), "源 access key")
	srcSK := flag.String("src-sk", env("SRC_SECRET_KEY", ""), "源 secret key")
	dstEndpoint := flag.String("dst-endpoint", env("DST_ENDPOINT", ""), "目标 endpoint")
	dstAK := flag.String("dst-ak", env("DST_ACCESS_KEY", ""), "目标 access key")
	dstSK := flag.String("dst-sk", env("DST_SECRET_KEY", ""), "目标 secret key")
	flag.Parse()

	if *bucket == "" {
		log.Fatal("必须指定 -bucket")
	}
	if *srcEndpoint == "" || *srcAK == "" || *srcSK == "" || *dstEndpoint == "" || *dstAK == "" || *dstSK == "" {
		log.Fatal("缺少 S3 连接配置：请通过环境变量 SRC_ENDPOINT/SRC_ACCESS_KEY/SRC_SECRET_KEY 与 DST_ENDPOINT/DST_ACCESS_KEY/DST_SECRET_KEY 提供（勿将密钥写入代码）")
	}
	ctx := context.Background()

	switch *mode {
	case "policy":
		doPolicy(ctx, client(ctx, *dstEndpoint, *dstAK, *dstSK), *bucket)
	case "mirror":
		doMirror(ctx,
			client(ctx, *srcEndpoint, *srcAK, *srcSK),
			client(ctx, *dstEndpoint, *dstAK, *dstSK),
			*bucket)
	default:
		log.Fatalf("未知 mode: %s", *mode)
	}
}
