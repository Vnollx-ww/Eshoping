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
	err := DB.Create(usr).Error
	return err
}
func GetUserByName(ctx context.Context, userName string) (*User, error) {
	user := new(User)
	//db.Where("user_name = ?", userName).First(&user)
	if err := DB.Where("user_name = ?", userName).First(&user).Error; err == nil {
		return user, nil
	} else if user.ID == 0 {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetUserByID(ctx context.Context, ID int64) (*User, error) {
	user := new(User)
	if err := DB.Where("ID = ?", ID).First(&user).Error; err == nil {
		return user, nil
	} else if user.ID == 0 {
		return nil, nil
	} else {
		return nil, err
	}
}
func UpdateName(ctx context.Context, user *User, name string) error {
	user.UserName = name
	err := DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdatePassword(ctx context.Context, usr *User, password string) error {
	usr.Password = password
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBalance(ctx context.Context, usr *User, addbalance int64) error {
	usr.Balance += addbalance
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateCost(ctx context.Context, usr *User, addcost int64) error {
	usr.Cost += addcost
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBalanceAndCost(ctx context.Context, usr *User, addbalance int64) error {
	usr.Balance -= addbalance
	usr.Cost += addbalance
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateAddress(ctx context.Context, usr *User, address string) error {
	usr.Address = address
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
