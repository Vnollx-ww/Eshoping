// Code generated by Kitex v0.7.3. DO NOT EDIT.

package userservice

import (
	user "Eshop/kitex_gen/user"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error)
	UserRegiter(ctx context.Context, req *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error)
	GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error)
	GetUserInfoByUserID(ctx context.Context, req *user.GetUserInfoByUserIDRequest, callOptions ...callopt.Option) (r *user.GetUserInfoByUserIDResponse, err error)
	UpdateName(ctx context.Context, req *user.UpdateNameRequest, callOptions ...callopt.Option) (r *user.UpdateNameResponse, err error)
	UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, callOptions ...callopt.Option) (r *user.UpdatePasswordResponse, err error)
	UpdateCost(ctx context.Context, req *user.UpdateCostRequest, callOptions ...callopt.Option) (r *user.UpdateCostResponse, err error)
	UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest, callOptions ...callopt.Option) (r *user.UpdateBalanceResponse, err error)
	UpdateBalanceAndCost(ctx context.Context, req *user.UpdateBalanceAndCostRequest, callOptions ...callopt.Option) (r *user.UpdateBalanceAndCostResponse, err error)
	UpdateAddress(ctx context.Context, req *user.UpdateAddressRequest, callOptions ...callopt.Option) (r *user.UpdateAddressResponse, err error)
	GetFriendList(ctx context.Context, req *user.GetFriendListRequest, callOptions ...callopt.Option) (r *user.GetFriendListResponse, err error)
	AddFriend(ctx context.Context, req *user.AddFriendRequest, callOptions ...callopt.Option) (r *user.AddFriendResponse, err error)
	DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest, callOptions ...callopt.Option) (r *user.DeleteFriendResponse, err error)
	GetMessageList(ctx context.Context, req *user.GetMessageListRequest, callOptions ...callopt.Option) (r *user.GetMessageListResponse, err error)
	SendMessage(ctx context.Context, req *user.SendMessageRequest, callOptions ...callopt.Option) (r *user.SendMessageResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, req)
}

func (p *kUserServiceClient) UserRegiter(ctx context.Context, req *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegiter(ctx, req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}

func (p *kUserServiceClient) GetUserInfoByUserID(ctx context.Context, req *user.GetUserInfoByUserIDRequest, callOptions ...callopt.Option) (r *user.GetUserInfoByUserIDResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfoByUserID(ctx, req)
}

func (p *kUserServiceClient) UpdateName(ctx context.Context, req *user.UpdateNameRequest, callOptions ...callopt.Option) (r *user.UpdateNameResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateName(ctx, req)
}

func (p *kUserServiceClient) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, callOptions ...callopt.Option) (r *user.UpdatePasswordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePassword(ctx, req)
}

func (p *kUserServiceClient) UpdateCost(ctx context.Context, req *user.UpdateCostRequest, callOptions ...callopt.Option) (r *user.UpdateCostResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCost(ctx, req)
}

func (p *kUserServiceClient) UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest, callOptions ...callopt.Option) (r *user.UpdateBalanceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateBalance(ctx, req)
}

func (p *kUserServiceClient) UpdateBalanceAndCost(ctx context.Context, req *user.UpdateBalanceAndCostRequest, callOptions ...callopt.Option) (r *user.UpdateBalanceAndCostResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateBalanceAndCost(ctx, req)
}

func (p *kUserServiceClient) UpdateAddress(ctx context.Context, req *user.UpdateAddressRequest, callOptions ...callopt.Option) (r *user.UpdateAddressResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateAddress(ctx, req)
}

func (p *kUserServiceClient) GetFriendList(ctx context.Context, req *user.GetFriendListRequest, callOptions ...callopt.Option) (r *user.GetFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, req)
}

func (p *kUserServiceClient) AddFriend(ctx context.Context, req *user.AddFriendRequest, callOptions ...callopt.Option) (r *user.AddFriendResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFriend(ctx, req)
}

func (p *kUserServiceClient) DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest, callOptions ...callopt.Option) (r *user.DeleteFriendResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFriend(ctx, req)
}

func (p *kUserServiceClient) GetMessageList(ctx context.Context, req *user.GetMessageListRequest, callOptions ...callopt.Option) (r *user.GetMessageListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMessageList(ctx, req)
}

func (p *kUserServiceClient) SendMessage(ctx context.Context, req *user.SendMessageRequest, callOptions ...callopt.Option) (r *user.SendMessageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendMessage(ctx, req)
}
