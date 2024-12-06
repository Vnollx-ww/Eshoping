// Code generated by Kitex v0.7.3. DO NOT EDIT.

package productservice

import (
	product "Eshop/kitex_gen/product"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddProduct(ctx context.Context, req *product.AddProductRequest, callOptions ...callopt.Option) (r *product.AddProductResponse, err error)
	GetProductInfo(ctx context.Context, req *product.GetProductInfoRequest, callOptions ...callopt.Option) (r *product.GetProductInfoResponse, err error)
	GetProductListInfo(ctx context.Context, callOptions ...callopt.Option) (r *product.GetProductListInfoResponse, err error)
	DelProduct(ctx context.Context, req *product.DelProductRequest, callOptions ...callopt.Option) (r *product.DelProductResponse, err error)
	UpdatePrice(ctx context.Context, req *product.UpdatePriceRequest, callOptions ...callopt.Option) (r *product.UpdatePriceResponse, err error)
	UpdateStock(ctx context.Context, req *product.UpdateStockRequest, callOptions ...callopt.Option) (r *product.UpdateStockResponse, err error)
	GetProductListInfoByUser(ctx context.Context, req *product.GetProductListInfoByUserRequest, callOptions ...callopt.Option) (r *product.GetProductListInfoByUserResponse, err error)
	UpdateStockAndSales(ctx context.Context, req *product.UpdateStockAndSalesRequest, callOptions ...callopt.Option) (r *product.UpdateStockAndSalesResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductServiceClient struct {
	*kClient
}

func (p *kProductServiceClient) AddProduct(ctx context.Context, req *product.AddProductRequest, callOptions ...callopt.Option) (r *product.AddProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddProduct(ctx, req)
}

func (p *kProductServiceClient) GetProductInfo(ctx context.Context, req *product.GetProductInfoRequest, callOptions ...callopt.Option) (r *product.GetProductInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProductInfo(ctx, req)
}

func (p *kProductServiceClient) GetProductListInfo(ctx context.Context, callOptions ...callopt.Option) (r *product.GetProductListInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProductListInfo(ctx)
}

func (p *kProductServiceClient) DelProduct(ctx context.Context, req *product.DelProductRequest, callOptions ...callopt.Option) (r *product.DelProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DelProduct(ctx, req)
}

func (p *kProductServiceClient) UpdatePrice(ctx context.Context, req *product.UpdatePriceRequest, callOptions ...callopt.Option) (r *product.UpdatePriceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePrice(ctx, req)
}

func (p *kProductServiceClient) UpdateStock(ctx context.Context, req *product.UpdateStockRequest, callOptions ...callopt.Option) (r *product.UpdateStockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateStock(ctx, req)
}

func (p *kProductServiceClient) GetProductListInfoByUser(ctx context.Context, req *product.GetProductListInfoByUserRequest, callOptions ...callopt.Option) (r *product.GetProductListInfoByUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProductListInfoByUser(ctx, req)
}

func (p *kProductServiceClient) UpdateStockAndSales(ctx context.Context, req *product.UpdateStockAndSalesRequest, callOptions ...callopt.Option) (r *product.UpdateStockAndSalesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateStockAndSales(ctx, req)
}
