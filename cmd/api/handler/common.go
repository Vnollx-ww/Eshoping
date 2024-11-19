package handler

import (
	"Eshop/kitex_gen/base"
	"Eshop/pkg/zaplog"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger = zaplog.GetLogger()

func BadBaseResponse(c *app.RequestContext, s string) {
	c.JSON(http.StatusBadRequest, base.Base{
		StatusCode: http.StatusBadRequest,
		StatusMsg:  s,
	})
}
