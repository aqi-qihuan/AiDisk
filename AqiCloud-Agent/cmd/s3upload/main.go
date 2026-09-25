// s3upload - 部署辅助工具：向 S3 兼容存储（RustFS/MinIO）上传文件并生成 presigned 下载链接
// 用法:
//
//	go run ./cmd/s3upload -file bin/agentpan-server -key deploy/agentpan-server -bucket deploy
//
// 环境变量: S3_ENDPOINT / S3_ACCESS_KEY / S3_SECRET_KEY / S3_BUCKET
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	file := flag.String("file", "", "本地待上传文件（必填）")
	key := flag.String("key", "", "S3 object key（必填）")
	bucket := flag.String("bucket", env("S3_BUCKET", "deploy"), "目标 bucket")
	list := flag.Bool("list", false, "仅列出所有 bucket")
	endpoint := flag.String("endpoint", env("S3_ENDPOINT", "http://134.175.206.158:9000"), "S3 endpoint")
	accessKey := flag.String("ak", env("S3_ACCESS_KEY", "aqi1015."), "access key")
	secretKey := flag.String("sk", env("S3_SECRET_KEY", "aqi1015."), "secret key")
	expire := flag.Int("expire-min", 60, "presigned URL 有效期（分钟）")
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(*accessKey, *secretKey, "")),
	)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// endpoint 走 HTTP 自定义地址
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(*endpoint)
		o.UsePathStyle = true
	})

	// -list 模式：仅列出 bucket
	if *list {
		out, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
		if err != nil {
			log.Fatalf("列举 bucket 失败: %v", err)
		}
		for _, b := range out.Buckets {
			fmt.Println(*b.Name)
		}
		return
	}

	if *file == "" || *key == "" {
		log.Fatal("必须指定 -file 和 -key")
	}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}
	defer f.Close()
	fi, _ := f.Stat()
	log.Printf("上传 %s (%d MB) -> %s/%s", *file, fi.Size()/1024/1024, *bucket, *key)

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(*key),
		Body:   f,
	})
	if err != nil {
		log.Fatalf("上传失败: %v", err)
	}
	log.Printf("上传成功")

	// 生成 presigned GET URL
	presign := s3.NewPresignClient(client)
	req, err := presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(*key),
	}, s3.WithPresignExpires(time.Duration(*expire)*time.Minute))
	if err != nil {
		log.Fatalf("生成 presigned URL 失败: %v", err)
	}
	fmt.Println(req.URL)
}
