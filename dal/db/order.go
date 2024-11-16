package db

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID      int64  `gorm:"not null" json:"userid,omitempty"`
	ProductName string `gorm:"varchar(40);not null" json:"productname,omitempty"`
	Number      int64  `gorm:"not null;default:0" json:"price,omitempty"`
	Cost        int64  `gorm:"not null;default:0" json:"cost,omitempty"`
	Address     string `gorm:"varchar(256);not null" json:"address,omitempty"`
	State       bool   `gorm:"default:false" json:"state,omitempty"`
}

func CreateOrder(ctx context.Context, order *Order) error {
	db := GetDB()
	err := db.Create(order).Error
	return err
}
func DeleteOrder(ctx context.Context, ID int64) error {
	order := new(Order)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&order).Error; err == nil {
		db.Unscoped().Delete(&order)
		return nil
	} else {
		return err
	}
}
func GetOrderByID(ctx context.Context, ID int64) (*Order, error) {
	order := new(Order)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&order).Error; err == nil {
		return order, nil
	} else {
		return nil, err
	}
}
func UpdateOrderState(ctx context.Context, order *Order) error {
	db := GetDB()
	order.State = true
	err := db.Save(&order).Error
	if err != nil {
		return err
	}
	return nil
}
func GetOrderListByUserID(ctx context.Context, userID int64) ([]*Order, error) {
	db := GetDB()
	var orders []*Order
	result := db.Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
func GetOrderListByProductName(ctx context.Context, productName string) ([]*Order, error) {
	db := GetDB()
	var orders []*Order
	result := db.Where("product_name = ?", productName).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
func GetOrderListByState(ctx context.Context, state bool) ([]*Order, error) {
	db := GetDB()
	var orders []*Order
	result := db.Where("state = ?", state).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
