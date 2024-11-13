package main

import (
	"Eshop/dal/db"
	orderlist "Eshop/kitex_gen/orderlist"
	"Eshop/kitex_gen/product"
	"Eshop/kitex_gen/product/productservice"
	"Eshop/kitex_gen/user"
	"Eshop/kitex_gen/user/userservice"
	"Eshop/pkg/middlerware"
	"context"
	"log"
)

// OrderListServiceImpl implements the last service interface defined in the IDL.
type OrderListServiceImpl struct {
	usrcli userservice.Client
	procli productservice.Client
}

func BadAddOrderResponse(s string) *orderlist.AddOrderResponse {
	return &orderlist.AddOrderResponse{StatusCode: -1, StatusMsg: s}
}
func GoodAddOrderResponse(s string) *orderlist.AddOrderResponse {
	return &orderlist.AddOrderResponse{StatusCode: 200, StatusMsg: s}
}
func BadDelOrderResponse(s string) *orderlist.DelOrderResponse {
	return &orderlist.DelOrderResponse{StatusCode: -1, StatusMsg: s}
}
func GoodDelOrderResponse(s string) *orderlist.DelOrderResponse {
	return &orderlist.DelOrderResponse{StatusCode: 200, StatusMsg: s}
}
func BadGetOrderListByUserIDResponse(s string) *orderlist.GetOrderListByUserIDResponse {
	return &orderlist.GetOrderListByUserIDResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetOrderListByUserIDResponse(s string, or []*orderlist.Order) *orderlist.GetOrderListByUserIDResponse {
	return &orderlist.GetOrderListByUserIDResponse{StatusCode: 200, StatusMsg: s, Orderlist: or}
}
func BadGetOrderListByProductName(s string) *orderlist.GetOrderListByProductNameResponse {
	return &orderlist.GetOrderListByProductNameResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetOrderListByProductNameResponse(s string, or []*orderlist.Order) *orderlist.GetOrderListByProductNameResponse {
	return &orderlist.GetOrderListByProductNameResponse{StatusCode: 200, StatusMsg: s, Orderlist: or}
}

// AddOrder implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) AddOrder(ctx context.Context, req *orderlist.AddOrderRequest) (resp *orderlist.AddOrderResponse, err error) {
	ol := req.Ol
	order := &db.Order{
		ProductName: ol.ProductName,
		UserID:      ol.UserId,
		Number:      ol.Number,
		Cost:        ol.Cost,
	}
	usr, err := db.GetUserByID(ctx, ol.UserId)
	if err != nil {
		log.Println(err)
		return BadAddOrderResponse("服务器内部错误"), err
	}
	pro, er := db.GetProductByName(ctx, ol.ProductName)
	if er != nil {
		log.Println(err)
		return BadAddOrderResponse("服务器内部错误"), err
	}
	if pro == nil {
		return BadAddOrderResponse("商品不存在"), err
	}
	if pro.Stock < order.Number {
		return BadAddOrderResponse("商品库存不足"), err
	}
	if usr == nil {
		return BadAddOrderResponse("用户不存在"), err
	}
	if usr.Balance < order.Cost {
		return BadAddOrderResponse("用户余额不足"), err
	}
	err = db.CreateOrder(ctx, order)
	if err != nil {
		log.Println(err)
		return BadAddOrderResponse("订单创建失败"), err
	}
	_, err = s.usrcli.UpdateCost(ctx, &user.UpdateCostRequest{UserId: order.UserID, Addcost: order.Cost})
	if err != nil {
		log.Println(err)
		return BadAddOrderResponse("用户消费总额更改失败"), err
	}
	_, err = s.usrcli.UpdateBalance(ctx, &user.UpdateBalanceRequest{UserId: order.UserID, Addbalance: -order.Cost})
	if err != nil {
		log.Println(err)
		return BadAddOrderResponse("用户余额更改失败"), err
	}
	_, err = s.procli.UpdateStock(ctx, &product.UpdateStockRequest{ProductId: int64(pro.ID), Addstock: -order.Number})
	if err != nil {
		log.Println(err)
		return BadAddOrderResponse("商品库存更改失败"), err
	}
	return GoodAddOrderResponse("订单创建成功"), nil
}

// DelOrder implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) DelOrder(ctx context.Context, req *orderlist.DelOrderRequest) (resp *orderlist.DelOrderResponse, err error) {
	er := db.DeleteOrder(ctx, req.OrderId)
	if er != nil {
		log.Println(er)
		return BadDelOrderResponse("订单删除失败"), nil
	}
	return GoodDelOrderResponse("订单删除成功"), nil
}

// GetOrderListByUserID implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByUserID(ctx context.Context, req *orderlist.GetOrderListByUserIDRequest) (resp *orderlist.GetOrderListByUserIDResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		return BadGetOrderListByUserIDResponse("token解析失败"), nil
	}
	orders, er := db.GetOrderListByUserID(ctx, mc.UserId)
	if er != nil {
		log.Println(err)
		return BadGetOrderListByUserIDResponse("用户订单列表获取失败"), nil
	}
	var or []*orderlist.Order
	for _, order := range orders {
		var ord orderlist.Order
		ord.OrderId = int64(order.ID)
		ord.UserId = order.UserID
		ord.Number = order.Number
		ord.Cost = order.Cost
		ord.ProductName = order.ProductName
		ord.CreateTime = order.CreatedAt.Format("2006-01-02 15:04:05")
		or = append(or, &ord)
	}
	return GoodGetOrderListByUserIDResponse("用户订单列表获取成功", or), nil
}

// GetOrderListByProductNameID implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByProductNameID(ctx context.Context, req *orderlist.GetOrderListByProductNameRequest) (resp *orderlist.GetOrderListByProductNameResponse, err error) {
	orders, er := db.GetOrderListByProductName(ctx, req.ProductName)
	if er != nil {
		log.Println(err)
		return BadGetOrderListByProductName("商品订单列表获取失败"), nil
	}
	var or []*orderlist.Order
	for _, order := range orders {
		var ord orderlist.Order
		ord.OrderId = int64(order.ID)
		ord.UserId = order.UserID
		ord.Number = order.Number
		ord.Cost = order.Cost
		ord.ProductName = order.ProductName
		ord.CreateTime = order.CreatedAt.Format("2006-01-02 15:04:05")
		or = append(or, &ord)
	}
	return GoodGetOrderListByProductNameResponse("商品订单列表获取成功", or), nil
}
