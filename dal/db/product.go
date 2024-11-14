package db

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `gorm:"unique;varchar(40);not null" json:"name,omitempty"`
	Price       int64  `gorm:"default:0" json:"price,omitempty"`
	Stock       int64  `gorm:"default:0" json:"stock,omitempty"`
}

func CreateProduct(ctx context.Context, pro *Product) error {
	db := GetDB()
	err := db.Create(pro).Error
	if err != nil {
		return err
	}
	return nil
}
func GetProductByName(ctx context.Context, productName string) (*Product, error) {
	db := GetDB()
	product := new(Product)
	if err := db.Where("product_name = ?", productName).First(&product).Error; err == nil {
		return product, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetProductByID(ctx context.Context, ID int64) (*Product, error) {
	db := GetDB()
	product := new(Product)
	if err := db.Where("ID = ?", ID).First(&product).Error; err == nil {
		return product, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetProductListInfo(ctx context.Context) ([]*Product, error) {
	db := GetDB()
	var products []*Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func DeleteProduct(ctx context.Context, pro *Product) error {
	db := GetDB()
	err := db.Unscoped().Delete(&pro).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateStock(ctx context.Context, pro *Product, stock int64) error {
	db := GetDB()
	pro.Stock += stock
	err := db.Save(&pro).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdatePrice(ctx context.Context, pro *Product, price int64) error {
	db := GetDB()
	pro.Price = price
	err := db.Save(&pro).Error
	if err != nil {
		return err
	}
	return nil
}
