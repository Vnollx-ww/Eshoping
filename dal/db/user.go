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
	UserName       string  `gorm:"unique;varchar(40);not null" json:"name,omitempty"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password,omitempty"`
	Balance        int64   `gorm:"default:0" json:"balance,omitempty"`
	Cost           int64   `gorm:"default:0" json:"cost,omitempty"`
	Address        string  `gorm:"varchar(256);not null" json:"address,omitempty"`
	Avatar         string  `gorm:"varchar(256);not null" json:"avatar,omitempty"`
	Friend         []int64 `gorm:"type:json" json:"friend,omitempty"`
	SendMessage    string  `gorm:"type:json" json:"send_message,omitempty"`    // 设置默认值为空 JSON 对象
	ReceiveMessage string  `gorm:"type:json" json:"receive_message,omitempty"` // 设置默认值为空 JSON 对象
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
func AddFriend(ctx context.Context, usr *User, friend int64) error {
	for _, v := range usr.Friend {
		if v == friend {
			return fmt.Errorf("已经是好友了，无需再次添加")
		}
	}
	if usr.Friend == nil {
		usr.Friend = []int64{}
	}
	usr.Friend = append(usr.Friend, friend)
	if err := DB.Save(&usr).Error; err != nil {
		return err
	}
	return nil
}
func DeleteFriend(ctx context.Context, usr *User, friend int64) error {
	err := DB.Exec(`
		UPDATE users
		SET friend = JSON_REMOVE(friend, JSON_UNQUOTE(JSON_SEARCH(friend, 'one', ?)))
		WHERE id = ?`, friend, usr.ID).Error

	return err
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
func ReceiveMessage(ctx context.Context, usr *User, tousr *User, msg string) error {
	var Receivemessage map[int64][][2]string
	err := json.Unmarshal([]byte(tousr.ReceiveMessage), &Receivemessage)
	if err != nil {
		return err
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	messageItem := [2]string{msg, currentTime} // 构建消息项：包含消息和时间
	if _, exists := Receivemessage[int64(usr.ID)]; !exists {
		Receivemessage[int64(usr.ID)] = [][2]string{}
	}
	Receivemessage[int64(usr.ID)] = append(Receivemessage[int64(usr.ID)], messageItem)
	receivemessage, err := json.Marshal(Receivemessage)
	if err != nil {
		return err
	}
	tousr.ReceiveMessage = string(receivemessage)
	err = DB.Save(&tousr).Error
	if err != nil {
		return err
	}
	return nil
}
