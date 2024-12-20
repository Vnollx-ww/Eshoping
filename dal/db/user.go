package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName              string `gorm:"unique;varchar(40);not null" json:"name,omitempty"`
	Password              string `gorm:"type:varchar(256);not null" json:"password,omitempty"`
	Balance               int64  `gorm:"default:0" json:"balance,omitempty"`
	Cost                  int64  `gorm:"default:0" json:"cost,omitempty"`
	Address               string `gorm:"varchar(256);not null" json:"address,omitempty"`
	Avatar                string `gorm:"varchar(256);not null" json:"avatar,omitempty"`
	Friend                string `gorm:"type:json" json:"friend,omitempty"`
	SendMessage           string `gorm:"type:json" json:"send_message,omitempty"` // 设置默认值为空 JSON 对象
	SendFriendApplication string `gorm:"type:json" json:"send_friend_application,omitempty"`
}

func CreateUser(ctx context.Context, usr *User) error {
	err := DB.Create(usr).Error
	return err
}
func GetUserByName(ctx context.Context, userName string) (*User, error) {
	user := new(User)
	//db.Where("user_name = ?", userName).First(&user)
	if err := DB.Where("user_name = ?", userName).First(&user).Error; err == nil {
		return user, nil
	} else if user.ID == 0 {
		return nil, nil
	} else {
		return nil, err
	}
}
func GetUserNameByID(ctx context.Context, ID int64) (string, error) {
	user := new(User)
	if err := DB.Where("ID = ?", ID).First(&user).Error; err == nil {
		return user.UserName, nil
	} else if user.ID == 0 {
		return "", nil
	} else {
		return "", err
	}
}
func GetAvatarByID(ctx context.Context, ID int64) (string, error) {
	user := new(User)
	if err := DB.Where("ID = ?", ID).First(&user).Error; err == nil {
		return user.Avatar, nil
	} else if user.ID == 0 {
		return "", nil
	} else {
		return "", err
	}
}
func GetUserByID(ctx context.Context, ID int64) (*User, error) {
	user := new(User)
	if err := DB.Where("ID = ?", ID).First(&user).Error; err == nil {
		return user, nil
	} else if user.ID == 0 {
		return nil, nil
	} else {
		return nil, err
	}
}
func UpdateNameAndAvatar(ctx context.Context, user *User, name string) error {
	user.UserName = name
	user.Avatar = fmt.Sprintf("http://localhost:9000/user/UserName=%s.jpg", name)
	err := DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdatePassword(ctx context.Context, usr *User, password string) error {
	usr.Password = password
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBalance(ctx context.Context, usr *User, addbalance int64) error {
	usr.Balance += addbalance
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateCost(ctx context.Context, usr *User, addcost int64) error {
	usr.Cost += addcost
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBalanceAndCost(ctx context.Context, usr *User, addbalance int64) error {
	usr.Balance -= addbalance
	usr.Cost += addbalance
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateAddress(ctx context.Context, usr *User, address string) error {
	usr.Address = address
	err := DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}

func AddFriend(ctx context.Context, usra *User, usrb *User) error {
	// 如果 usra.Friend 或 usrb.Friend 为 nil，则初始化为空 map
	var usrafriend map[int64]int
	var usrbfriend map[int64]int
	err := json.Unmarshal([]byte(usra.Friend), &usrafriend)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(usrb.Friend), &usrbfriend)
	if err != nil {
		return err
	}
	_, exists := usrafriend[int64(usrb.ID)]
	if exists {
		return fmt.Errorf("已经是好友了，无需再次添加")
	}
	// 使用数据库事务，确保原子性
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 添加好友关系
	usrafriend[int64(usrb.ID)] = 1 // 将 usrb 添加到 usra 的好友列表
	usrbfriend[int64(usra.ID)] = 1 // 将 usra 添加到 usrb 的好友列表
	// 将 map 转换为 JSON 格式
	usraJSON, err := json.Marshal(usrafriend)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化 用户A 的 Friend 错误: %v", err)
	}
	usrbJSON, err := json.Marshal(usrbfriend)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化 用户b 的 Friend 错误: %v", err)
	}
	usra.Friend = string(usraJSON)
	usrb.Friend = string(usrbJSON)
	// 保存 usra 和 usrb
	if err := tx.Save(&usra).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&usrb).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
func DeleteFriend(ctx context.Context, usra *User, usrb *User) error {
	// 如果 usra.Friend 或 usrb.Friend 为 nil，则初始化为空 map
	var usrafriend map[int64]int
	var usrbfriend map[int64]int
	err := json.Unmarshal([]byte(usra.Friend), &usrafriend)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(usrb.Friend), &usrbfriend)
	if err != nil {
		return err
	}

	// 检查是否是好友，如果不是，则直接返回
	_, exists := usrafriend[int64(usrb.ID)]
	if !exists {
		return fmt.Errorf("已经不再是好友了，无需再次删除")
	}

	// 使用数据库事务，确保原子性
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除好友关系
	delete(usrafriend, int64(usrb.ID)) // 从 usra 的好友列表中删除 usrb
	delete(usrbfriend, int64(usra.ID)) // 从 usrb 的好友列表中删除 usra

	// 更新好友字段
	usraJSON, err := json.Marshal(usrafriend)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化 用户A 的 Friend 错误: %v", err)
	}
	usrbJSON, err := json.Marshal(usrbfriend)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化 用户B 的 Friend 错误: %v", err)
	}
	usra.Friend = string(usraJSON)
	usrb.Friend = string(usrbJSON)

	// 删除 SendMessage 中与对方的聊天记录
	var sendMessageA map[int64][][2]string
	var sendMessageB map[int64][][2]string
	err = json.Unmarshal([]byte(usra.SendMessage), &sendMessageA)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("反序列化 用户A 的 SendMessage 错误: %v", err)
	}
	err = json.Unmarshal([]byte(usrb.SendMessage), &sendMessageB)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("反序列化 用户B 的 SendMessage 错误: %v", err)
	}

	// 删除与对方的聊天记录
	delete(sendMessageA, int64(usrb.ID)) // 删除与 usrb 的聊天记录
	delete(sendMessageB, int64(usra.ID)) // 删除与 usra 的聊天记录

	// 更新 SendMessage 字段
	usraSendMessageJSON, err := json.Marshal(sendMessageA)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化 用户A 的 SendMessage 错误: %v", err)
	}
	usrbSendMessageJSON, err := json.Marshal(sendMessageB)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化 用户B 的 SendMessage 错误: %v", err)
	}
	usra.SendMessage = string(usraSendMessageJSON)
	usrb.SendMessage = string(usrbSendMessageJSON)

	// 保存更新后的数据
	if err := tx.Save(&usra).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&usrb).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}
func SendMessage(ctx context.Context, usr *User, tousr *User, msg string) error {
	var Sendmessage map[int64][][2]string
	err := json.Unmarshal([]byte(usr.SendMessage), &Sendmessage)
	if err != nil {
		return err
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	messageItem := [2]string{msg, currentTime} // 构建消息项：包含消息和时间
	//发送方添加发送消息到数据库
	if _, exists := Sendmessage[int64(tousr.ID)]; !exists {
		Sendmessage[int64(tousr.ID)] = [][2]string{} // 如果 key 对应的值不存在，初始化一个空的切片
	}
	Sendmessage[int64(tousr.ID)] = append(Sendmessage[int64(tousr.ID)], messageItem) // 将新的消息项添加到该 target 的数组中
	sendmessage, err := json.Marshal(Sendmessage)                                    // 序列化更新后的 SendMessage 数据并存储回 usr.SendMessage
	if err != nil {
		return err
	}
	usr.SendMessage = string(sendmessage)
	err = DB.Save(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func GetUserListByContent(ctx context.Context, content string) ([]*User, error) {
	var u []*User
	err := DB.Where("user_name LIKE ?", "%"+content+"%").Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
func SendFriendApplication(ctx context.Context, usr *User, tousr *User) error {
	// 1. 反序列化 SendFriendApplication 字段为 map[int64]bool
	var sendFriendApplicationMap map[int64]bool
	err := json.Unmarshal([]byte(tousr.SendFriendApplication), &sendFriendApplicationMap)
	if err != nil {
		return err
	}
	_, ok := sendFriendApplicationMap[int64(usr.ID)]
	if ok {
		delete(sendFriendApplicationMap, int64(usr.ID))
	}
	// 2. 将当前用户的请求状态设置为 true，表示已发送请求
	sendFriendApplicationMap[int64(usr.ID)] = true
	// 3. 将 map 序列化回 JSON 字符串
	updatedData, err := json.Marshal(sendFriendApplicationMap)
	if err != nil {
		return err
	}
	// 4. 更新 usr 的 SendFriendApplication 字段
	tousr.SendFriendApplication = string(updatedData)
	// 假设有一个 `db.Save` 方法来更新数据库
	if err := DB.Save(tousr).Error; err != nil {
		return err
	}
	return nil
}
func DeleteFriendApplication(ctx context.Context, usr *User, tousr *User) error {
	var sendFriendApplication map[int64]bool
	err := json.Unmarshal([]byte(tousr.SendFriendApplication), &sendFriendApplication)
	if err != nil {
		return err
	}
	// 删除 userb.ID 这一键
	delete(sendFriendApplication, int64(usr.ID))
	// 将更新后的 map 序列化回 JSON
	updatedSendFriendApplicationJSON, err := json.Marshal(sendFriendApplication)
	if err != nil {
		return err
	}
	// 更新 usra.SendFriendApplication
	tousr.SendFriendApplication = string(updatedSendFriendApplicationJSON)
	err = DB.Save(tousr).Error
	if err != nil {
		return err
	}
	return nil
}
func IsFriendJudge(ctx context.Context, usr *User, ID int64) (bool, error) {
	var usrfriend map[int64]int
	err := json.Unmarshal([]byte(usr.Friend), &usrfriend)
	if err != nil {
		return false, err
	}
	_, ok := usrfriend[ID]
	if ok {
		return true, nil
	} else {
		return false, nil
	}
}
