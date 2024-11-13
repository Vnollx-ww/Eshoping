package main

import (
	"Eshop/cmd/api/handler"
	"Eshop/pkg/viper"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var (
	apiConfig     = viper.Init("api")
	apiServerName = apiConfig.Viper.GetString("server.name")
	apiServerAddr = fmt.Sprintf("%s:%d", apiConfig.Viper.GetString("server.host"), apiConfig.Viper.GetInt("server.port"))
	etcdAddress   = fmt.Sprintf("%s:%d", apiConfig.Viper.GetString("Etcd.host"), apiConfig.Viper.GetInt("Etcd.port"))
	serverTLSKey  = apiConfig.Viper.GetString("Hertz.tls.keyFile")
	serverTLSCert = apiConfig.Viper.GetString("Hertz.tls.certFile")
)

func LoadHtml(hz *server.Hertz) {
	hz.GET("/", func(ctx context.Context, c *app.RequestContext) {
		// 你可以通过 c.File() 返回文件
		c.File("D:/GolandProgram/Eshoping/web/html/index.html")
	})
	hz.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/login.html")
	})
	hz.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/register.html")
	})
	hz.GET("/homepage", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/homepage.html")
	})
	hz.GET("/shop", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/shop.html")
	})
	hz.GET("/recharge", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/recharge.html")
	})
	hz.GET("/updatename", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/updatename.html")
	})
	hz.GET("/updatepassword", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/updatepassword.html")
	})
	hz.GET("/getorderlist", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/getorderlist.html")
	})
}
func registerGroup(hz *server.Hertz) {
	hz.POST("/register", handler.Register)
	hz.POST("/login", handler.Login)
	hz.POST("/userinfo", handler.GetUserInfo)
	hz.POST("/updatename", handler.UpdateName)
	hz.POST("/updatepassword", handler.UpdatePassword)
	hz.POST("/updatecost", handler.UpdateCost)
	hz.POST("/updatebalance", handler.UpdateBalance)
	hz.POST("/addproduct", handler.AddProduct)
	hz.POST("/delproduct", handler.DelProduct)
	hz.POST("/productinfo", handler.GetProductInfo)
	hz.POST("/updatestock", handler.Updatestock)
	hz.POST("/updateprice", handler.UpdatePrice)
	hz.POST("/createorder", handler.CreateOrder)
	hz.POST("/deleteorder", handler.DeleteOrder)
	hz.POST("/orderlistbyuserid", handler.GetOrderListByUserID)
	hz.POST("/orderlistbyproductname", handler.GetOrderListByProductName)
}

func main() {
	hz := server.New(server.WithHostPorts("localhost:8889"))
	hz.Static("/images", "./web/images")
	LoadHtml(hz)
	registerGroup(hz)
	hz.Spin()
}
