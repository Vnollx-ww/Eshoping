package main

import (
	orderlist "Eshop/kitex_gen/orderlist/orderlistservice"
	"Eshop/pkg/jaeger"
	"Eshop/pkg/viper"
	"Eshop/pkg/zaplog"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
	"log"
	"net"
)

var (
	config                  = viper.Init("order")
	serviceName             = config.Viper.GetString("server.name")
	serviceAddr             = fmt.Sprintf("%s:%d", config.Viper.GetString("server.host"), config.Viper.GetInt("server.port"))
	etcdAddr                = fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	kafkaAddr               = fmt.Sprintf("%s:%d", config.Viper.GetString("kafka.host"), config.Viper.GetInt("kafka.port"))
	logger      *zap.Logger = zaplog.GetLogger()
)

func main() {
	/*ok := []attribute.KeyValue{
		attribute.String("IsOK?", "YES"),
	}
	bad:=[]attribute.KeyValue{
		attribute.String("IsOK?", "NO"),
	}*/
	ctx := context.Background()
	shutdown, err := jaeger.SetupTracer(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = shutdown(ctx)
	}()
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		log.Fatalln(err)
	}
	orderlistserviceimpl := new(OrderListServiceImpl)
	// 初始化etcd
	s := orderlist.NewServer(orderlistserviceimpl,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err := s.Run(); err != nil {
		logger.Fatal("Service stopped with error", zap.String("serviceName", serviceName), zap.String("error", err.Error()))
	}
}
