package main

import (
	"Eshop/cmd/api/rpc"
	orderlist "Eshop/kitex_gen/orderlist/orderlistservice"
	"Eshop/pkg/viper"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

var (
	config        = viper.Init("order")
	userconfig    = viper.Init("user")
	productconfig = viper.Init("product")
	serviceName   = config.Viper.GetString("server.name")
	serviceAddr   = fmt.Sprintf("%s:%d", config.Viper.GetString("server.host"), config.Viper.GetInt("server.port"))
	etcdAddr      = fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	//signingKey  = config.Viper.GetString("JWT.signingKey")
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8885")
	if err != nil {
		log.Fatalln(err)
	}
	orderlistserviceimpl := new(OrderListServiceImpl)
	orderlistserviceimpl.usrcli = rpc.GetUserClient()
	orderlistserviceimpl.procli = rpc.GerProductClient()
	// 初始化etcd
	s := orderlist.NewServer(orderlistserviceimpl,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err := s.Run(); err != nil {
		logger.Fatalf("%v stopped with error: %v", serviceName, err.Error())
	}
}
