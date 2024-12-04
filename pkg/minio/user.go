package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
)

func UserUploadFileToMinio(c context.Context, file *multipart.FileHeader, username string) (string, error) {
	// 设置 MinIO 上传的桶名称
	bucketName := "user"
	// 上传文件到 MinIO
	objectName := fmt.Sprintf("UserName=%s.jpg", username)
	// 打开文件内容
	uploadedFile, err := file.Open()
	if err != nil {
		log.Println("打开文件时出错:", err)
		return "", err
	}
	defer uploadedFile.Close()
	// 上传文件到 MinIO
	_, err = MinioClient.PutObject(
		c,
		bucketName,
		objectName,
		uploadedFile,
		file.Size, // 使用上传文件的大小
		minio.PutObjectOptions{
			ContentType: "image/jpeg", // 你可以根据文件的类型来设置 ContentType
		},
	)
	if err != nil {
		log.Println("上传文件到 MinIO 时出错:", err)
		return "", err
	}

	// 构建文件的 URL
	fileURL := fmt.Sprintf("http://localhost:9000/%s/%s", bucketName, objectName)
	return fileURL, nil
}
func UserDeleteFileMinio(c context.Context, username string) error {
	bucketName := "user"
	objectName := fmt.Sprintf("UserName=%s.jpg", username)
	// 删除文件
	err := MinioClient.RemoveObject(c, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalf("删除头像失败: %v", err)
		return err
	}
	return nil
}
