// Code generated by Kitex v0.7.3. DO NOT EDIT.

package userservice

import (
	user "Eshop/kitex_gen/user"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserLogin":            kitex.NewMethodInfo(userLoginHandler, newUserServiceUserLoginArgs, newUserServiceUserLoginResult, false),
		"UserRegiter":          kitex.NewMethodInfo(userRegiterHandler, newUserServiceUserRegiterArgs, newUserServiceUserRegiterResult, false),
		"GetUserInfo":          kitex.NewMethodInfo(getUserInfoHandler, newUserServiceGetUserInfoArgs, newUserServiceGetUserInfoResult, false),
		"GetUserInfoByUserID":  kitex.NewMethodInfo(getUserInfoByUserIDHandler, newUserServiceGetUserInfoByUserIDArgs, newUserServiceGetUserInfoByUserIDResult, false),
		"UpdateName":           kitex.NewMethodInfo(updateNameHandler, newUserServiceUpdateNameArgs, newUserServiceUpdateNameResult, false),
		"UpdatePassword":       kitex.NewMethodInfo(updatePasswordHandler, newUserServiceUpdatePasswordArgs, newUserServiceUpdatePasswordResult, false),
		"UpdateCost":           kitex.NewMethodInfo(updateCostHandler, newUserServiceUpdateCostArgs, newUserServiceUpdateCostResult, false),
		"UpdateBalance":        kitex.NewMethodInfo(updateBalanceHandler, newUserServiceUpdateBalanceArgs, newUserServiceUpdateBalanceResult, false),
		"UpdateBalanceAndCost": kitex.NewMethodInfo(updateBalanceAndCostHandler, newUserServiceUpdateBalanceAndCostArgs, newUserServiceUpdateBalanceAndCostResult, false),
		"UpdateAddress":        kitex.NewMethodInfo(updateAddressHandler, newUserServiceUpdateAddressArgs, newUserServiceUpdateAddressResult, false),
		"GetFriendList":        kitex.NewMethodInfo(getFriendListHandler, newUserServiceGetFriendListArgs, newUserServiceGetFriendListResult, false),
		"AddFriend":            kitex.NewMethodInfo(addFriendHandler, newUserServiceAddFriendArgs, newUserServiceAddFriendResult, false),
		"DeleteFriend":         kitex.NewMethodInfo(deleteFriendHandler, newUserServiceDeleteFriendArgs, newUserServiceDeleteFriendResult, false),
		"GetMessageList":       kitex.NewMethodInfo(getMessageListHandler, newUserServiceGetMessageListArgs, newUserServiceGetMessageListResult, false),
		"SendMessage":          kitex.NewMethodInfo(sendMessageHandler, newUserServiceSendMessageArgs, newUserServiceSendMessageResult, false),
		"GetUserListByContent": kitex.NewMethodInfo(getUserListByContentHandler, newUserServiceGetUserListByContentArgs, newUserServiceGetUserListByContentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": `idl\user.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserLoginArgs)
	realResult := result.(*user.UserServiceUserLoginResult)
	success, err := handler.(user.UserService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserLoginArgs() interface{} {
	return user.NewUserServiceUserLoginArgs()
}

func newUserServiceUserLoginResult() interface{} {
	return user.NewUserServiceUserLoginResult()
}

func userRegiterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserRegiterArgs)
	realResult := result.(*user.UserServiceUserRegiterResult)
	success, err := handler.(user.UserService).UserRegiter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserRegiterArgs() interface{} {
	return user.NewUserServiceUserRegiterArgs()
}

func newUserServiceUserRegiterResult() interface{} {
	return user.NewUserServiceUserRegiterResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func getUserInfoByUserIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoByUserIDArgs)
	realResult := result.(*user.UserServiceGetUserInfoByUserIDResult)
	success, err := handler.(user.UserService).GetUserInfoByUserID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoByUserIDArgs() interface{} {
	return user.NewUserServiceGetUserInfoByUserIDArgs()
}

func newUserServiceGetUserInfoByUserIDResult() interface{} {
	return user.NewUserServiceGetUserInfoByUserIDResult()
}

func updateNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateNameArgs)
	realResult := result.(*user.UserServiceUpdateNameResult)
	success, err := handler.(user.UserService).UpdateName(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateNameArgs() interface{} {
	return user.NewUserServiceUpdateNameArgs()
}

func newUserServiceUpdateNameResult() interface{} {
	return user.NewUserServiceUpdateNameResult()
}

func updatePasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdatePasswordArgs)
	realResult := result.(*user.UserServiceUpdatePasswordResult)
	success, err := handler.(user.UserService).UpdatePassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdatePasswordArgs() interface{} {
	return user.NewUserServiceUpdatePasswordArgs()
}

func newUserServiceUpdatePasswordResult() interface{} {
	return user.NewUserServiceUpdatePasswordResult()
}

func updateCostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateCostArgs)
	realResult := result.(*user.UserServiceUpdateCostResult)
	success, err := handler.(user.UserService).UpdateCost(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateCostArgs() interface{} {
	return user.NewUserServiceUpdateCostArgs()
}

func newUserServiceUpdateCostResult() interface{} {
	return user.NewUserServiceUpdateCostResult()
}

func updateBalanceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateBalanceArgs)
	realResult := result.(*user.UserServiceUpdateBalanceResult)
	success, err := handler.(user.UserService).UpdateBalance(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateBalanceArgs() interface{} {
	return user.NewUserServiceUpdateBalanceArgs()
}

func newUserServiceUpdateBalanceResult() interface{} {
	return user.NewUserServiceUpdateBalanceResult()
}

func updateBalanceAndCostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateBalanceAndCostArgs)
	realResult := result.(*user.UserServiceUpdateBalanceAndCostResult)
	success, err := handler.(user.UserService).UpdateBalanceAndCost(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateBalanceAndCostArgs() interface{} {
	return user.NewUserServiceUpdateBalanceAndCostArgs()
}

func newUserServiceUpdateBalanceAndCostResult() interface{} {
	return user.NewUserServiceUpdateBalanceAndCostResult()
}

func updateAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateAddressArgs)
	realResult := result.(*user.UserServiceUpdateAddressResult)
	success, err := handler.(user.UserService).UpdateAddress(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateAddressArgs() interface{} {
	return user.NewUserServiceUpdateAddressArgs()
}

func newUserServiceUpdateAddressResult() interface{} {
	return user.NewUserServiceUpdateAddressResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFriendListArgs)
	realResult := result.(*user.UserServiceGetFriendListResult)
	success, err := handler.(user.UserService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFriendListArgs() interface{} {
	return user.NewUserServiceGetFriendListArgs()
}

func newUserServiceGetFriendListResult() interface{} {
	return user.NewUserServiceGetFriendListResult()
}

func addFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAddFriendArgs)
	realResult := result.(*user.UserServiceAddFriendResult)
	success, err := handler.(user.UserService).AddFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAddFriendArgs() interface{} {
	return user.NewUserServiceAddFriendArgs()
}

func newUserServiceAddFriendResult() interface{} {
	return user.NewUserServiceAddFriendResult()
}

func deleteFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceDeleteFriendArgs)
	realResult := result.(*user.UserServiceDeleteFriendResult)
	success, err := handler.(user.UserService).DeleteFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceDeleteFriendArgs() interface{} {
	return user.NewUserServiceDeleteFriendArgs()
}

func newUserServiceDeleteFriendResult() interface{} {
	return user.NewUserServiceDeleteFriendResult()
}

func getMessageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetMessageListArgs)
	realResult := result.(*user.UserServiceGetMessageListResult)
	success, err := handler.(user.UserService).GetMessageList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetMessageListArgs() interface{} {
	return user.NewUserServiceGetMessageListArgs()
}

func newUserServiceGetMessageListResult() interface{} {
	return user.NewUserServiceGetMessageListResult()
}

func sendMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceSendMessageArgs)
	realResult := result.(*user.UserServiceSendMessageResult)
	success, err := handler.(user.UserService).SendMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceSendMessageArgs() interface{} {
	return user.NewUserServiceSendMessageArgs()
}

func newUserServiceSendMessageResult() interface{} {
	return user.NewUserServiceSendMessageResult()
}

func getUserListByContentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserListByContentArgs)
	realResult := result.(*user.UserServiceGetUserListByContentResult)
	success, err := handler.(user.UserService).GetUserListByContent(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserListByContentArgs() interface{} {
	return user.NewUserServiceGetUserListByContentArgs()
}

func newUserServiceGetUserListByContentResult() interface{} {
	return user.NewUserServiceGetUserListByContentResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserLogin(ctx context.Context, req *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args user.UserServiceUserLoginArgs
	_args.Req = req
	var _result user.UserServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserRegiter(ctx context.Context, req *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args user.UserServiceUserRegiterArgs
	_args.Req = req
	var _result user.UserServiceUserRegiterResult
	if err = p.c.Call(ctx, "UserRegiter", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (r *user.GetUserInfoResponse, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.Req = req
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfoByUserID(ctx context.Context, req *user.GetUserInfoByUserIDRequest) (r *user.GetUserInfoByUserIDResponse, err error) {
	var _args user.UserServiceGetUserInfoByUserIDArgs
	_args.Req = req
	var _result user.UserServiceGetUserInfoByUserIDResult
	if err = p.c.Call(ctx, "GetUserInfoByUserID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateName(ctx context.Context, req *user.UpdateNameRequest) (r *user.UpdateNameResponse, err error) {
	var _args user.UserServiceUpdateNameArgs
	_args.Req = req
	var _result user.UserServiceUpdateNameResult
	if err = p.c.Call(ctx, "UpdateName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (r *user.UpdatePasswordResponse, err error) {
	var _args user.UserServiceUpdatePasswordArgs
	_args.Req = req
	var _result user.UserServiceUpdatePasswordResult
	if err = p.c.Call(ctx, "UpdatePassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateCost(ctx context.Context, req *user.UpdateCostRequest) (r *user.UpdateCostResponse, err error) {
	var _args user.UserServiceUpdateCostArgs
	_args.Req = req
	var _result user.UserServiceUpdateCostResult
	if err = p.c.Call(ctx, "UpdateCost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest) (r *user.UpdateBalanceResponse, err error) {
	var _args user.UserServiceUpdateBalanceArgs
	_args.Req = req
	var _result user.UserServiceUpdateBalanceResult
	if err = p.c.Call(ctx, "UpdateBalance", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBalanceAndCost(ctx context.Context, req *user.UpdateBalanceAndCostRequest) (r *user.UpdateBalanceAndCostResponse, err error) {
	var _args user.UserServiceUpdateBalanceAndCostArgs
	_args.Req = req
	var _result user.UserServiceUpdateBalanceAndCostResult
	if err = p.c.Call(ctx, "UpdateBalanceAndCost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAddress(ctx context.Context, req *user.UpdateAddressRequest) (r *user.UpdateAddressResponse, err error) {
	var _args user.UserServiceUpdateAddressArgs
	_args.Req = req
	var _result user.UserServiceUpdateAddressResult
	if err = p.c.Call(ctx, "UpdateAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *user.GetFriendListRequest) (r *user.GetFriendListResponse, err error) {
	var _args user.UserServiceGetFriendListArgs
	_args.Req = req
	var _result user.UserServiceGetFriendListResult
	if err = p.c.Call(ctx, "GetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddFriend(ctx context.Context, req *user.AddFriendRequest) (r *user.AddFriendResponse, err error) {
	var _args user.UserServiceAddFriendArgs
	_args.Req = req
	var _result user.UserServiceAddFriendResult
	if err = p.c.Call(ctx, "AddFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest) (r *user.DeleteFriendResponse, err error) {
	var _args user.UserServiceDeleteFriendArgs
	_args.Req = req
	var _result user.UserServiceDeleteFriendResult
	if err = p.c.Call(ctx, "DeleteFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMessageList(ctx context.Context, req *user.GetMessageListRequest) (r *user.GetMessageListResponse, err error) {
	var _args user.UserServiceGetMessageListArgs
	_args.Req = req
	var _result user.UserServiceGetMessageListResult
	if err = p.c.Call(ctx, "GetMessageList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendMessage(ctx context.Context, req *user.SendMessageRequest) (r *user.SendMessageResponse, err error) {
	var _args user.UserServiceSendMessageArgs
	_args.Req = req
	var _result user.UserServiceSendMessageResult
	if err = p.c.Call(ctx, "SendMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserListByContent(ctx context.Context, req *user.GetUserListByContentRequest) (r *user.GetUserListByContentResponse, err error) {
	var _args user.UserServiceGetUserListByContentArgs
	_args.Req = req
	var _result user.UserServiceGetUserListByContentResult
	if err = p.c.Call(ctx, "GetUserListByContent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
