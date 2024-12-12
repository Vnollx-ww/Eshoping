package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/product"
	"Eshop/pkg/minio"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"net/http"
)

func AddProduct(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductName string `form:"productname"`
		Stock       int64  `form:"stock"`
		Price       int64  `form:"price"`
		Token       string `form:"token"`
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	file, err := c.FormFile("productimage")
	if err != nil {
		logger.Error("文件上传错误", zap.Error(err))
		BadBaseResponse(c, "图片上传失败")
		return
	}
	fileURL := fmt.Sprintf("http://localhost:9000/product/ProductName=%s.jpg", reqbody.ProductName)
	req := &product.AddProductRequest{
		Name:         reqbody.ProductName,
		Stock:        reqbody.Stock,
		Price:        reqbody.Price,
		Productimage: fileURL,
		Token:        reqbody.Token,
	}
	res, err := rpc.AddProduct(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
		return
	}
	err = minio.ProductUploadFileToMinio(ctx, file, reqbody.ProductName)
	if err != nil {
		logger.Error("商品图片上传到 MinIO 失败", zap.Error(err))
		BadBaseResponse(c, "商品图片上传到 MinIO 失败")
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.DelProductRequest{
		Productid: reqbody.ProductId,
	}
	res, err := rpc.DelProduct(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.GetProductInfoRequest{
		Productid: reqbody.ProductId,
	}
	res, err := rpc.GetProductInfo(ctx, req)
	if res == nil {
		return
	}
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, product.GetProductInfoResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Product:    res.Product,
	})
}
func GetProductListInfo(ctx context.Context, c *app.RequestContext) {
	res, err := rpc.GetProductListInfo(ctx)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.UpdateStockRequest{
		ProductId: reqbody.ProductId,
		Addstock:  reqbody.AddStock,
	}
	res, err := rpc.UpdateStock(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.UpdatePriceRequest{
		ProductId: reqbody.ProductId,
		Newprice_: reqbody.Price,
	}
	res, err := rpc.UpdatePrice(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, product.UpdatePriceResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func GetProductListInfoByUser(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &product.GetProductListInfoByUserRequest{
		Token: reqbody.Token,
	}
	res, err := rpc.GetProductListInfoByUser(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, product.GetProductListInfoByUserResponse{
		StatusMsg:   res.StatusMsg,
		StatusCode:  res.StatusCode,
		Productlist: res.Productlist,
	})
}
