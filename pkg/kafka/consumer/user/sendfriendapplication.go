package user

import (
	"Eshop/dal/db"
	"Eshop/pkg/zaplog"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"log"
	"time"
)

var logger *zap.Logger = zaplog.GetLogger()

type SendFriendApplicationConsumer struct {
	consumer sarama.Consumer
	group    string
}
type SendFriendApplicationMessage struct {
	UserID   int64 `json:"user_id"`
	ToUserId int64 `json:"touser_id"`
}

func init() {
	SendFriendApplicationConsumer, err := NewSendFriendApplicationConsumer([]string{kafkaAddr})
	if err != nil {
		log.Println("无法接收消息到 Kafka:", err)
		return
	}
	go func() {
		SendFriendApplicationConsumer.Listen()
	}()
}
func NewSendFriendApplicationConsumer(brokerList []string) (*SendFriendApplicationConsumer, error) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}
	return &SendFriendApplicationConsumer{
		consumer: consumer,
		group:    "sendfriendapplication-group",
	}, nil
}
func (c *SendFriendApplicationConsumer) Listen() {
	log.Println("listenSendFriendApplication")
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{kafkaAddr}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka client: %v", err)
	}
	defer client.Close()
	// 刷新元数据
	err = client.RefreshMetadata("send-friendapplication")
	if err != nil {
		log.Fatalf("Error refreshing metadata: %v", err)
	}

	partitionConsumer, err := c.consumer.ConsumePartition("send-friendapplication", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()
	ticker := time.NewTicker(time.Second / 20)
	defer ticker.Stop()
	for msg := range partitionConsumer.Messages() {
		<-ticker.C
		var event SendFriendApplicationMessage
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("反序列化好友申请事件失败: %v", err)
			continue
		}
		err = SendFriendApplication(event.UserID, event.ToUserId)
		if err != nil {
			log.Printf("好友申请失败: %v", err)
		}
	}
}
func SendFriendApplication(userID int64, toUserId int64) error {
	ctx := context.Background()
	usr, err := db.GetUserByID(ctx, userID)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return err
	}
	tousr, err := db.GetUserByID(ctx, toUserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return err
	}
	err = db.SendFriendApplication(ctx, usr, tousr)
	if err != nil {
		logger.Error("发送好友请求失败,服务器出错：", zap.Error(err))
		return err
	}
	log.Println("发送好友请求成功！")
	return nil
}
