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
	apiServerAddr = fmt.Sprintf("%s:%d", apiConfig.Viper.GetString("server.host"), apiConfig.Viper.GetInt("server.port"))
)

func LoadHtml(hz *server.Hertz) {
	hz.GET("/", func(ctx context.Context, c *app.RequestContext) {
		// 你可以通过 c.File() 返回文件
		c.File("D:/GolandProgram/Eshoping/web/html/index.html")
		//c.File("/app/web/html/index.html")
	})
	hz.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/login.html")
		//c.File("/app/web/html/user/login.html")
	})

	hz.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/register.html")
		//c.File("/app/web/html/user/register.html")
	})
	hz.GET("/userinfo", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/userinfo.html")
		//c.File("/app/web/html/user/homepage.html")
	})
	hz.GET("/homepage", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/homepage.html")
		//c.File("/app/web/html/user/homepage.html")
	})
	hz.GET("/shop", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/shop.html")
		//c.File("/app/web/html/user/shop.html")
	})
	hz.GET("/recharge", func(ctx context.Context, c *app.RequestContext) {
		//c.File("/app/web/html/user/recharge.html")
		c.File("D:/GolandProgram/Eshoping/web/html/user/recharge.html")
	})
	hz.GET("/updatename", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updatename.html")
		//c.File("/app/web/html/user/updatename.html")
	})
	hz.GET("/updatepassword", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updatepassword.html")
		//c.File("/app/web/html/user/updatepassword.html")
	})
	hz.GET("/updateaddress", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updateaddress.html")
		//c.File("/app/web/html/user/updateaddress.html")
	})
	hz.GET("/updateavatar", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updateavatar.html")
		//c.File("/app/web/html/user/updateavatar.html")
	})
	hz.GET("/getorderlist", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/getorderlist.html")
		//c.File("/app/web/html/user/getorderlist.html")
	})
	hz.GET("/adminlogin", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/adminlogin.html")
		//c.File("/app/web/html/admin/adminlogin.html")
	})
	hz.GET("/admin", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/admin/admin.html")
		//c.File("/app/web/html/admin/admin.html")
	})

	hz.GET("/addproduct", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/addproduct.html")
		//c.File("/app/web/html/user/addproduct.html")
	})
	hz.GET("/deleteproduct", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/deleteproduct.html")
		//c.File("/app/web/html/user/deleteproduct.html")
	})
	hz.GET("/updateprice", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updateprice.html")
		//c.File("/app/web/html/user/updateprice.html")
	})
	hz.GET("/updatestock", func(ctx context.Context, c *app.RequestContext) {
		c.File("D:/GolandProgram/Eshoping/web/html/user/updatestock.html")
		//c.File("/app/web/html/user/updatestock.html")
	})
	hz.GET("/getorderlistbystate", func(ctx context.Context, c *app.RequestContext) {
		//c.File("/app/web/html/admin/getorderlistbystate.html")
		c.File("D:/GolandProgram/Eshoping/web/html/admin/getorderlistbystate.html")
	})
	hz.GET("/friend", func(ctx context.Context, c *app.RequestContext) {
		//c.File("/app/web/html/user/friend.html")
		c.File("D:/GolandProgram/Eshoping/web/html/user/friend.html")
	})
	hz.GET("/targetuser", func(ctx context.Context, c *app.RequestContext) {
		//c.File("/app/web/html/user/target.html")
		c.File("D:/GolandProgram/Eshoping/web/html/user/targetuser.html")
	})
	hz.GET("/deletefriend", func(ctx context.Context, c *app.RequestContext) {
		//c.File("/app/web/html/user/deletefriend.html")
		c.File("D:/GolandProgram/Eshoping/web/html/user/deletefriend.html")
	})
}
func registerGroup(hz *server.Hertz) {
	user := hz.Group("/user")
	{
		user.GET("/captcha", handler.GetCaptcha)
		user.POST("/register", handler.Register)
		user.POST("/adminlogin", handler.AdminLogin)
		user.POST("/login", handler.Login)
		user.POST("/getuserinfo", handler.GetUserInfo)
		user.POST("/getuserinfobyuserid", handler.GetUserInfoByUserID)
		user.POST("/updatepassword", handler.UpdatePassword)
		user.POST("/updatename", handler.UpdateName)
		user.POST("/updatecost", handler.UpdateCost)
		user.POST("/updatebalance", handler.UpdateBalance)
		user.POST("/updateaddress", handler.UpdateAddress)
		user.POST("/updateavatar", handler.UpdateAvatar)
		user.POST("/getfriendlist", handler.GetFriendList)
		user.POST("/addfriend", handler.AddFriend)
		user.POST("/deletefriend", handler.DelFriend)
		user.POST("/getmessagelist", handler.GetMessageList)
		user.POST("/sendmessage", handler.SendMessage)
		user.GET("/ws", handler.WebSocketConnections)
		user.POST("/getuserlistbycontent", handler.GetUserListByContent)
	}
	product := hz.Group("/product")
	{
		product.POST("/addproduct", handler.AddProduct)
		product.POST("/deleteproduct", handler.DelProduct)
		product.POST("/getproductinfo", handler.GetProductInfo)
		product.GET("/getproductlistinfo", handler.GetProductListInfo)
		product.POST("/updatestock", handler.Updatestock)
		product.POST("/updateprice", handler.UpdatePrice)
		product.POST("/getproductlistinfobyuser", handler.GetProductListInfoByUser)
	}
	order := hz.Group("/order")
	{
		order.POST("/addorder", handler.CreateOrder)
		order.POST("/deleteorder", handler.DeleteOrder)
		order.POST("/updatestate", handler.UpdateOrderState)
		order.POST("/orderlistbyuserid", handler.GetOrderListByUserID)
		order.POST("/orderlistbyproductname", handler.GetOrderListByProductName)
		order.POST("/orderlistbystate", handler.GetOrderListByState)
	}
}

func main() {
	//ws :=handler.NewWebSocketServer()
	//go ws.HandleMessages()
	hz := server.New(server.WithHostPorts(apiServerAddr))
	hz.NoHijackConnPool = true
	LoadHtml(hz)
	registerGroup(hz)
	hz.Spin()
}
