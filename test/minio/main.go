package main

import (
	"Eshop/dal/db"
	"Eshop/pkg/zaplog"
	"context"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"log"
)

var MinioClient *minio.Client
var logger *zap.Logger = zaplog.GetLogger()

func main() {

	usr, err := db.GetUserByID(context.Background(), 17)
	tousr, err := db.GetUserByID(context.Background(), 16)
	err = db.SendMessage(context.Background(), usr, tousr, "哇咔咔!")
	if err != nil {
		log.Println("发送失败")
	}
}
