package main

import (
	"Eshop/pkg/zaplog"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"log"
	"os"
)

var MinioClient *minio.Client
var logger *zap.Logger = zaplog.GetLogger()

func main() {
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
	ctx := context.Background()
	uploadFileToMinIO(ctx)
}
func uploadFileToMinIO(c context.Context) {
	// 设置 MinIO 上传的桶名称
	bucketName := "eshoping"

	// 上传文件到 MinIO
	objectName := "UserId=3.jpg"
	filePath := `C:\Users\吴恩宇\Desktop\3WAK.jpg`
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("打开文件时出错:", err)
		return
	}
	defer file.Close()
	_, err = MinioClient.PutObject(
		c,
		bucketName,
		objectName,
		file,
		-1, // 使用文件的默认大小
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)
	if err != nil {
		log.Println("上传文件到 MinIO 时出错: %v", err)
	}
	fileURL := fmt.Sprintf("http://localhost:9000/%s/%s", bucketName, objectName)
	log.Println(fileURL)
}
