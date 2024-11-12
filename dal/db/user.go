package db

import (
	"context"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;varchar(40);not null" json:"name,omitempty"`
	Password string `gorm:"type:varchar(256);not null" json:"password,omitempty"`
	Balance  int64  `gorm:"default:0" json:"balance,omitempty"`
	Cost     int64  `gorm:"default:0" json:"cost,omitempty"`
}

func CreateUser(ctx context.Context, usr *User) error {
	db := GetDB()
	err := db.Create(usr).Error
	return err
}
func GetUserByName(ctx context.Context, userName string) (*User, error) {
	user := new(User)
	db := GetDB()
	//db.Where("user_name = ?", userName).First(&user)
	if err := db.Where("user_name = ?", userName).First(&user).Error; err == nil {
		return user, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetUserByID(ctx context.Context, ID int64) (*User, error) {
	user := new(User)
	db := GetDB()
	//db.Where("user_Id = ?", userId).First(&user)
	if err := db.Where("ID = ?", ID).First(&user).Error; err == nil {
		return user, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
func UpdateName(ctx context.Context, ID int64, name string) error {
	user := new(User)
	db := GetDB()
	//db.Where("user_Id = ?", userId).First(&user)
	if err := db.Where("ID = ?", ID).First(&user).Error; err == nil {
		user.UserName = name
		db.Save(&user)
		return nil
	} else {
		return err
	}
}
func UpdatePassword(ctx context.Context, ID int64, password string) error {
	user := new(User)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&user).Error; err == nil {
		user.Password = password
		db.Save(&user)
		return nil
	} else {
		return err
	}
}
func UpdateBalance(ctx context.Context, ID int64, addbalance int64) error {
	user := new(User)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&user).Error; err == nil {
		user.Balance += addbalance
		db.Save(&user)
		return nil
	} else {
		return err
	}
}
func UpdateCost(ctx context.Context, ID int64, addcost int64) error {
	user := new(User)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&user).Error; err == nil {
		user.Cost += addcost
		db.Save(&user)
		return nil
	} else {
		return err
	}
}
