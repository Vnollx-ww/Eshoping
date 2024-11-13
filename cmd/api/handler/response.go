package handler

import (
	"Eshop/kitex_gen/base"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func BadBaseResponse(c *app.RequestContext, s string) {
	c.JSON(http.StatusBadRequest, base.Base{
		StatusCode: http.StatusBadRequest,
		StatusMsg:  s,
	})
}
