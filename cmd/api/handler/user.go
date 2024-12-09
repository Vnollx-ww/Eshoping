package handler

import (
	"Eshop/cmd/api/rpc"
	"Eshop/dal/rs"
	"Eshop/kitex_gen/base"
	"Eshop/kitex_gen/user"
	captcha2 "Eshop/pkg/captcha"
	"Eshop/pkg/minio"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func AdminLogin(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Username string
		Password string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
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
		Captcha  string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	key := "captcha:" + c.ClientIP()
	captcha, err := rs.GetKeyValue(ctx, key)
	if err != nil {
		logger.Error("验证码已过期，请刷新重试", zap.Error(err))
		BadBaseResponse(c, "验证码已过期")
	}
	if captcha != reqbody.Captcha {
		BadBaseResponse(c, "验证码错误，请重试")
		return
	}
	req := &user.UserLoginRequest{
		Username: reqbody.Username,
		Password: reqbody.Password,
		Captcha:  reqbody.Captcha,
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
		Username string `form:"username"`
		Password string `form:"password"`
		Address  string `form:"address"`
		Captcha  string `form:"captcha"`
	}

	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	file, err := c.FormFile("avatar")
	if err != nil {
		logger.Error("文件上传错误", zap.Error(err))
		BadBaseResponse(c, "头像上传失败")
		return
	}
	fileURL := fmt.Sprintf("http://localhost:9000/user/UserName=%s.jpg", reqbody.Username)
	key := "captcha:" + c.ClientIP()
	captcha, err := rs.GetKeyValue(ctx, key)
	if err != nil {
		logger.Error("验证码已过期，请刷新重试", zap.Error(err))
		BadBaseResponse(c, "验证码已过期")
	}
	if captcha != reqbody.Captcha {
		BadBaseResponse(c, "验证码错误，请重试")
		return
	}
	req := &user.UserRegisterRequest{
		Username: reqbody.Username,
		Password: reqbody.Password,
		Address:  reqbody.Address,
		Captcha:  reqbody.Captcha,
		Avatar:   fileURL,
	}
	res, _ := rpc.Register(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	err = minio.UserUploadFileToMinio(ctx, file, reqbody.Username) // 传递用户 ID，这里假设为 12345
	if err != nil {
		logger.Error("头像上传到 MinIO 失败", zap.Error(err))
		BadBaseResponse(c, "头像上传到 MinIO 失败")
		return
	}
	c.JSON(http.StatusOK, user.UserLoginResponse{
		UserId:     res.UserId,
		Token:      res.Token,
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
	})
}

func GetCaptcha(ctx context.Context, c *app.RequestContext) {
	captcha := captcha2.GenerateCaptcha(c)
	key := "captcha:" + c.ClientIP()
	err := rs.SetKey(ctx, key, captcha, 2*time.Minute)
	if err != nil {
		logger.Error("验证码存入redis失败", zap.Error(err))
	}
}
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token string `json:"token"`
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
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
		logger.Error("前后端数据绑定错误", zap.Error(err))
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
func UpdateAddress(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token   string `json:"token"`
		Address string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.UpdateAddressRequest{
		Token:   reqbody.Token,
		Address: reqbody.Address,
	}
	res, _ := rpc.UpdateAddress(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.UpdateAddressResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func UpdateAvatar(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token string `form:"token"`
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	file, err := c.FormFile("avatar")
	if err != nil {
		logger.Error("文件上传错误", zap.Error(err))
		BadBaseResponse(c, "头像上传失败")
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
	err = minio.UserUploadFileToMinio(ctx, file, res.User.Name) // 传递用户 ID，这里假设为 12345
	if err != nil {
		logger.Error("头像上传到 MinIO 失败", zap.Error(err))
		BadBaseResponse(c, "头像上传到 MinIO 失败")
		return
	}
	c.JSON(http.StatusOK, base.Base{
		StatusCode: 200,
		StatusMsg:  "头像修改成功！",
	})
}
func GetUserInfoByUserID(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		UserId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.GetUserInfoByUserIDRequest{
		UserId: reqbody.UserId,
	}
	res, _ := rpc.GetUserInfoByUserID(ctx, req)
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
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.GetFriendListRequest{
		Token: reqbody.Token,
	}
	res, _ := rpc.GetFriendList(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.GetFriendListResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Friend:     res.Friend,
	})
}
func AddFriend(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token    string
		ToUserId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.AddFriendRequest{
		Token:    reqbody.Token,
		TouserId: reqbody.ToUserId,
	}
	res, _ := rpc.AddFriend(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.AddFriendResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func DelFriend(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token    string
		ToUserId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.DeleteFriendRequest{
		Token:    reqbody.Token,
		TouserId: reqbody.ToUserId,
	}
	res, _ := rpc.DeleteFriend(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.DeleteFriendResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func GetMessageList(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token    string
		ToUserId int64
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.GetMessageListRequest{
		Token:    reqbody.Token,
		TouserId: reqbody.ToUserId,
	}
	res, _ := rpc.GetMessageList(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.GetMessageListResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Message:    res.Message,
	})
}
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Token    string
		ToUserId int64
		Content  string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.SendMessageRequest{
		Token:    reqbody.Token,
		TouserId: reqbody.ToUserId,
		Content:  reqbody.Content,
	}
	res, _ := rpc.SendMessage(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.SendMessageResponse{
		StatusMsg:  res.StatusMsg,
		StatusCode: res.StatusCode,
		Succed:     true,
	})
}
func GetUserListByContent(ctx context.Context, c *app.RequestContext) {
	var reqbody struct {
		Content string
	}
	if err := c.Bind(&reqbody); err != nil {
		logger.Error("前后端数据绑定错误", zap.Error(err))
		BadBaseResponse(c, "无效的请求格式")
		return
	}
	req := &user.GetUserListByContentRequest{
		Content: reqbody.Content,
	}
	res, _ := rpc.GetUserListByContent(ctx, req)
	if res.StatusCode == -1 {
		BadBaseResponse(c, res.StatusMsg)
		return
	}
	c.JSON(http.StatusOK, user.GetUserListByContentResponse{
		StatusMsg:    res.StatusMsg,
		StatusCode:   res.StatusCode,
		Userinfolist: res.Userinfolist,
	})
}
