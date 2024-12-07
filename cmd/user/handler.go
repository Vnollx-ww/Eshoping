package main

import (
	"Eshop/dal/db"
	"Eshop/dal/rs"
	"Eshop/kitex_gen/user"
	user2 "Eshop/pkg/kafka/producer/user"
	"Eshop/pkg/middlerware"
	"Eshop/pkg/minio"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"go.uber.org/zap"
	"time"
)

const cacheKey string = "friendlist"

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

func BadLoginResponse(s string) *user.UserLoginResponse {
	return &user.UserLoginResponse{StatusCode: -1, StatusMsg: s}
}
func GoodLoginResponse(s string, token string, ID int64) *user.UserLoginResponse {
	return &user.UserLoginResponse{StatusCode: 200, StatusMsg: s, Token: token, UserId: ID}
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
func BadGetUserInfoByUserIDResponse(s string) *user.GetUserInfoByUserIDResponse {
	return &user.GetUserInfoByUserIDResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetUserInfoByUserIDResponse(s string, usr *user.User) *user.GetUserInfoByUserIDResponse {
	return &user.GetUserInfoByUserIDResponse{StatusCode: 200, StatusMsg: s, User: usr}
}
func BadGetFriendListResponse(s string) *user.GetFriendListResponse {
	return &user.GetFriendListResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetFriendListResponse(s string, friend []*user.FriendInfo) *user.GetFriendListResponse {
	return &user.GetFriendListResponse{StatusCode: 200, StatusMsg: s, Friend: friend}
}
func BadAddFriendResponse(s string) *user.AddFriendResponse {
	return &user.AddFriendResponse{StatusCode: -1, StatusMsg: s}
}
func GoodAddFriendResponse(s string, flag bool) *user.AddFriendResponse {
	return &user.AddFriendResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadDeleteFriendResponse(s string) *user.DeleteFriendResponse {
	return &user.DeleteFriendResponse{StatusCode: -1, StatusMsg: s}
}
func GoodDeleteFriendResponse(s string, flag bool) *user.DeleteFriendResponse {
	return &user.DeleteFriendResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadGetMessageListResponse(s string) *user.GetMessageListResponse {
	return &user.GetMessageListResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetMessageListResponse(s string, message []*user.Message) *user.GetMessageListResponse {
	return &user.GetMessageListResponse{StatusCode: 200, StatusMsg: s, Message: message}
}
func BadSendMessageResponse(s string) *user.SendMessageResponse {
	return &user.SendMessageResponse{StatusCode: -1, StatusMsg: s}
}
func GoodSendMessageResponse(s string, flag bool) *user.SendMessageResponse {
	return &user.SendMessageResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		logger.Error("登录失败，服务器内部错误：", zap.Error(err))
		return BadLoginResponse("登录失败；服务器内部错误"), nil
	}
	if usr == nil {
		return BadLoginResponse("用户不存在"), nil
	}
	encryptedPassword := middlerware.MD5Encrypt(req.Password)
	if encryptedPassword != usr.Password {
		return BadLoginResponse("密码错误"), nil
	}

	token, err := middlerware.NewToken(int64(usr.ID))
	if err != nil {
		logger.Error("Token生成失败：", zap.Error(err))
		return BadLoginResponse("Token生成失败"), nil
	}
	return GoodLoginResponse("欢迎你！ "+req.Username, token, int64(usr.ID)), nil
}

// UserRegiter implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegiter(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		logger.Error("注册失败，服务器内部错误：", zap.Error(err))
		return BadRegisterResponse("注册失败：服务器内部错误"), nil
	}
	if usr != nil {
		return BadRegisterResponse("用户已存在"), nil
	}
	encryptedPassword := middlerware.MD5Encrypt(req.Password)
	usr = &db.User{UserName: req.Username, Password: encryptedPassword, Address: req.Address, Avatar: req.Avatar} // 设置为 NULL}
	usr.SendMessage = "{}"
	usr.Friend = "{}"
	err = db.CreateUser(ctx, usr)
	if err != nil {
		logger.Error("注册失败：", zap.Error(err))
		return BadRegisterResponse("注册失败"), nil
	}
	return GoodRegisterResponse("注册成功", int64(usr.ID)), nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadGetUserInfoResponse("token解析失败"), nil
	}
	cachekey := fmt.Sprintf("userinfo:%d", mc.UserId)
	u, err := rs.GetUserInfo(ctx, cachekey)
	if err == nil && u != nil {
		return GoodGetUserInfoResponse("获取用户信息成功", u), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadGetUserInfoResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadGetUserInfoResponse("用户不存在"), nil
	}
	u = &user.User{Name: usr.UserName, Id: int64(usr.ID), Balance: usr.Balance, Cost: usr.Cost, Address: usr.Address, Avatar: usr.Avatar}
	userdata, err := json.Marshal(u)
	if err != nil {
		logger.Error("用户信息数据序列化失败：", zap.Error(err))
	}
	err = rs.SetKey(ctx, cachekey, userdata, 10*time.Minute)
	if err != nil {
		logger.Error("缓存设置失败：", zap.Error(err))
	}
	return GoodGetUserInfoResponse("获取用户信息成功", u), nil
}

// UpdateName implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateName(ctx context.Context, req *user.UpdateNameRequest) (resp *user.UpdateNameResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadUpdateNameResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadUpdateNameResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdateNameResponse("用户不存在"), nil
	}
	err = minio.UserCopyFileMinio(ctx, usr.UserName, req.Newname_)
	if err != nil {
		logger.Error("用户头像路径修改失败：", zap.Error(err))
		return BadUpdateNameResponse("用户头像路径修改失败"), nil
	}
	kafkaProducer, err := user2.NewKafkaProducer([]string{kafkaAddr}) //初始化kafka生产者
	if err != nil {
		logger.Error("kafka生产者创建失败：", zap.Error(err))
		return BadUpdateNameResponse("Kafka生产者创建失败"), err
	}
	err = kafkaProducer.SendDeleteAvatarEvent(usr.UserName)
	if err != nil {
		logger.Error("删除旧路径消息创建成功，但更新消息发送失败：", zap.Error(err))
		return BadUpdateNameResponse("删除旧路径消息创建成功，但更新消息发送失败"), err
	}
	err = db.UpdateNameAndAvatar(ctx, usr, req.Newname_)
	if err != nil {
		logger.Error("修改用户名失败：", zap.Error(err))
		return BadUpdateNameResponse("修改用户名失败"), nil
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateNameResponse("欢迎你！"+usr.UserName, true), nil
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (resp *user.UpdatePasswordResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadUpdatePasswordResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
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
		logger.Error("修改密码失败：", zap.Error(err))
		return BadUpdatePasswordResponse("修改密码失败"), nil
	}
	return GoodUpdatePasswordResponse("修改密码成功", true), nil
}

// UpdateCost implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateCost(ctx context.Context, req *user.UpdateCostRequest) (resp *user.UpdateCostResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadUpdateCostResponse("token解析失败"), err
	}
	lockKey := rs.RedisLockKeyPrefix + string(mc.UserId) + "cost"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		logger.Error("获取Redis锁失败：", zap.Error(err)) //锁已存在
		return BadUpdateCostResponse("获取花费锁失败，稍后重试"), nil
	}
	if !locked {
		return BadUpdateCostResponse("花费操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadUpdateCostResponse("获取用户信息失败：服务器内部错误"), err
	}
	if usr == nil {
		return BadUpdateCostResponse("用户不存在"), nil
	}
	err = db.UpdateCost(ctx, usr, req.Addcost)
	if err != nil {
		logger.Error("修改花费失败：", zap.Error(err))
		return BadUpdateCostResponse("修改花费失败"), err
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateCostResponse("修改花费成功", true), nil
}

// UpdateBalance implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBalance(ctx context.Context, req *user.UpdateBalanceRequest) (resp *user.UpdateBalanceResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadUpdateBalanceResponse("token解析失败"), err
	}
	lockKey := rs.RedisLockKeyPrefix + string(mc.UserId) + "balance"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		logger.Error("获取Redis锁失败：", zap.Error(err)) //锁已存在
		return BadUpdateBalanceResponse("获取余额锁失败，稍后重试"), nil
	}
	if !locked {
		return BadUpdateBalanceResponse("余额操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadUpdateBalanceResponse("获取用户信息失败：服务器内部错误"), err
	}
	if usr == nil {
		return BadUpdateBalanceResponse("用户不存在"), err
	}
	err = db.UpdateBalance(ctx, usr, req.Addbalance)
	if err != nil {
		logger.Error("修改余额失败：", zap.Error(err))
		return BadUpdateBalanceResponse("修改余额失败"), err
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateBalanceResponse("修改余额成功", true), err
}

// UpdateAddress implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateAddress(ctx context.Context, req *user.UpdateAddressRequest) (resp *user.UpdateAddressResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadUpdateAddressResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadUpdateAddressResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdateAddressResponse("用户不存在"), nil
	}
	err = db.UpdateAddress(ctx, usr, req.Address)
	if err != nil {
		logger.Error("用户收货地址修改失败：", zap.Error(err))
		return BadUpdateAddressResponse("用户收货地址修改失败"), nil
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateAddressResponse("用户收获地址修改成功", true), nil
}

// UpdateBalanceAndCost implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBalanceAndCost(ctx context.Context, req *user.UpdateBalanceAndCostRequest) (resp *user.UpdateBalanceAndCostResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadUpdateBalanceAndCostResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadUpdateBalanceAndCostResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadUpdateBalanceAndCostResponse("用户不存在"), nil
	}
	err = db.UpdateBalanceAndCost(ctx, usr, req.Number)
	if err != nil {
		logger.Error("用户余额和花费修改失败：", zap.Error(err))
		return BadUpdateBalanceAndCostResponse("用户余额和花费修改失败"), nil
	}
	cacheKey := fmt.Sprintf("userinfo:%d", mc.UserId)
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateBalanceAndCostResponse("用户余额和花费修改成功", true), nil
}

// GetUserInfoByUserID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfoByUserID(ctx context.Context, req *user.GetUserInfoByUserIDRequest) (resp *user.GetUserInfoByUserIDResponse, err error) {
	usr, err := db.GetUserByID(ctx, req.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadGetUserInfoByUserIDResponse("获取用户信息失败：服务器内部错误"), nil
	}
	if usr == nil {
		return BadGetUserInfoByUserIDResponse("用户不存在"), nil
	}
	u := &user.User{Name: usr.UserName, Id: int64(usr.ID), Balance: usr.Balance, Cost: usr.Cost, Address: usr.Address, Avatar: usr.Avatar}
	return GoodGetUserInfoByUserIDResponse("获取用户信息成功", u), nil
}

// GetFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFriendList(ctx context.Context, req *user.GetFriendListRequest) (resp *user.GetFriendListResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadGetFriendListResponse("token解析失败"), nil
	}
	flist, err := rs.GetFriendList(ctx, cacheKey+fmt.Sprintf("%d", mc.UserId))
	if err == nil && flist != nil {
		return GoodGetFriendListResponse("获取好友列表信息成功", flist), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadGetFriendListResponse("获取用户信息失败：服务器内部错误"), nil
	}
	var frlist map[int64]int
	err = json.Unmarshal([]byte(usr.Friend), &frlist)
	if err != nil {
		logger.Error("反序列化好友列表失败：", zap.Error(err))
		return BadGetFriendListResponse("反序列化好友列表失败"), nil
	}
	var friendlist []*user.FriendInfo
	for userid := range frlist {
		f := &user.FriendInfo{}
		f.Name, err = db.GetUserNameByID(ctx, userid)
		if err != nil {
			return BadGetFriendListResponse(err.Error()), nil
		}
		f.Avatar, err = db.GetAvatarByID(ctx, userid)
		if err != nil {
			return BadGetFriendListResponse(err.Error()), nil
		}
		f.UserId = userid
		friendlist = append(friendlist, f)
	}
	//log.Println(friendlist)
	cacheddatabytes, err := json.Marshal(friendlist)
	if err != nil {
		logger.Error("好友列表数据序列化失败：", zap.Error(err))
	}
	err = rs.SetKey(ctx, cacheKey+fmt.Sprintf("%d", mc.UserId), cacheddatabytes, 10*time.Minute)
	if err != nil {
		logger.Error("缓存设置失败：", zap.Error(err))
	}
	return GoodGetFriendListResponse("获取好友列表信息成功", friendlist), nil
}

// AddFriend implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddFriend(ctx context.Context, req *user.AddFriendRequest) (resp *user.AddFriendResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadAddFriendResponse("token解析失败"), nil
	}
	usra, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取目标用户信息失败，服务器内部错误：", zap.Error(err))
		return BadAddFriendResponse("获取目标用户信息失败：服务器内部错误"), nil
	}
	usrb, err := db.GetUserByID(ctx, req.TouserId)
	if err != nil {
		logger.Error("获取目标用户信息失败，服务器内部错误：", zap.Error(err))
		return BadAddFriendResponse("获取目标用户信息失败：服务器内部错误"), nil
	}
	err = db.AddFriend(ctx, usra, usrb)
	if err != nil {
		logger.Error("好友添加失败：", zap.Error(err))
		return BadAddFriendResponse("好友添加失败"), nil
	}
	err = rs.DelKey(ctx, cacheKey+fmt.Sprintf("%d", mc.UserId)) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodAddFriendResponse("好友添加成功", true), nil
}

// DeleteFriend implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest) (resp *user.DeleteFriendResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadDeleteFriendResponse("token解析失败"), nil
	}
	usra, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取目标用户信息失败，服务器内部错误：", zap.Error(err))
		return BadDeleteFriendResponse("获取目标用户信息失败：服务器内部错误"), nil
	}
	usrb, err := db.GetUserByID(ctx, req.TouserId)
	if err != nil {
		logger.Error("获取目标用户信息失败，服务器内部错误：", zap.Error(err))
		return BadDeleteFriendResponse("获取目标用户信息失败：服务器内部错误"), nil
	}
	err = db.DeleteFriend(ctx, usra, usrb)
	if err != nil {
		logger.Error("好友删除失败：", zap.Error(err))
		return BadDeleteFriendResponse("好友删除失败"), nil
	}
	err = rs.DelKey(ctx, cacheKey+fmt.Sprintf("%d", mc.UserId)) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodDeleteFriendResponse("好友删除成功", true), nil
}

// GetMessageList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetMessageList(ctx context.Context, req *user.GetMessageListRequest) (resp *user.GetMessageListResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadGetMessageListResponse("token解析失败"), nil
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadGetMessageListResponse("获取用户信息失败：服务器内部错误"), nil
	}
	tousr, err := db.GetUserByID(ctx, req.TouserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return BadGetMessageListResponse("获取用户信息失败：服务器内部错误"), nil
	}
	var smsg map[int64][][2]string
	smsg = make(map[int64][][2]string)
	err = json.Unmarshal([]byte(usr.SendMessage), &smsg)
	if err != nil {
		logger.Error("消息记录反序列化失败", zap.Error(err))
		return BadGetMessageListResponse("消息记录反序列化失败"), nil
	}
	var amsg map[int64][][2]string
	amsg = make(map[int64][][2]string)
	err = json.Unmarshal([]byte(tousr.SendMessage), &amsg)
	if err != nil {
		logger.Error("消息记录反序列化失败", zap.Error(err))
		return BadGetMessageListResponse("消息记录反序列化失败"), nil
	}
	var message []*user.Message
	if smsMessages, exists := smsg[int64(tousr.ID)]; exists {
		for _, msg := range smsMessages {
			message = append(message, &user.Message{
				Content:    msg[0],
				CreateTime: msg[1],
				UserId:     int64(usr.ID),
			})
		}
	}
	if amsMessages, exists := amsg[int64(usr.ID)]; exists {
		for _, msg := range amsMessages {
			message = append(message, &user.Message{
				Content:    msg[0],
				CreateTime: msg[1],
				UserId:     int64(tousr.ID),
			})
		}
	}
	return GoodGetMessageListResponse("获取消息记录成功", message), nil
}

// SendMessage implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendMessage(ctx context.Context, req *user.SendMessageRequest) (resp *user.SendMessageResponse, err error) {
	kafkaProducer, err := user2.NewKafkaProducer([]string{kafkaAddr}) //初始化kafka生产者
	if err != nil {
		logger.Error("kafka生产者创建失败：", zap.Error(err))
		return BadSendMessageResponse("Kafka生产者创建失败"), err
	}
	err = kafkaProducer.SendSendMessageEvent(req.Token, req.Content, req.TouserId)
	if err != nil {
		logger.Error("发送消息创建成功，但消息发送失败：", zap.Error(err))
		return BadSendMessageResponse("发送消息创建成功，但消息发送失败"), err
	}
	return GoodSendMessageResponse("发送消息成功", true), nil
}
