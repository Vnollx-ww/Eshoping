package main

import (
	"Eshop/dal/db"
	user "Eshop/kitex_gen/user"
	"Eshop/pkg/middlerware"
	"context"
	"log"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	log.Println(req)
	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		log.Println(err.Error())
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "登录失败：服务器内部错误",
		}
		return res, nil
	}
	if usr == nil {
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		log.Println(res)
		return res, nil
	}
	if req.Password != usr.Password {
		log.Println("用户名或密码错误")
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
		}
		return res, nil
	}
	token, err := middlerware.NewToken(int64(usr.ID))
	if err != nil {
		log.Println(err.Error())
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "token生成失败",
		}
		return res, nil
	}
	res := &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "欢迎你! " + req.Username,
		UserId:     int64(usr.ID),
		Token:      token,
	}
	return res, nil
}

// UserRegiter implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegiter(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	usr, err := db.GetUserByName(ctx, req.Username)
	log.Println(req)
	if err != nil {
		log.Println(err.Error())
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	}
	if usr != nil {
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "用户已存在",
		}
		return res, nil
	}
	usr = &db.User{UserName: req.Username, Password: req.Password}
	log.Println(usr)
	err = db.CreateUser(ctx, usr)
	if err != nil {
		log.Println(err.Error())
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败",
		}
		return res, nil
	}
	res := &user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     int64(usr.ID),
	}
	return res, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		log.Println(err.Error())
		res := &user.GetUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "token解析失败",
		}
		return res, nil
	}
	log.Println(mc)
	usr, err := db.GetUserByID(ctx, mc.UserId)
	log.Println(req)
	if err != nil {
		log.Println(err.Error())
		res := &user.GetUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：服务器内部错误",
		}
		return res, nil
	}
	if usr == nil {
		res := &user.GetUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		return res, nil
	}
	u := &user.User{
		Name:    usr.UserName,
		Id:      int64(usr.ID),
		Balance: usr.Balance,
		Cost:    usr.Cost,
	}
	res := &user.GetUserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		User:       u,
	}
	return res, nil
}

// UpdateName implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateName(ctx context.Context, req *user.UpdateNameRequest) (resp *user.UpdateNameResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		log.Println(err.Error())
		res := &user.UpdateNameResponse{
			StatusCode: -1,
			StatusMsg:  "token解析失败",
		}
		return res, nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	log.Println(req)
	if err != nil {
		log.Println(err.Error())
		res := &user.UpdateNameResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：服务器内部错误",
		}
		return res, nil
	}
	if usr == nil {
		res := &user.UpdateNameResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		return res, nil
	}
	err = db.UpdateName(ctx, usr, req.Newname_)
	if err != nil {
		log.Println(err)
		res := &user.UpdateNameResponse{
			StatusCode: 0,
			StatusMsg:  "修改用户名失败",
			Succed:     false,
		}
		return res, nil
	}
	res := &user.UpdateNameResponse{
		StatusCode: 0,
		StatusMsg:  "欢迎你！" + usr.UserName,
		Succed:     true,
	}
	return res, nil
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (resp *user.UpdatePasswordResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		res := &user.UpdatePasswordResponse{
			StatusCode: -1,
			StatusMsg:  "token解析失败",
		}
		return res, nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	//log.Println(req)
	if err != nil {
		log.Println(err.Error())
		res := &user.UpdatePasswordResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：服务器内部错误",
		}
		return res, nil
	}
	if usr == nil {
		res := &user.UpdatePasswordResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		return res, nil
	}
	if usr.Password != req.Oldpassword {
		res := &user.UpdatePasswordResponse{
			StatusCode: -1,
			StatusMsg:  "用户旧密码错误",
		}
		return res, nil
	}
	err = db.UpdatePassword(ctx, usr, req.Newpassword_)
	if err != nil {
		log.Println(err)
		res := &user.UpdatePasswordResponse{
			StatusCode: 0,
			StatusMsg:  "修改密码失败",
			Succed:     false,
		}
		return res, nil
	}
	res := &user.UpdatePasswordResponse{
		StatusCode: 0,
		StatusMsg:  "修改密码成功",
		Succed:     true,
	}
	return res, nil
}

// UpdateCost implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateCost(ctx context.Context, req *user.UpdateCostRequest) (resp *user.UpdateCostResponse, err error) {
	usr, err := db.GetUserByID(ctx, req.UserId)
	log.Println(req)
	if err != nil {
		log.Println(err.Error())
		res := &user.UpdateCostResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：服务器内部错误",
		}
		return res, nil
	}
	if usr == nil {
		res := &user.UpdateCostResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		return res, nil
	}
	err = db.UpdateCost(ctx, req.UserId, req.Addcost)
	if err != nil {
		log.Println(err)
		res := &user.UpdateCostResponse{
			StatusCode: 0,
			StatusMsg:  "修改花费失败",
			Succed:     false,
		}
		return res, nil
	}
	res := &user.UpdateCostResponse{
		StatusCode: 0,
		StatusMsg:  "修改花费成功",
		Succed:     true,
	}
	return res, nil
}

// UpdateBalance implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest) (resp *user.UpdateBalanceResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		log.Println(err.Error())
		res := &user.UpdateBalanceResponse{
			StatusCode: -1,
			StatusMsg:  "token解析失败",
		}
		return res, nil
	}
	if err != nil {
		log.Println(err.Error())
		res := &user.UpdateBalanceResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：服务器内部错误",
		}
		return res, nil
	}
	if usr == nil {
		res := &user.UpdateBalanceResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
		return res, nil
	}
	err = db.UpdateBalance(ctx, usr, req.Addbalance)
	if err != nil {
		//log.Println(err)
		res := &user.UpdateBalanceResponse{
			StatusCode: -1,
			StatusMsg:  "修改余额失败",
			Succed:     false,
		}
		return res, nil
	}
	res := &user.UpdateBalanceResponse{
		StatusCode: 0,
		StatusMsg:  "修改余额成功",
		Succed:     true,
	}
	return res, nil
}
