package rpc

import (
	"Eshop/kitex_gen/product"
	"Eshop/kitex_gen/product/productservice"
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
	productClient productservice.Client
)

func InitProduct(config *viper.Config) {
	etcdAddr := fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	serviceName := config.Viper.GetString("server.name")
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := productservice.NewClient(
		serviceName,
		//client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r), // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		panic(err)
	}
	productClient = c
}
func GerProductClient() productservice.Client {
	return productClient
}
func AddProduct(ctx context.Context, req *product.AddProductRequest) (*product.AddProductResponse, error) {
	return productClient.AddProduct(ctx, req)
}
func DelProduct(ctx context.Context, req *product.DelProductRequest) (*product.DelProductResponse, error) {
	return productClient.DelProduct(ctx, req)
}
func GetProductInfo(ctx context.Context, req *product.GetProductInfoRequest) (*product.GetProductInfoResponse, error) {
	return productClient.GetProductInfo(ctx, req)
}
func GetProductListInfo(ctx context.Context) (*product.GetProductListInfoResponse, error) {
	return productClient.GetProductListInfo(ctx)
}
func UpdatePrice(ctx context.Context, req *product.UpdatePriceRequest) (*product.UpdatePriceResponse, error) {
	return productClient.UpdatePrice(ctx, req)
}
func UpdateStock(ctx context.Context, req *product.UpdateStockRequest) (*product.UpdateStockResponse, error) {
	return productClient.UpdateStock(ctx, req)
}
