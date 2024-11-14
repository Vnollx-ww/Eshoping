package main

import (
	"Eshop/dal/db"
	product "Eshop/kitex_gen/product"
	"context"
	"log"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

func BadAddProductResponse(s string) *product.AddProductResponse {
	return &product.AddProductResponse{StatusCode: -1, StatusMsg: s}
}
func GoodAddProductResponse(s string, ID int64) *product.AddProductResponse {
	return &product.AddProductResponse{StatusCode: 200, StatusMsg: s, ProductId: ID}
}
func BadGetProductInfoResponse(s string) *product.GetProductInfoResponse {
	return &product.GetProductInfoResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetProductInfoResponse(s string, pro *product.Product) *product.GetProductInfoResponse {
	return &product.GetProductInfoResponse{StatusCode: 200, StatusMsg: s, Product: pro}
}
func BadGetProductListInfoResponse(s string) *product.GetProductListInfoResponse {
	return &product.GetProductListInfoResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetProductListInfoResponse(s string, prolist []*product.Product) *product.GetProductListInfoResponse {
	return &product.GetProductListInfoResponse{StatusCode: 200, StatusMsg: s, Productlist: prolist}
}
func BadDelProductResponse(s string) *product.DelProductResponse {
	return &product.DelProductResponse{StatusCode: -1, StatusMsg: s}
}
func GoodDelProductResponse(s string, flag bool) *product.DelProductResponse {
	return &product.DelProductResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdatePriceResponse(s string) *product.UpdatePriceResponse {
	return &product.UpdatePriceResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdatePriceResponse(s string, flag bool) *product.UpdatePriceResponse {
	return &product.UpdatePriceResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdateStockResponse(s string) *product.UpdateStockResponse {
	return &product.UpdateStockResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateStockResponse(s string, flag bool) *product.UpdateStockResponse {
	return &product.UpdateStockResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}

// AddProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) AddProduct(ctx context.Context, req *product.AddProductRequest) (resp *product.AddProductResponse, err error) {
	pro, err := db.GetProductByName(ctx, req.Name)
	if err != nil {
		log.Println(err)
		return BadAddProductResponse("添加失败，服务器内部错误"), err
	}
	if pro != nil {
		return BadAddProductResponse("商品已存在"), nil
	}
	pro = &db.Product{
		ProductName: req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	err = db.CreateProduct(ctx, pro)
	if err != nil {
		log.Println(err)
		return BadAddProductResponse("添加商品失败"), err
	}
	return GoodAddProductResponse("添加商品成功", int64(pro.ID)), nil
}

// GetProductInfo implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductInfo(ctx context.Context, req *product.GetProductInfoRequest) (resp *product.GetProductInfoResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.Productid)
	if err != nil {
		log.Println(err)
		return BadGetProductInfoResponse("获取商品信息失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadGetProductInfoResponse("商品不存在"), nil
	}
	p := &product.Product{
		Name:  pro.ProductName,
		Id:    int64(pro.ID),
		Price: pro.Price,
		Stock: pro.Stock,
	}
	return GoodGetProductInfoResponse("获取商品信息成功", p), nil
}

// DelProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DelProduct(ctx context.Context, req *product.DelProductRequest) (resp *product.DelProductResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.Productid)
	if err != nil {
		log.Println(err)
		return BadDelProductResponse("删除商品失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadDelProductResponse("商品不存在"), nil
	}
	err = db.DeleteProduct(ctx, pro)
	if err != nil {
		log.Println(err)
		return BadDelProductResponse("删除商品失败"), nil
	}
	return GoodDelProductResponse("删除商品成功", true), nil
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdatePrice(ctx context.Context, req *product.UpdatePriceRequest) (resp *product.UpdatePriceResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return BadUpdatePriceResponse("修改商品价格失败，服务器内部错误"), nil
	}
	if pro == nil {
		return BadUpdatePriceResponse("商品不存在"), nil
	}
	err = db.UpdatePrice(ctx, pro, req.Newprice_)
	if err != nil {
		log.Println(err)
		return BadUpdatePriceResponse("修改商品价格失败"), nil
	}
	return GoodUpdatePriceResponse("修改商品价格成功", true), nil
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateStock(ctx context.Context, req *product.UpdateStockRequest) (resp *product.UpdateStockResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return BadUpdateStockResponse("修改商品库存失败，服务器内部错误"), nil
	}
	if pro == nil {
		return BadUpdateStockResponse("商品不存在"), nil
	}
	err = db.UpdateStock(ctx, pro, req.Addstock)
	if err != nil {
		log.Println(err)
		return BadUpdateStockResponse("修改商品库存失败"), nil
	}
	return GoodUpdateStockResponse("修改商品库存成功", true), nil
}

// GetProductListInfo implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductListInfo(ctx context.Context) (resp *product.GetProductListInfoResponse, err error) {
	prolist, err := db.GetProductListInfo(ctx)
	if err != nil {
		log.Println(err)
		return BadGetProductListInfoResponse("获取商品列表信息失败"), nil
	}
	var productlist []*product.Product
	for _, pro := range prolist {
		var p product.Product
		p.Id = int64(pro.ID)
		p.Name = pro.ProductName
		p.Price = pro.Price
		p.Stock = pro.Stock
		productlist = append(productlist, &p)
	}
	return GoodGetProductListInfoResponse("获取商品列表信息成功", productlist), nil
}
