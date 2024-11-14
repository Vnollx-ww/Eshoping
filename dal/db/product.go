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
func DeleteProduct(ctx context.Context, ID int64) error {
	pro := new(Product)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&pro).Error; err == nil {
		db.Unscoped().Delete(&pro)
		return nil
	} else {
		return err
	}
}
func UpdateStock(ctx context.Context, ID int64, stock int64) error {
	pro := new(Product)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&pro).Error; err == nil {
		pro.Stock += stock
		db.Save(&pro)
		return nil
	} else {
		return err
	}
}
func UpdatePrice(ctx context.Context, ID int64, price int64) error {
	pro := new(Product)
	db := GetDB()
	if err := db.Where("ID = ?", ID).First(&pro).Error; err == nil {
		pro.Price = price
		db.Save(&pro)
		return nil
	} else {
		return err
	}
}
