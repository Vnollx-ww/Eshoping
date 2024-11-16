package main

import (
	"Eshop/dal/db"
	"Eshop/dal/rs"
	"Eshop/kitex_gen/user"
	"Eshop/pkg/middlerware"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"log"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

func BadLoginResponse(s string) *user.UserLoginResponse {
	return &user.UserLoginResponse{StatusCode: -1, StatusMsg: s}
}
func GoodLoginResponse(s string, token string) *user.UserLoginResponse {
	return &user.UserLoginResponse{StatusCode: 200, StatusMsg: s, Token: token}
}
func BadRegisterResponse(s string) *user.UserRegisterResponse {
	return &user.UserRegisterResponse{StatusCode: -1, StatusMsg: s}
}
func GoodRegisterResponse(s string, ID int64) *user.UserRegisterResponse {
	return &user.UserRegisterResponse{StatusCode: 200, StatusMsg: s, UserId: ID}
}
func BadGetUserInfoResponse(s string) *user.GetUserInfoResponse {
	return &user.GetUserInfoResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetUserInfoResponse(s string, usr *user.User) *user.GetUserInfoResponse {
	return &user.GetUserInfoResponse{StatusCode: 200, StatusMsg: s, User: usr}
}
func BadUpdateNameResponse(s string) *user.UpdateNameResponse {
	return &user.UpdateNameResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateNameResponse(s string, flag bool) *user.UpdateNameResponse {
	return &user.UpdateNameResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdatePasswordResponse(s string) *user.UpdatePasswordResponse {
	return &user.UpdatePasswordResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdatePasswordResponse(s string, flag bool) *user.UpdatePasswordResponse {
	return &user.UpdatePasswordResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdateCostResponse(s string) *user.UpdateCostResponse {
	return &user.UpdateCostResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateCostResponse(s string, flag bool) *user.UpdateCostResponse {
	return &user.UpdateCostResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdateBalanceResponse(s string) *user.UpdateBalanceResponse {
	return &user.UpdateBalanceResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateBalanceResponse(s string, flag bool) *user.UpdateBalanceResponse {
	return &user.UpdateBalanceResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdateAddressResponse(s string) *user.UpdateAddressResponse {
	return &user.UpdateAddressResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateAddressResponse(s string, flag bool) *user.UpdateAddressResponse {
	return &user.UpdateAddressResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		log.Println(err.Error())
		return BadLoginResponse("登录失败；服务器内部错误"), nil
	}
	if usr == nil {
		return BadLoginResponse("用户不存在"), nil
	}
	if req.Password != usr.Password {
		return BadLoginResponse("密码错误"), nil
	}
	token, err := middlerware.NewToken(int64(usr.ID))
	if err != nil {
		log.Println(err.Error())
		return BadLoginResponse("Token生成失败"), nil
	}
	return GoodLoginResponse("欢迎你！ "+req.Username, token), nil
}

// UserRegiter implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegiter(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		log.Println(err)
		return BadRegisterResponse("注册失败：服务器内部错误"), nil
	}
	if usr != nil {
		return BadRegisterResponse("用户已存在"), nil
	}
	usr = &db.User{UserName: req.Username, Password: req.Password, Address: req.Address}
	err = db.CreateUser(ctx, usr)
	if err != nil {
		log.Println(err)
		return BadRegisterResponse("注册失败"), nil
	}
	return GoodRegisterResponse("注册成功", int64(usr.ID)), nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		log.Println(err)
		return BadGetUserInfoResponse("token解析失败"), nil
	}
	rc := rs.GetRedisClient()
	cachekey := fmt.Sprintf("userinfo:%d", mc.UserId)
	cacheuserdata, err := rc.Get(ctx, cachekey).Result()
	if err == nil && cacheuserdata != "" {
		var cacheduser user.User
		err := json.Unmarshal([]byte(cacheuserdata), &cacheduser)
		if err == nil {
			return GoodGetUserInfoResponse("从缓存获取用户信息成功", &cacheduser), nil
		}
		rc.Del(ctx, cachekey)
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadGetUserInfoResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadGetUserInfoResponse("用户不存在"), nil
	}
	u := &user.User{Name: usr.UserName, Id: int64(usr.ID), Balance: usr.Balance, Cost: usr.Cost, Address: usr.Address}
	userdata, err := json.Marshal(u)
	if err == nil {
		rc.Set(ctx, cachekey, userdata, 5*time.Minute)
	}
	return GoodGetUserInfoResponse("获取用户信息成功", u), nil
}

// UpdateName implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateName(ctx context.Context, req *user.UpdateNameRequest) (resp *user.UpdateNameResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		log.Println(err)
		return BadUpdateNameResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadUpdateNameResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdateNameResponse("用户不存在"), nil
	}
	err = db.UpdateName(ctx, usr, req.Newname_)
	if err != nil {
		log.Println(err)
		return BadUpdateNameResponse("修改用户名失败"), nil
	}
	rc := rs.GetRedisClient()
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	rc.Del(ctx, cacheKey) // 删除缓存
	return GoodUpdateNameResponse("欢迎你！"+usr.UserName, true), nil
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (resp *user.UpdatePasswordResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		return BadUpdatePasswordResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadUpdatePasswordResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdatePasswordResponse("用户不存在"), nil
	}
	if usr.Password != req.Oldpassword {
		return BadUpdatePasswordResponse("用户旧密码错误"), nil
	}
	err = db.UpdatePassword(ctx, usr, req.Newpassword_)
	if err != nil {
		log.Println(err)
		return BadUpdatePasswordResponse("修改密码失败"), nil
	}
	return GoodUpdatePasswordResponse("修改密码成功", true), nil
}

// UpdateCost implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateCost(ctx context.Context, req *user.UpdateCostRequest) (resp *user.UpdateCostResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		return BadUpdateCostResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadUpdateCostResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdateCostResponse("用户不存在"), nil
	}
	err = db.UpdateCost(ctx, usr, req.Addcost)
	if err != nil {
		log.Println(err)
		return BadUpdateCostResponse("修改花费失败"), nil
	}
	rc := rs.GetRedisClient()
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	rc.Del(ctx, cacheKey) // 删除缓存
	return GoodUpdateCostResponse("修改花费成功", true), nil
}

// UpdateBalance implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest) (resp *user.UpdateBalanceResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdateBalanceResponse("用户不存在"), nil
	}
	err = db.UpdateBalance(ctx, usr, req.Addbalance)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceResponse("修改余额失败"), nil
	}
	rc := rs.GetRedisClient()
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	rc.Del(ctx, cacheKey) // 删除缓存
	return GoodUpdateBalanceResponse("修改余额成功", true), nil
}

// UpdateAddress implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateAddress(ctx context.Context, req *user.UpdateAddressRequest) (resp *user.UpdateAddressResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		return BadUpdateAddressResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadUpdateAddressResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		log.Println(err)
		return BadUpdateAddressResponse("用户不存在"), nil
	}
	err = db.UpdateAddress(ctx, usr, req.Address)
	if err != nil {
		log.Println(err)
		return BadUpdateAddressResponse("用户收货地址修改失败"), nil
	}
	rc := rs.GetRedisClient()
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	rc.Del(ctx, cacheKey) // 删除缓存
	return GoodUpdateAddressResponse("用户收获地址修改成功", true), nil
}
