package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/base"
	"Eshop/kitex_gen/orderlist"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		UserId      int64
		ProductName string
		Number      int64
		Cost        int64
	}
	if err := c.Bind(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "无效的请求格式",
		})
		return
	}
	ol := &orderlist.Order{
		UserId:      reqbody.UserId,
		ProductName: reqbody.ProductName,
		Number:      reqbody.Number,
		Cost:        reqbody.Cost,
	}
	req := &orderlist.AddOrderRequest{
		Ol: ol,
	}
	res, _ := rpc.CreateOrder(ctx, req)
	if res.StatusCode == -1 {
		log.Println(res)
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  res.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, orderlist.AddOrderResponse{
		OrderId:    res.OrderId,
		StatusCode: http.StatusOK,
		StatusMsg:  "订单创建成功",
	})
}
func DeleteOrder(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		OrderId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "无效的请求格式",
		})
		return
	}
	req := &orderlist.DelOrderRequest{
		OrderId: reqbody.OrderId,
	}
	res, _ := rpc.DeleteOrder(ctx, req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "删除订单失败",
		})
		return
	}
	c.JSON(http.StatusOK, orderlist.DelOrderResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "订单删除成功",
		Succed:     true,
	})
}
func GetOrderListByUserID(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		UserId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "无效的请求格式",
		})
		return
	}
	req := &orderlist.GetOrderListByUserIDRequest{
		UserId: reqbody.UserId,
	}
	res, _ := rpc.GetOrderListByUserID(ctx, req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "获取用户订单列表失败",
		})
	}
	c.JSON(http.StatusOK, orderlist.GetOrderListByUserIDResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "获取用户订单列表成功",
		Orderlist:  res.Orderlist,
	})
}
func GetOrderListByProductName(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductName string
	}
	if err := c.Bind(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "无效的请求格式",
		})
		return
	}
	req := &orderlist.GetOrderListByProductNameRequest{
		ProductName: reqbody.ProductName,
	}
	res, _ := rpc.GetOrderListByProductName(ctx, req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusBadRequest, base.Base{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "获取商品订单列表失败",
		})
	}
	c.JSON(http.StatusOK, orderlist.GetOrderListByProductNameResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "获取商品订单列表成功",
		Orderlist:  res.Orderlist,
	})
}
