package consumer

import (
	"Eshop/pkg/minio"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type ProductImageConsumer struct {
	consumer sarama.Consumer
	group    string
}
type DeleteProductImageMessage struct {
	Productname string `json:"productname"`
}

func init() {
	ProductImageConsumer, err := NewProductImageConsumer([]string{kafkaAddr})
	if err != nil {
		log.Println("无法接收消息到 Kafka:", err)
		return
	}
	go func() {
		ProductImageConsumer.Listen()
	}()
}
func NewProductImageConsumer(brokerList []string) (*ProductImageConsumer, error) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}
	return &ProductImageConsumer{
		consumer: consumer,
		group:    "ProductImage-group",
	}, nil
}
func (c *ProductImageConsumer) Listen() {
	log.Println("listenProductImage")
	partitionConsumer, err := c.consumer.ConsumePartition("ProductImage-delete", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()
	ticker := time.NewTicker(time.Second / 20)
	defer ticker.Stop()
	for msg := range partitionConsumer.Messages() {
		<-ticker.C
		var event DeleteProductImageMessage
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("反序列化删除商品图片失败: %v", err)
			continue
		}
		err = DeleteAvatar(event.Productname)
		if err != nil {
			log.Printf("删除商品图片失败: %v", err)
		}
	}
}
func DeleteProductImage(productname string) error {
	ctx := context.Background()
	return minio.ProductDeleteFileMinio(ctx, productname)
}
