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
	err := DB.Create(pro).Error
	if err != nil {
		return err
	}
	return nil
}
func GetProductByName(ctx context.Context, productName string) (*Product, error) {
	product := new(Product)
	if err := DB.Where("product_name = ?", productName).First(&product).Error; err == nil {
		return product, nil
	} else if product.ID == 0 {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetProductByID(ctx context.Context, ID int64) (*Product, error) {
	product := new(Product)
	if err := DB.Where("ID = ?", ID).First(&product).Error; err == nil {
		return product, nil
	} else if product.ID == 0 {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetProductListInfo(ctx context.Context) ([]*Product, error) {
	var products []*Product
	if err := DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func DeleteProduct(ctx context.Context, pro *Product) error {
	err := DB.Unscoped().Delete(&pro).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateStock(ctx context.Context, pro *Product, stock int64) error {
	pro.Stock += stock
	err := DB.Save(&pro).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdatePrice(ctx context.Context, pro *Product, price int64) error {
	pro.Price = price
	err := DB.Save(&pro).Error
	if err != nil {
		return err
	}
	return nil
}
