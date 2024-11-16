package rpc

import (
	"Eshop/kitex_gen/orderlist"
	"Eshop/kitex_gen/orderlist/orderlistservice"
	"Eshop/pkg/viper"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var (
	orderClient orderlistservice.Client
)

func InitOrder(config *viper.Config) {
	etcdAddr := fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	serviceName := config.Viper.GetString("server.name")
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}
	c, err := orderlistservice.NewClient(
		serviceName, // mux
		client.WithRPCTimeout(30*time.Second),
		client.WithConnectTimeout(30000*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		panic(err)
	}
	orderClient = c
}
func CreateOrder(ctx context.Context, req *orderlist.AddOrderRequest) (*orderlist.AddOrderResponse, error) {
	return orderClient.AddOrder(ctx, req)
}
func DeleteOrder(ctx context.Context, req *orderlist.DelOrderRequest) (*orderlist.DelOrderResponse, error) {
	return orderClient.DelOrder(ctx, req)
}
func UpdateOrderState(ctx context.Context, req *orderlist.UpdateOrderStateRequest) (*orderlist.UpdateOrderStateResponse, error) {
	return orderClient.UpdateOrderState(ctx, req)
}
func GetOrderListByUserID(ctx context.Context, req *orderlist.GetOrderListByUserIDRequest) (*orderlist.GetOrderListByUserIDResponse, error) {
	return orderClient.GetOrderListByUserID(ctx, req)
}
func GetOrderListByProductName(ctx context.Context, req *orderlist.GetOrderListByProductNameRequest) (*orderlist.GetOrderListByProductNameResponse, error) {
	return orderClient.GetOrderListByProductNameID(ctx, req)
}
func GetOrderListByState(ctx context.Context, req *orderlist.GetOrderListByStateRequest) (*orderlist.GetOrderListByStateResponse, error) {
	return orderClient.GetOrderListByState(ctx, req)
}
