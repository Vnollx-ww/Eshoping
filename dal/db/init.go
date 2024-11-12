package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func init() {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "Eshop"
	username := "root"
	password := "ab147890"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, _ := gorm.Open(driverName, args)
	//迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Order{})
	//初始化后的db赋给全局变量DB，没赋值的化DB就是一个nil
	DB = db
}
