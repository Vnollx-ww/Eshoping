package main

import (
	"Eshop/kitex_gen/product/productservice"
	"Eshop/pkg/viper"
	"Eshop/pkg/zaplog"
	"fmt"
	_ "github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
	"log"
	"net"
)

var (
	config                  = viper.Init("product")
	serviceName             = config.Viper.GetString("server.name")
	serviceAddr             = fmt.Sprintf("%s:%d", config.Viper.GetString("server.host"), config.Viper.GetInt("server.port"))
	etcdAddr                = fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	logger      *zap.Logger = zaplog.GetLogger()
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 初始化etcd
	s := productservice.NewServer(new(ProductServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)

	if err := s.Run(); err != nil {
		logger.Fatal("Service stopped with error", zap.String("serviceName", serviceName), zap.String("error", err.Error()))
	}
}
