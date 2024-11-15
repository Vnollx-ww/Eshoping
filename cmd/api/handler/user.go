package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/base"
	"Eshop/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

func AdminLogin(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Username string
		Password string
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	password := "vnollx"
	username := "root"
	if username == reqbody.Username && password == reqbody.Password {
		c.JSON(http.StatusOK, base.Base{
			StatusCode: 200,
			StatusMsg:  "管理员登录成功",
		})
	}
}
func Login(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Username string
		Password string
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UserLoginRequest{
		Username: reqbody.Username,
		Password: reqbody.Password,
	}
	res, _ := rpc.Login(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UserLoginResponse{
		UserId:     res.UserId,
		Token:      res.Token,
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
	})
}
func Register(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Username string
		Password string
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UserRegisterRequest{
		Username: reqbody.Username,
		Password: reqbody.Password,
	}
	res, _ := rpc.Register(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UserLoginResponse{
		UserId:     res.UserId,
		Token:      res.Token,
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
	})
}
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token string `json:"token"`
	}
	if err := c.Bind(&reqbody); err != nil {
		log.Println(err)
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.GetUserInfoRequest{
		Token: reqbody.Token,
	}
	res, _ := rpc.GetUserInfo(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.GetUserInfoResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		User:       res.User,
	})
}
func UpdateName(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token   string `json:"token"`
		NewName string `json:"newname"`
	}
	if err := c.Bind(&reqbody); err != nil {
		log.Println(err)
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UpdateNameRequest{
		Token:    reqbody.Token,
		Newname_: reqbody.NewName,
	}
	res, _ := rpc.UpdateName(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UpdateNameResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func UpdatePassword(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token   string `json:"token"`
		OldPass string `json:"oldpassword"`
		NewPass string `json:"newpassword"`
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UpdatePasswordRequest{
		Token:        reqbody.Token,
		Oldpassword:  reqbody.OldPass,
		Newpassword_: reqbody.NewPass,
	}
	res, _ := rpc.UpdatePassword(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UpdatePasswordResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func UpdateBalance(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token      string `json:"token"`
		AddBalance int64  `json:"balance"`
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UpdateBalanceRequest{
		Token:      reqbody.Token,
		Addbalance: reqbody.AddBalance,
	}
	res, _ := rpc.UpdateBalance(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UpdateBalanceResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func UpdateCost(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token   string `json:"token"`
		AddCost int64
	}
	if err := c.Bind(&reqbody); err != nil {
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UpdateCostRequest{
		Token:   reqbody.Token,
		Addcost: reqbody.AddCost,
	}
	res, _ := rpc.UpdateCost(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UpdateCostResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
