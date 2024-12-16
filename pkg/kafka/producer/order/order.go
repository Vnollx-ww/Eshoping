package order

import (
	"Eshop/pkg/kafka/consumer/order"
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"go.opentelemetry.io/otel"
	"log"
)

// cd D:
// cd kafka_2.12-3.9.0
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
	tracer := otel.Tracer("order-service")
	// 创建 Span
	_, span := tracer.Start(context.Background(), "SendCreateOrderEvent")
	defer span.End()
	message := &order.CreateOrderMessage{Token: token, Amount: amount, Number: number, Productid: productid}
	msgBytes, err := json.Marshal(message)
	if err != nil {
		span.RecordError(err)
		log.Printf("消息序列化失败: %v", err)
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: "order-create",
		Value: sarama.StringEncoder(msgBytes),
	}
	// 发送消息到 Kafka
	log.Println(message.Amount)
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		span.RecordError(err)
		log.Printf("发送消息到kafka失败: %v", err)
		return err
	}
	log.Println("发送订单消息到kafka成功")
	return nil
}
