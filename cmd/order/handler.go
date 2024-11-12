package main

import (
	"Eshop/dal/db"
	orderlist "Eshop/kitex_gen/orderlist"
	"Eshop/kitex_gen/product"
	"Eshop/kitex_gen/product/productservice"
	"Eshop/kitex_gen/user"
	"Eshop/kitex_gen/user/userservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
)

// OrderListServiceImpl implements the last service interface defined in the IDL.
type OrderListServiceImpl struct {
	usrcli userservice.Client
	procli productservice.Client
}

func NewUserClient(addr string) (userservice.Client, error) {
	return userservice.NewClient("user", client.WithHostPorts(addr))
}
func NewProductClient(addr string) (productservice.Client, error) {
	return productservice.NewClient("product", client.WithHostPorts(addr))
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
	log.Println(order)
	usr, err := db.GetUserByID(ctx, ol.UserId)
	if err != nil {
		log.Println(err)
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误",
		}
		return res, err
	}
	if usr == nil {
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		return res, nil
	}
	if usr.Balance < order.Cost {
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "余额不足",
		}
		return res, nil
	}
	pro, er := db.GetProductByName(ctx, ol.ProductName)
	if er != nil {
		log.Println(err)
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误",
		}
		return res, err
	}
	if pro == nil {
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "商品不存在",
		}
		return res, nil
	}
	if pro.Stock < order.Number {
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "商品库存不足",
		}
		return res, nil
	}
	err = db.CreateOrder(ctx, order)
	if err != nil {
		log.Println(err)
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "订单创建失败",
		}
		return res, nil
	}
	_, err = s.usrcli.UpdateCost(ctx, &user.UpdateCostRequest{UserId: order.UserID, Addcost: order.Cost})
	if err != nil {
		log.Println(err)
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "用户消费更改失败",
		}
		return res, nil
	}
	_, err = s.usrcli.UpdateBalance(ctx, &user.UpdateBalanceRequest{UserId: order.UserID, Addbalance: -order.Cost})
	if err != nil {
		log.Println(err)
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "用户余额更改失败",
		}
		return res, nil
	}
	_, err = s.procli.UpdateStock(ctx, &product.UpdateStockRequest{ProductId: int64(pro.ID), Addstock: -order.Number})
	if err != nil {
		log.Println(err)
		res := &orderlist.AddOrderResponse{
			StatusCode: -1,
			StatusMsg:  "商品库存更改失败",
		}
		return res, nil
	}
	res := &orderlist.AddOrderResponse{
		StatusCode: 200,
		StatusMsg:  "订单创建成功",
	}
	return res, nil
}

// DelOrder implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) DelOrder(ctx context.Context, req *orderlist.DelOrderRequest) (resp *orderlist.DelOrderResponse, err error) {
	er := db.DeleteOrder(ctx, req.OrderId)
	if er != nil {
		log.Println(err)
		res := &orderlist.DelOrderResponse{
			StatusCode: -1,
			StatusMsg:  "订单删除失败",
		}
		return res, nil
	}
	res := &orderlist.DelOrderResponse{
		StatusCode: 200,
		StatusMsg:  "订单删除成功",
	}
	return res, nil
}

// GetOrderListByUserID implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByUserID(ctx context.Context, req *orderlist.GetOrderListByUserIDRequest) (resp *orderlist.GetOrderListByUserIDResponse, err error) {
	orders, er := db.GetOrderListByUserID(ctx, req.UserId)
	if er != nil {
		log.Println(err)
		res := &orderlist.GetOrderListByUserIDResponse{
			StatusCode: -1,
			StatusMsg:  "用户订单列表获取失败",
		}
		return res, nil
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
	res := &orderlist.GetOrderListByUserIDResponse{
		StatusCode: 200,
		StatusMsg:  "用户订单列表获取成功",
		Orderlist:  or,
	}
	return res, nil
}

// GetOrderListByProductNameID implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByProductNameID(ctx context.Context, req *orderlist.GetOrderListByProductNameRequest) (resp *orderlist.GetOrderListByProductNameResponse, err error) {
	orders, er := db.GetOrderListByProductName(ctx, req.ProductName)
	if er != nil {
		log.Println(err)
		res := &orderlist.GetOrderListByProductNameResponse{
			StatusCode: -1,
			StatusMsg:  "商品订单列表获取失败",
		}
		return res, nil
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
	res := &orderlist.GetOrderListByProductNameResponse{
		StatusCode: 200,
		StatusMsg:  "商品订单列表获取成功",
		Orderlist:  or,
	}
	return res, nil
}
