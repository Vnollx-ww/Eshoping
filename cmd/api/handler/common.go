package handler

import (
	"Eshop/kitex_gen/base"
	"Eshop/pkg/viper"
	"Eshop/pkg/zaplog"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger = zaplog.GetLogger()
var config = viper.Init("order")
var kafkaAddr = fmt.Sprintf("%s:%d", config.Viper.GetString("kafka.host"), config.Viper.GetInt("kafka.port"))

func BadBaseResponse(c *app.RequestContext, s string) {
	c.JSON(http.StatusBadRequest, base.Base{
		StatusCode: http.StatusBadRequest,
		StatusMsg:  s,
	})
}
