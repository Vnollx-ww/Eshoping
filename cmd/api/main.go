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
		c.File("D:/GolandProgram/Eshoping/web/html/user/login.html")
	})
	hz.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/register.html")
	})
	hz.GET("/homepage", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/homepage.html")
	})
	hz.GET("/shop", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/shop.html")
	})
	hz.GET("/recharge", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/recharge.html")
	})
	hz.GET("/updatename", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updatename.html")
	})
	hz.GET("/updatepassword", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updatepassword.html")
	})
	hz.GET("/getorderlist", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/getorderlist.html")
	})
	hz.GET("/admin", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/admin.html")
	})
	hz.GET("/addproduct", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/addproduct.html")
	})
	hz.GET("/deleteproduct", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/deleteproduct.html")
	})
	hz.GET("/updateprice", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/updateprice.html")
	})
	hz.GET("/updatestock", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/updatestock.html")
	})
}
func registerGroup(hz *server.Hertz) {
	user := hz.Group("/user")
	{
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
		user.POST("/getuserinfo", handler.GetUserInfo)
		user.POST("/updatepassword", handler.UpdatePassword)
		user.POST("/updatename", handler.UpdateName)
		user.POST("/updatecost", handler.UpdateCost)
		user.POST("/updatebalance", handler.UpdateBalance)
	}
	product := hz.Group("/product")
	{
		product.POST("/addproduct", handler.AddProduct)
		product.POST("/deleteproduct", handler.DelProduct)
		product.POST("/getproductinfo", handler.GetProductInfo)
		product.GET("/getproductlistinfo", handler.GetProductListInfo)
		product.POST("/updatestock", handler.Updatestock)
		product.POST("/updateprice", handler.UpdatePrice)
	}
	order := hz.Group("/order")
	{
		order.POST("/addorder", handler.CreateOrder)
		order.POST("/deleteorder", handler.DeleteOrder)
		order.POST("/orderlistbyuserid", handler.GetOrderListByUserID)
		order.POST("/orderlistbyproductname", handler.GetOrderListByProductName)
	}
}

func main() {
	hz := server.New(server.WithHostPorts("localhost:8889"))
	hz.Static("/images", "./web/images")
	LoadHtml(hz)
	registerGroup(hz)
	hz.Spin()
}
