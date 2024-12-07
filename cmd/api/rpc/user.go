package rpc

import (
	"Eshop/kitex_gen/user"
	"Eshop/kitex_gen/user/userservice"
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
	userClient userservice.Client
)

func InitUser(config *viper.Config) {
	etcdAddr := fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	serviceName := config.Viper.GetString("server.name")
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		serviceName,
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}
func GetUserClient() userservice.Client {
	return userClient
}
func Register(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	return userClient.UserRegiter(ctx, req)
}
func Login(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	return userClient.UserLogin(ctx, req)
}
func GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	return userClient.GetUserInfo(ctx, req)
}
func UpdateName(ctx context.Context, req *user.UpdateNameRequest) (*user.UpdateNameResponse, error) {
	return userClient.UpdateName(ctx, req)
}
func UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (*user.UpdatePasswordResponse, error) {
	return userClient.UpdatePassword(ctx, req)
}
func UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest) (*user.UpdateBalanceResponse, error) {
	return userClient.UpdateBalance(ctx, req)
}
func UpdateCost(ctx context.Context, req *user.UpdateCostRequest) (*user.UpdateCostResponse, error) {
	return userClient.UpdateCost(ctx, req)
}
func UpdateAddress(ctx context.Context, req *user.UpdateAddressRequest) (*user.UpdateAddressResponse, error) {
	return userClient.UpdateAddress(ctx, req)
}
func UpdateBalanceAndCost(ctx context.Context, req *user.UpdateBalanceAndCostRequest) (*user.UpdateBalanceAndCostResponse, error) {
	return userClient.UpdateBalanceAndCost(ctx, req)
}
func GetUserInfoByUserID(ctx context.Context, req *user.GetUserInfoByUserIDRequest) (*user.GetUserInfoByUserIDResponse, error) {
	return userClient.GetUserInfoByUserID(ctx, req)
}
func GetFriendList(ctx context.Context, req *user.GetFriendListRequest) (*user.GetFriendListResponse, error) {
	return userClient.GetFriendList(ctx, req)
}
func AddFriend(ctx context.Context, req *user.AddFriendRequest) (*user.AddFriendResponse, error) {
	return userClient.AddFriend(ctx, req)
}
func DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest) (*user.DeleteFriendResponse, error) {
	return userClient.DeleteFriend(ctx, req)
}
func GetMessageList(ctx context.Context, req *user.GetMessageListRequest) (*user.GetMessageListResponse, error) {
	return userClient.GetMessageList(ctx, req)
}
func SendMessage(ctx context.Context, req *user.SendMessageRequest) (*user.SendMessageResponse, error) {
	return userClient.SendMessage(ctx, req)
}
