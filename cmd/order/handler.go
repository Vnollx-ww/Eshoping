package main

import (
	"Eshop/dal/db"
	"Eshop/dal/rs"
	"Eshop/kitex_gen/orderlist"
	"Eshop/kitex_gen/product/productservice"
	"Eshop/kitex_gen/user/userservice"
	_ "Eshop/pkg/jaeger"
	order2 "Eshop/pkg/kafka/producer/order"
	"Eshop/pkg/middlerware"
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
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
func BadGetOrderListByProductNameResponse(s string) *orderlist.GetOrderListByProductNameResponse {
	return &orderlist.GetOrderListByProductNameResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetOrderListByProductNameResponse(s string, or []*orderlist.Order) *orderlist.GetOrderListByProductNameResponse {
	return &orderlist.GetOrderListByProductNameResponse{StatusCode: 200, StatusMsg: s, Orderlist: or}
}
func BadGetOrderListByStateResponse(s string) *orderlist.GetOrderListByStateResponse {
	return &orderlist.GetOrderListByStateResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetOrderListByStateResponse(s string, or []*orderlist.Order) *orderlist.GetOrderListByStateResponse {
	return &orderlist.GetOrderListByStateResponse{StatusCode: 200, StatusMsg: s, Orderlist: or}
}
func BadUpdateOrderStateResponse(s string) *orderlist.UpdateOrderStateResponse {
	return &orderlist.UpdateOrderStateResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateOrderStateResponse(s string) *orderlist.UpdateOrderStateResponse {
	return &orderlist.UpdateOrderStateResponse{StatusCode: 200, StatusMsg: s, Succed: true}
}

// AddOrder implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) AddOrder(ctx context.Context, req *orderlist.AddOrderRequest) (resp *orderlist.AddOrderResponse, err error) {
	tracer := otel.Tracer("order-service")
	// 创建根 Span
	ctx, span := tracer.Start(ctx, "AddOrder")
	defer span.End()
	ol := req.Ol
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		span.RecordError(err) // 记录错误到 Span
		logger.Error("token解析失败：", zap.Error(err))
		return BadAddOrderResponse("token解析失败"), err
	}
	usrCtx, usrSpan := tracer.Start(ctx, "GetUserByID")
	usr, err := db.GetUserByID(usrCtx, mc.UserId)
	usrSpan.End()
	if err != nil {
		span.RecordError(err)
		logger.Error("服务器内部错误：", zap.Error(err))
		return BadAddOrderResponse("服务器内部错误"), err
	}
	order := &db.Order{
		ProductName: ol.ProductName,
		UserID:      mc.UserId,
		Number:      ol.Number,
		Cost:        ol.Cost,
		Address:     usr.Address,
	}
	proCtx, proSpan := tracer.Start(ctx, "GetProductByName")
	pro, er := db.GetProductByName(proCtx, ol.ProductName)
	proSpan.End()
	if er != nil {
		span.RecordError(er)
		logger.Error("服务器内部错误：", zap.Error(err))
		return BadAddOrderResponse("服务器内部错误"), err
	}
	if pro == nil {
		span.RecordError(fmt.Errorf("商品不存在"))
		return BadAddOrderResponse("商品不存在"), nil
	}
	if pro.Stock < order.Number {
		span.RecordError(fmt.Errorf("商品库存不足"))
		return BadAddOrderResponse("商品库存不足"), nil
	}
	if usr.Balance < order.Cost {
		span.RecordError(fmt.Errorf("用户余额不足"))
		return BadAddOrderResponse("用户余额不足"), nil
	}
	orderCtx, orderSpan := tracer.Start(ctx, "CreateOrder")
	err = db.CreateOrder(orderCtx, order)
	orderSpan.End()
	if err != nil {
		span.RecordError(fmt.Errorf("订单创建失败"))
		logger.Error("订单创建失败：", zap.Error(err))
		return BadAddOrderResponse("订单创建失败"), err
	}
	// 发送 Kafka 消息
	_, kafkaSpan := tracer.Start(ctx, "SendCreateOrderEvent")
	kafkaProducer, err := order2.NewKafkaProducer([]string{kafkaAddr})
	if err != nil {
		span.RecordError(err)
		kafkaSpan.End()
		logger.Error("kafka生产者创建失败：", zap.Error(err))
		return BadAddOrderResponse("Kafka生产者创建失败"), err
	}
	err = kafkaProducer.SendCreateOrderEvent(req.Token, order.Cost, order.Number, int64(pro.ID))
	kafkaSpan.End()
	if err != nil {
		logger.Error("订单创建成功，但更新消息发送失败：", zap.Error(err))
		return BadAddOrderResponse("订单创建成功，但更新消息发送失败"), err
	}
	u, err := db.GetUserByID(ctx, pro.UserID)
	if err != nil {
		logger.Error("服务器内部错误：", zap.Error(err))
		return BadAddOrderResponse("服务器内部错误"), err
	}
	err = db.UpdateBalance(ctx, u, order.Cost)
	if err != nil {
		logger.Error("服务器内部错误：", zap.Error(err))
		return BadAddOrderResponse("服务器内部错误"), err
	}
	cacheKey := fmt.Sprintf("userinfo:%d", pro.UserID)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodAddOrderResponse("订单创建成功"), nil
}

// DelOrder implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) DelOrder(ctx context.Context, req *orderlist.DelOrderRequest) (resp *orderlist.DelOrderResponse, err error) {
	er := db.DeleteOrder(ctx, req.OrderId)
	if er != nil {
		logger.Error("订单删除失败：", zap.Error(er))
		return BadDelOrderResponse("订单删除失败"), err
	}
	return GoodDelOrderResponse("订单删除成功"), nil
}

// GetOrderListByUserID implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByUserID(ctx context.Context, req *orderlist.GetOrderListByUserIDRequest) (resp *orderlist.GetOrderListByUserIDResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadGetOrderListByUserIDResponse("token解析失败"), err
	}
	var or []*orderlist.Order
	orders, err := db.GetOrderListByUserID(ctx, mc.UserId)
	if err != nil {
		logger.Error("用户订单列表获取失败：", zap.Error(err))
		return BadGetOrderListByUserIDResponse("用户订单列表获取失败"), err
	}
	for _, order := range orders {
		var ord orderlist.Order
		ord.OrderId = int64(order.ID)
		ord.UserId = order.UserID
		ord.Number = order.Number
		ord.Cost = order.Cost
		ord.ProductName = order.ProductName
		ord.Address = order.Address
		ord.Isdeliver = order.State
		ord.CreateTime = order.CreatedAt.Format("2006-01-02 15:04:05")
		or = append(or, &ord)
	}
	return GoodGetOrderListByUserIDResponse("用户订单列表获取成功", or), nil
}

// GetOrderListByProductNameID implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByProductNameID(ctx context.Context, req *orderlist.GetOrderListByProductNameRequest) (resp *orderlist.GetOrderListByProductNameResponse, err error) {
	orders, er := db.GetOrderListByProductName(ctx, req.ProductName)
	if er != nil {
		logger.Error("商品订单列表获取失败：", zap.Error(er))
		return BadGetOrderListByProductNameResponse("商品订单列表获取失败"), err
	}
	var or []*orderlist.Order
	for _, order := range orders {
		var ord orderlist.Order
		ord.OrderId = int64(order.ID)
		ord.UserId = order.UserID
		ord.Number = order.Number
		ord.Cost = order.Cost
		ord.ProductName = order.ProductName
		ord.Address = order.Address
		ord.Isdeliver = order.State
		ord.CreateTime = order.CreatedAt.Format("2006-01-02 15:04:05")
		or = append(or, &ord)
	}
	return GoodGetOrderListByProductNameResponse("商品订单列表获取成功", or), nil
}

// GetOrderListByState implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) GetOrderListByState(ctx context.Context, req *orderlist.GetOrderListByStateRequest) (resp *orderlist.GetOrderListByStateResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadGetOrderListByStateResponse("token解析失败"), err
	}
	pro, err := db.GetProductListInfoByUser(ctx, mc.UserId)
	if err != nil {
		logger.Error("订单未发货列表获取失败：", zap.Error(err))
		return BadGetOrderListByStateResponse("订单未发货列表获取成功失败"), err
	}
	var or []*orderlist.Order
	for _, p := range pro {
		orders, err := db.GetOrderListByState(ctx, req.State, p.ProductName)
		if err != nil {
			logger.Error("用户订单列表获取失败：", zap.Error(err))
			return BadGetOrderListByStateResponse("用户订单列表获取失败"), err
		}
		for _, order := range orders {
			var ord orderlist.Order
			ord.OrderId = int64(order.ID)
			ord.UserId = order.UserID
			ord.Number = order.Number
			ord.Cost = order.Cost
			ord.ProductName = order.ProductName
			ord.Address = order.Address
			ord.Isdeliver = order.State
			ord.CreateTime = order.CreatedAt.Format("2006-01-02 15:04:05")
			or = append(or, &ord)
		}
	}
	return GoodGetOrderListByStateResponse("订单未发货列表获取成功", or), nil
}

// UpdateOrderState implements the OrderListServiceImpl interface.
func (s *OrderListServiceImpl) UpdateOrderState(ctx context.Context, req *orderlist.UpdateOrderStateRequest) (resp *orderlist.UpdateOrderStateResponse, err error) {
	order, err := db.GetOrderByID(ctx, req.OrderId)
	if err != nil {
		logger.Error("修改订单状态失败，服务器内部错误：", zap.Error(err))
		return BadUpdateOrderStateResponse("修改订单状态失败，服务器内部错误"), err
	}
	err = db.UpdateOrderState(ctx, order)
	if err != nil {
		logger.Error("订单状态修改失败：", zap.Error(err))
		return BadUpdateOrderStateResponse("订单状态修改失败"), err
	}
	return GoodUpdateOrderStateResponse("订单状态修改成功"), nil
}
