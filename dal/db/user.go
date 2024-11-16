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
	Address  string `gorm:"varchar(256);not null" json:"address,omitempty"`
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
func UpdateName(ctx context.Context, user *User, name string) error {
	db := GetDB()
	user.UserName = name
	err := db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdatePassword(ctx context.Context, usr *User, password string) error {
	db := GetDB()
	usr.Password = password
	err := db.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBalance(ctx context.Context, usr *User, addbalance int64) error {
	db := GetDB()
	usr.Balance += addbalance
	err := db.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateCost(ctx context.Context, usr *User, addcost int64) error {
	db := GetDB()
	usr.Cost += addcost
	err := db.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateAddress(ctx context.Context, usr *User, address string) error {
	db := GetDB()
	usr.Address = address
	err := db.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
