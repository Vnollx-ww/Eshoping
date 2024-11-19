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
func BadUpdateBalanceAndCostResponse(s string) *user.UpdateBalanceAndCostResponse {
	return &user.UpdateBalanceAndCostResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateBalanceAndCostResponse(s string, flag bool) *user.UpdateBalanceAndCostResponse {
	return &user.UpdateBalanceAndCostResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
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
	cachekey := fmt.Sprintf("userinfo:%d", mc.UserId)
	u, err := rs.GetUserInfo(ctx, cachekey)
	if err == nil && u != nil {
		return GoodGetUserInfoResponse("获取用户信息成功", u), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadGetUserInfoResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadGetUserInfoResponse("用户不存在"), nil
	}
	u = &user.User{Name: usr.UserName, Id: int64(usr.ID), Balance: usr.Balance, Cost: usr.Cost, Address: usr.Address}
	userdata, err := json.Marshal(u)
	if err != nil {
		log.Println("用户信息数据序列化失败:", err)
	}
	err = rs.SetKey(ctx, cachekey, userdata, 10*time.Minute)
	if err != nil {
		log.Println("缓存设置失败:", err)
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
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		log.Println("删除缓存失败:", err)
	}
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
		return BadUpdateCostResponse("token解析失败"), err
	}
	lockKey := rs.RedisLockKeyPrefix + string(mc.UserId) + "cost"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		log.Println("获取 Redis 锁失败:", err) //锁已存在
		return BadUpdateCostResponse("获取花费锁失败，稍后重试"), nil
	}
	if !locked {
		return BadUpdateCostResponse("花费操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	usr, err := db.GetUserByID(ctx, mc.UserId)
	log.Println(usr.UserName)
	if err != nil {
		log.Println(err)
		return BadUpdateCostResponse("获取用户信息失败：服务器内部错误"), err
	}
	if usr == nil {
		return BadUpdateCostResponse("用户不存在"), nil
	}
	err = db.UpdateCost(ctx, usr, req.Addcost)
	if err != nil {
		log.Println(err)
		return BadUpdateCostResponse("修改花费失败"), err
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		log.Println("删除缓存失败:", err)
	}
	return GoodUpdateCostResponse("修改花费成功", true), nil
}

// UpdateBalance implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest) (resp *user.UpdateBalanceResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceResponse("token解析失败"), err
	}
	lockKey := rs.RedisLockKeyPrefix + string(mc.UserId) + "balance"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		log.Println("获取 Redis 锁失败:", err) //锁已存在
		return BadUpdateBalanceResponse("获取余额锁失败，稍后重试"), nil
	}
	if !locked {
		return BadUpdateBalanceResponse("余额操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	usr, err := db.GetUserByID(ctx, mc.UserId)
	log.Println(usr.UserName)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceResponse("获取用户信息失败：服务器内部错误"), err
	}
	if usr == nil {
		return BadUpdateBalanceResponse("用户不存在"), err
	}
	err = db.UpdateBalance(ctx, usr, req.Addbalance)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceResponse("修改余额失败"), err
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		log.Println("删除缓存失败:", err)
	}
	return GoodUpdateBalanceResponse("修改余额成功", true), err
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
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		log.Println("删除缓存失败:", err)
	}
	return GoodUpdateAddressResponse("用户收获地址修改成功", true), nil
}

// UpdateBalanceAndCost implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBalanceAndCost(ctx context.Context, req *user.UpdateBalanceAndCostRequest) (resp *user.UpdateBalanceAndCostResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceAndCostResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceAndCostResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		log.Println(err)
		return BadUpdateBalanceAndCostResponse("用户不存在"), nil
	}
	err = db.UpdateBalanceAndCost(ctx, usr, req.Number)
	if err != nil {
		log.Println(err)
		return BadUpdateBalanceAndCostResponse("用户余额和花费修改失败"), nil
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		log.Println("删除缓存失败:", err)
	}
	return GoodUpdateBalanceAndCostResponse("用户余额和花费修改成功", true), nil
}
