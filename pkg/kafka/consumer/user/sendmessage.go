package user

import (
	"Eshop/dal/db"
	"Eshop/pkg/middlerware"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"log"
	"time"
)

/*
PS D:\kafka_2.12-3.9.0> .\bin\windows\kafka-topics.bat --delete --topic send-message --bootstrap-server localhost:9092
PS D:\kafka_2.12-3.9.0> .\bin\windows\kafka-topics.bat --create --topic send-message --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
*/
type SendMessageConsumer struct {
	consumer sarama.Consumer
	group    string
}
type SendMessageMessage struct {
	Token    string `json:"token"`
	Content  string `json:"content"`
	ToUserId int64  `json:"touser_id"`
}

func init() {
	SendMessageConsumer, err := NewSendMessageConsumer([]string{kafkaAddr})
	if err != nil {
		log.Println("无法接收消息到 Kafka:", err)
		return
	}
	go func() {
		SendMessageConsumer.Listen()
	}()
}
func NewSendMessageConsumer(brokerList []string) (*SendMessageConsumer, error) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}
	return &SendMessageConsumer{
		consumer: consumer,
		group:    "sendmessage-group",
	}, nil
}
func (c *SendMessageConsumer) Listen() {
	log.Println("listenSendMessage")
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{kafkaAddr}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka client: %v", err)
	}
	defer client.Close()
	// 刷新元数据
	err = client.RefreshMetadata("send-message")
	if err != nil {
		log.Fatalf("Error refreshing metadata: %v", err)
	}
	partitionConsumer, err := c.consumer.ConsumePartition("send-message", 0, sarama.OffsetNewest)
	if err != nil {
		log.Println("message")
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()
	ticker := time.NewTicker(time.Second / 100)
	defer ticker.Stop()
	for msg := range partitionConsumer.Messages() {
		<-ticker.C
		var event SendMessageMessage
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("反序列化发送消息事件失败: %v", err)
			continue
		}
		err = SendMessage(event.Token, event.Content, event.ToUserId)
		if err != nil {
			log.Printf("发送消息失败: %v", err)
		}
	}
}

func SendMessage(Token string, content string, ToUserId int64) error {
	ctx := context.Background()
	log.Println("尝试调用发送消息")
	mc, err := middlerware.ParseToken(Token)
	if err != nil {
		logger.Error("token解析失败：", zap.Error(err))
		return err
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return err
	}
	tousr, err := db.GetUserByID(ctx, ToUserId)
	if err != nil {
		logger.Error("获取用户信息失败，服务器内部错误：", zap.Error(err))
		return err
	}
	err = db.SendMessage(ctx, usr, tousr, content)
	if err != nil {
		logger.Error("发送消息失败,服务器出错：", zap.Error(err))
		return err
	}
	return nil
}
