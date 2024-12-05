package product

import (
	"Eshop/pkg/kafka/consumer/product"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

// 初始化 Kafka 生产者
func NewKafkaProducer(brokerList []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %v", err)
	}

	return &KafkaProducer{producer: producer}, nil
}
func (k *KafkaProducer) SendDeleteProductImageEvent(productname string) error {
	message := &product.DeleteProductImageMessage{Productname: productname}
	msgBytes, err := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: "avatar-delete",
		Value: sarama.StringEncoder(msgBytes),
	}
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		log.Printf("发送删除商品图片消息到kafka失败: %v", err)
		return err
	}
	log.Println("发送删除商品图片消息到kafka成功")
	return nil
}
