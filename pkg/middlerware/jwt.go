package middlerware

import (
	"Eshop/kitex_gen/base"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

var Mysecret = []byte("vnollxvnollx")

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewToken(username string) (string, error) {
	c := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("JWT_ISSUER"),
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Mysecret)
}
func ParseToken(Token string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(Token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Mysecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效token")
}
func JWTAuthMiddleware() app.HandlerFunc {
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
		c.Set("username", mc.Username)
		// 继续处理请求
		c.Next(ctx)
	}
}
