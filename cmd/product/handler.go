package main

import (
	"Eshop/dal/db"
	product "Eshop/kitex_gen/product"
	"context"
	"log"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// AddProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) AddProduct(ctx context.Context, req *product.AddProductRequest) (resp *product.AddProductResponse, err error) {
	pro, err := db.GetProductByName(ctx, req.Name)
	if err != nil {
		log.Println(err)
		res := &product.AddProductResponse{
			StatusCode: -1,
			StatusMsg:  "添加失败，服务器内部错误",
		}
		return res, err
	}
	if pro != nil {
		res := &product.AddProductResponse{
			StatusCode: -1,
			StatusMsg:  "商品已存在",
		}
		return res, nil
	}
	pro = &db.Product{
		ProductName: req.Name,
	}
	err = db.CreateProduct(ctx, pro)
	if err != nil {
		log.Println(err)
		res := &product.AddProductResponse{
			StatusCode: -1,
			StatusMsg:  "添加商品失败",
		}
		return res, err
	}
	res := &product.AddProductResponse{
		StatusCode: 200,
		StatusMsg:  "添加商品成功",
		ProductId:  int64(pro.ID),
	}
	return res, nil
}

// GetProductInfo implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductInfo(ctx context.Context, req *product.GetProductInfoRequest) (resp *product.GetProductInfoResponse, err error) {

	pro, err := db.GetProductByID(ctx, req.Productid)
	if err != nil {
		log.Println(err)
		res := &product.GetProductInfoResponse{
			StatusCode: -1,
			StatusMsg:  "获取商品信息失败，服务器内部错误",
		}
		return res, err
	}
	if pro == nil {
		res := &product.GetProductInfoResponse{
			StatusCode: -1,
			StatusMsg:  "商品不存在",
		}
		return res, nil
	}
	p := &product.Product{
		Name:  pro.ProductName,
		Id:    int64(pro.ID),
		Price: pro.Price,
		Stock: pro.Stock,
	}
	res := &product.GetProductInfoResponse{
		StatusCode: 200,
		StatusMsg:  "获取商品信息成功",
		Product:    p,
	}
	return res, nil
}

// DelProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DelProduct(ctx context.Context, req *product.DelProductRequest) (resp *product.DelProductResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.Productid)
	if err != nil {
		log.Println(err)
		res := &product.DelProductResponse{
			StatusCode: -1,
			StatusMsg:  "删除商品失败，服务器内部错误",
		}
		return res, err
	}
	if pro == nil {
		res := &product.DelProductResponse{
			StatusCode: -1,
			StatusMsg:  "商品不存在",
		}
		return res, nil
	}
	err = db.DeleteProduct(ctx, req.Productid)
	if err != nil {
		log.Println(err)
		res := &product.DelProductResponse{
			StatusCode: -1,
			StatusMsg:  "删除商品失败",
		}
		return res, nil
	}
	res := &product.DelProductResponse{
		StatusCode: 200,
		StatusMsg:  "删除商品成功",
		Succed:     true,
	}
	return res, nil
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdatePrice(ctx context.Context, req *product.UpdatePriceRequest) (resp *product.UpdatePriceResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		res := &product.UpdatePriceResponse{
			StatusCode: -1,
			StatusMsg:  "修改商品价格失败，服务器内部错误",
		}
		return res, nil
	}
	if pro == nil {
		res := &product.UpdatePriceResponse{
			StatusCode: -1,
			StatusMsg:  "商品不存在",
		}
		return res, nil
	}
	err = db.UpdatePrice(ctx, req.ProductId, req.Newprice_)
	if err != nil {
		log.Println(err)
		res := &product.UpdatePriceResponse{
			StatusCode: -1,
			StatusMsg:  "修改商品价格失败",
		}
		return res, nil
	}
	res := &product.UpdatePriceResponse{
		StatusCode: 200,
		StatusMsg:  "修改商品价格成功",
		Succed:     true,
	}
	return res, nil
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateStock(ctx context.Context, req *product.UpdateStockRequest) (resp *product.UpdateStockResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		res := &product.UpdateStockResponse{
			StatusCode: -1,
			StatusMsg:  "修改商品库存失败，服务器内部错误",
		}
		return res, nil
	}
	if pro == nil {
		res := &product.UpdateStockResponse{
			StatusCode: -1,
			StatusMsg:  "商品不存在",
		}
		return res, nil
	}
	err = db.UpdateStock(ctx, req.ProductId, req.Addstock)
	if err != nil {
		log.Println(err)
		res := &product.UpdateStockResponse{
			StatusCode: -1,
			StatusMsg:  "修改商品库存失败",
		}
		return res, nil
	}
	res := &product.UpdateStockResponse{
		StatusCode: 200,
		StatusMsg:  "修改商品库存成功",
	}
	return res, nil
}
