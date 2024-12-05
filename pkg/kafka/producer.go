package kafka

import (
	"Eshop/pkg/kafka/consumer"
	_ "Eshop/pkg/kafka/consumer"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

// cd D:
// cd kafka_2.13-3.9.0
// .\bin\windows\kafka-server-start.bat .\config\server.properties
// Kafka 生产者结构体
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

func (k *KafkaProducer) SendCreateOrderEvent(token string, amount int64, number int64, productid int64) error {
	message := &consumer.CreateOrderMessage{Token: token, Amount: amount, Number: number, Productid: productid}
	msgBytes, err := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: "order-create",
		Value: sarama.StringEncoder(msgBytes),
	}
	// 发送消息到 Kafka
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		log.Printf("发送消息到kafka失败: %v", err)
		return err
	}
	log.Println("发送订单消息到kafka成功")
	return nil
}
func (k *KafkaProducer) SendDeleteAvatarEvent(username string) error {
	message := &consumer.DeleteAvatarMessage{Username: username}
	msgBytes, err := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: "avatar-delete",
		Value: sarama.StringEncoder(msgBytes),
	}
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		log.Printf("发送删除头像消息到kafka失败: %v", err)
		return err
	}
	log.Println("发送删除头像消息到kafka成功")
	return nil
}
func (k *KafkaProducer) SendDeleteProductImageEvent(productname string) error {
	message := &consumer.DeleteProductImageMessage{Productname: productname}
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
