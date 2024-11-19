package kafka

import (
	_ "Eshop/pkg/kafka/consumer"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

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
	message := fmt.Sprintf(`{
		"token": "%s",
		"amount": %d,
        "number":%d,
        "productid":%d
	}`, token, amount, number, productid)

	msg := &sarama.ProducerMessage{
		Topic: "order-create",
		Value: sarama.StringEncoder(message),
	}
	// 发送消息到 Kafka
	_, _, err := k.producer.SendMessage(msg)
	if err != nil {
		log.Printf("发送消息到kafka失败: %v", err)
		return err
	}
	log.Println("发送订单消息到kafka成功")
	return nil
}