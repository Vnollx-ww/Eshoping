package consumer

import (
	"Eshop/cmd/api/rpc"
	"Eshop/kitex_gen/product"
	"Eshop/kitex_gen/product/productservice"
	"Eshop/kitex_gen/user"
	"Eshop/kitex_gen/user/userservice"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type UserBalanceAndCostConsumer struct {
	consumer sarama.Consumer
	group    string
}
type ProductStockConsumer struct {
	consumer sarama.Consumer
	group    string
}
type CreateOrderMessage struct {
	Token     string `json:"token"`
	Amount    int64  `json:"amount"`
	Number    int64  `json:"number"`
	Productid int64  `json:"productid"`
}

var userclient userservice.Client
var proclient productservice.Client

func init() {
	UserBalanceAndCostConsumer, err := NewUserBalanceAndCostConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Println("无接收消息到 Kafka:", err)
		return
	}
	ProductStockConsumer, err := NewProductStockConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Println("无接收消息到 Kafka:", err)
		return
	}

	userclient = rpc.GetUserClient() // 初始化用户客户端
	proclient = rpc.GerProductClient()
	go func() {
		UserBalanceAndCostConsumer.ListenBalanceAndCost()
	}()
	go func() {
		ProductStockConsumer.ListenStock()
	}()
}
func NewUserBalanceAndCostConsumer(brokerList []string) (*UserBalanceAndCostConsumer, error) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}
	return &UserBalanceAndCostConsumer{
		consumer: consumer,
		group:    "user-balanceandcost-group",
	}, nil
}

func NewProductStockConsumer(brokerList []string) (*ProductStockConsumer, error) {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return nil, err
	}
	return &ProductStockConsumer{
		consumer: consumer,
		group:    "product-stock-group",
	}, nil
}

func (c *UserBalanceAndCostConsumer) ListenBalanceAndCost() {
	log.Println("listenbalanceandcost")
	partitionConsumer, err := c.consumer.ConsumePartition("order-create", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()
	ticker := time.NewTicker(time.Second / 20)
	defer ticker.Stop()
	for msg := range partitionConsumer.Messages() {
		<-ticker.C
		var event CreateOrderMessage
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("反序列化更新用户余额和花费事件失败: %v", err)
			continue
		}
		err = UpdateUserBalanceAndCost(event.Token, event.Amount)
		if err != nil {
			log.Printf("更新用户余额和花费失败: %v", err)
		}
	}
}
func (c *ProductStockConsumer) ListenStock() {
	log.Println("listenstock")
	partitionConsumer, err := c.consumer.ConsumePartition("order-create", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()
	ticker := time.NewTicker(time.Second / 20) // 每秒钟最多处理 20 条消息
	defer ticker.Stop()
	for msg := range partitionConsumer.Messages() {
		<-ticker.C
		var event CreateOrderMessage
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("反序列化更新商品库存事件失败: %v", err)
			continue
		}
		err = UpdateProductStock(event.Productid, event.Number)
		if err != nil {
			log.Printf("更新商品库存失败: %v", err)
		}
	}
}
func UpdateUserBalanceAndCost(token string, amount int64) error {
	ctx := context.Background()
	req := &user.UpdateBalanceAndCostRequest{
		Token:  token,
		Number: amount,
	}
	log.Println("尝试调用余额和花费")
	_, err := userclient.UpdateBalanceAndCost(ctx, req)
	return err
}
func UpdateProductStock(ID int64, stock int64) error {
	ctx := context.Background()
	req := &product.UpdateStockRequest{
		ProductId: ID,
		Addstock:  -stock,
	}
	log.Println("尝试调用库存")
	_, err := proclient.UpdateStock(ctx, req)
	return err
}
