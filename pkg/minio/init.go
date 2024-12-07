package minio

import (
	"Eshop/pkg/zaplog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

// D:\minio.exe server D:\minio --console-address :9090
var MinioClient *minio.Client
var logger *zap.Logger = zaplog.GetLogger()

func init() {
	endpoint := "localhost:9000"
	accessKeyID := "KzpxenSDJNHgVVGOoiEr"
	secretAccessKey := "yyMiXKDqxNjzQnqYgzA71CZXMxyir99gm0mjqx47"

	// 初始化一个minio客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		logger.Fatal("初始化MinioClient错误:", zap.Error(err))
	}
	MinioClient = minioClient
}
