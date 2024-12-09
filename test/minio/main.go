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
	err = db.AddFriend(context.Background(), usr, tousr)
	if err != nil {
		log.Println("添加失败")
	}
}
