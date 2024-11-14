package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func AddProduct(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductName string
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.AddProductRequest{
		Name: reqbody.ProductName,
	}
	res, _ := rpc.AddProduct(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, product.AddProductResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		ProductId:  res.ProductId,
	})
}
func DelProduct(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.DelProductRequest{
		Productid: reqbody.ProductId,
	}
	res, _ := rpc.DelProduct(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, product.DelProductResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func GetProductInfo(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.GetProductInfoRequest{
		Productid: reqbody.ProductId,
	}
	res, _ := rpc.GetProductInfo(ctx, req)
	if res == nil {
		return
	}
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, product.GetProductInfoResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Product:    res.Product,
	})
}
func GetProductListInfo(ctx context.Context, c *app.RequestContext) {
	res, _ := rpc.GetProductListInfo(ctx)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, product.GetProductListInfoResponse{
		StatusMsg:   res.StatusMsg,
		StatusCode:  res.StatusCode,
		Productlist: res.Productlist,
	})
}
func Updatestock(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductId int64
		AddStock  int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.UpdateStockRequest{
		ProductId: reqbody.ProductId,
		Addstock:  reqbody.AddStock,
	}
	res, _ := rpc.UpdateStock(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, product.UpdateStockResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func UpdatePrice(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductId int64
		Price     int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.UpdatePriceRequest{
		ProductId: reqbody.ProductId,
		Newprice_: reqbody.Price,
	}
	res, _ := rpc.UpdatePrice(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, product.UpdatePriceResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
