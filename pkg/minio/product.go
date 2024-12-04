package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
)

func ProductUploadFileToMinio(c context.Context, file *multipart.FileHeader, productname string) (string, error) {
	// 设置 MinIO 上传的桶名称
	bucketName := "product"
	// 上传文件到 MinIO
	objectName := fmt.Sprintf("ProductName=%s.jpg", productname)
	// 打开文件内容
	uploadedFile, err := file.Open()
	if err != nil {
		log.Println("打开文件时出错:", err)
		return "", err
	}
	defer uploadedFile.Close()
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
	fileURL := fmt.Sprintf("http://localhost:9000/%s/%s", bucketName, objectName)
	return fileURL, nil
}
func ProductDeleteFileMinio(c context.Context, productname string) error {
	bucketName := "product"
	objectName := fmt.Sprintf("ProductName=%s.jpg", productname)
	// 删除文件
	err := MinioClient.RemoveObject(c, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalf("删除商品图片失败: %v", err)
		return err
	}
	return nil
}
