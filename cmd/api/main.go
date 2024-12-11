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
		//c.File("D:/GolandProgram/Eshoping/web/html/index.html")
		c.File("/app/web/html/index.html")
	})
	hz.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/login.html")
		c.File("/app/web/html/user/user/login.html")
	})
	hz.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/register.html")
		c.File("/app/web/html/user/user/register.html")
	})
	hz.GET("/userinfo", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/userinfo.html")
		c.File("/app/web/html/user/user/userinfo.html")
	})
	hz.GET("/homepage", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/homepage.html")
		c.File("/app/web/html/user/user/homepage.html")
	})
	hz.GET("/recharge", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/recharge.html")
		c.File("/app/web/html/user/user/recharge.html")
	})
	hz.GET("/updatename", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/updatename.html")
		c.File("/app/web/html/user/user/updatename.html")
	})
	hz.GET("/updatepassword", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/updatepassword.html")
		c.File("/app/web/html/user/user/updatepassword.html")
	})
	hz.GET("/updateaddress", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/updateaddress.html")
		c.File("/app/web/html/user/user/updateaddress.html")
	})
	hz.GET("/updateavatar", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/user/updateavatar.html")
		c.File("/app/web/html/user/user/updateavatar.html")
	})
	hz.GET("/getorderlist", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/order/getorderlist.html")
		c.File("/app/web/html/user/order/getorderlist.html")
	})
	hz.GET("/addproduct", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/product/addproduct.html")
		c.File("/app/web/html/user/product/addproduct.html")
	})
	hz.GET("/deleteproduct", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/product/deleteproduct.html")
		c.File("/app/web/html/user/product/deleteproduct.html")
	})
	hz.GET("/updateprice", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/product/updateprice.html")
		c.File("/app/web/html/user/product/updateprice.html")
	})
	hz.GET("/updatestock", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/product/updatestock.html")
		c.File("/app/web/html/user/product/updatestock.html")
	})
	hz.GET("/getorderlistbystate", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/order/getorderlistbystate.html")
		c.File("/app/web/html/user/order/getorderlistbystate.html")
	})
	hz.GET("/friend", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/friend/friend.html")
		c.File("/app/web/html/user/friend/friend.html")
	})
	hz.GET("/targetuser", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/friend/targetuser.html")
		c.File("/app/web/html/user/friend/targetuser.html")
	})
	hz.GET("/deletefriend", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/friend/deletefriend.html")
		c.File("/app/web/html/user/friend/deletefriend.html")
	})
	hz.GET("/friendrequest", func(ctx context.Context, c *app.RequestContext) {
		//c.File("D:/GolandProgram/Eshoping/web/html/user/friend/friendrequest.html")
		c.File("/app/web/html/user/friend/friendrequest.html")
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
		user.POST("/sendfriendapplication", handler.SendFriendApplication)
		user.POST("/getfriendapplicationlist", handler.GetFriendApplicationList)
		user.POST("/rejectfriendapplication", handler.RejectFriendApplication)
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
