package middlerware

/*func RateLimitMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取请求头中的Authorization字段
		if string(c.Request.URI().Path()) == "/login" {
			c.Next(ctx)
			return
		}
		// 定义一个结构体，用来解析请求体中的 token
		var requestBody struct {
			Token string `json:"token"` // 假设token在请求体中是 "token" 字段
		}
		// 解析请求体
		if err := c.Bind(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, base.Base{
				StatusCode: http.StatusBadRequest,
				StatusMsg:  "请求体格式有误",
			})
			c.Abort()
			return
		}
		// 获取token
		token := requestBody.Token
		// 如果token为空，返回401 Unauthorized
		if token == "" {
			c.JSON(http.StatusUnauthorized, base.Base{
				StatusCode: http.StatusUnauthorized,
				StatusMsg:  "请求体中token为空",
			})
			c.Abort()
			return
		}
		// 解析token
		mc, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, base.Base{
				StatusCode: http.StatusUnauthorized,
				StatusMsg:  "无效的token",
			})
			c.Abort()
			return
		}
		// 将用户名保存到上下文中
		c.Set("userid", mc.UserId)
		// 继续处理请求
		c.Next(ctx)
	}
}*/
