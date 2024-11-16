package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/orderlist"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token       string
		ProductName string
		Number      int64
		Cost        int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	if reqbody.Number == 0 {
		BadBaseResponse(c, "购买数量不能为0")
	}
	ol := &orderlist.Order{
		ProductName: reqbody.ProductName,
		Number:      reqbody.Number,
		Cost:        reqbody.Cost,
	}

	req := &orderlist.AddOrderRequest{
		Ol:    ol,
		Token: reqbody.Token,
	}
	res, _ := rpc.CreateOrder(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, orderlist.AddOrderResponse{
		OrderId:    res.OrderId,
		StatusCode: http.StatusOK,
		StatusMsg:  res.StatusMsg,
	})
}
func DeleteOrder(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token   string `json:"token"`
		OrderId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &orderlist.DelOrderRequest{
		OrderId: reqbody.OrderId,
	}
	res, _ := rpc.DeleteOrder(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, orderlist.DelOrderResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  res.StatusMsg,
		Succed:     true,
	})
}
func GetOrderListByUserID(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token string
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &orderlist.GetOrderListByUserIDRequest{
		Token: reqbody.Token,
	}
	log.Println(reqbody.Token)
	res, _ := rpc.GetOrderListByUserID(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
	}
	c.JSON(http.StatusOK, orderlist.GetOrderListByUserIDResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  res.StatusMsg,
		Orderlist:  res.Orderlist,
	})

}
func GetOrderListByProductName(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		ProductName string
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &orderlist.GetOrderListByProductNameRequest{
		ProductName: reqbody.ProductName,
	}
	res, _ := rpc.GetOrderListByProductName(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, orderlist.GetOrderListByProductNameResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  res.StatusMsg,
		Orderlist:  res.Orderlist,
	})
}
func GetOrderListByState(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		state bool
	}
	if err := c.Bind(&reqbody); err != nil {
		log.Println(err)
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &orderlist.GetOrderListByStateRequest{
		State: reqbody.state,
	}
	res, _ := rpc.GetOrderListByState(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, orderlist.GetOrderListByStateResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  res.StatusMsg,
		Orderlist:  res.Orderlist,
	})
}
func UpdateOrderState(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		OrderId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		log.Println(err)
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &orderlist.UpdateOrderStateRequest{
		OrderId: reqbody.OrderId,
	}
	res, _ := rpc.UpdateOrderState(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, orderlist.UpdateOrderStateResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  res.StatusMsg,
		Succed:     true,
	})
}
