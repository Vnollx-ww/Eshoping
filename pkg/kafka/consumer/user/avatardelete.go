package user

import (
	"Eshop/pkg/minio"
	"Eshop/pkg/viper"
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

var (
	config    = viper.Init("user")
	kafkaAddr = fmt.Sprintf("%s:%d", config.Viper.GetString("kafka.host"), config.Viper.GetInt("kafka.port"))
)

type AvatarConsumer struct {
	consumer sarama.Consumer
	group    string
}
type DeleteAvatarMessage struct {
	Username string `json:"username"`
}

func init() {
	AvatarConsumer, err := NewAvatarConsumer([]string{kafkaAddr})
	if err != nil {
		log.Println("无法接收消息到 Kafka:", err)
		return
	}
	go func() {
		AvatarConsumer.Listen()
	}()
}
func NewAvatarConsumer(brokerList []string) (*AvatarConsumer, error) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}
	return &AvatarConsumer{
		consumer: consumer,
		group:    "avatar-group",
	}, nil
}
func (c *AvatarConsumer) Listen() {
	log.Println("listenavatar")
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{kafkaAddr}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka client: %v", err)
	}
	defer client.Close()

	// 刷新元数据
	err = client.RefreshMetadata("avatar-delete")
	if err != nil {
		log.Fatalf("Error refreshing metadata: %v", err)
	}
	partitionConsumer, err := c.consumer.ConsumePartition("avatar-delete", 0, sarama.OffsetNewest)
	if err != nil {
		log.Println("avatar")
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()
	ticker := time.NewTicker(time.Second / 20)
	defer ticker.Stop()
	for msg := range partitionConsumer.Messages() {
		<-ticker.C
		var event DeleteAvatarMessage
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("反序列化删除头像事件失败: %v", err)
			continue
		}
		err = DeleteAvatar(event.Username)
		if err != nil {
			log.Printf("删除头像失败: %v", err)
		}
	}
}
func DeleteAvatar(username string) error {
	ctx := context.Background()
	return minio.UserDeleteFileMinio(ctx, username)
}
