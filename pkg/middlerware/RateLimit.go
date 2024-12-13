package middlerware

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

func init() {
	conf := config.NewDefaultConfig()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
}
func RateLimitMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		resource := string(c.Request.Path())
		_, err := flow.LoadRules([]*flow.Rule{
			{
				// 资源名称
				Resource:               resource,
				TokenCalculateStrategy: flow.Direct,
				ControlBehavior:        flow.Reject,
				Threshold:              100,
				StatIntervalInMs:       1000,
				RelationStrategy:       flow.CurrentResource,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		e, b := sentinel.Entry(resource, sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			c.JSON(http.StatusBadRequest, res{
				StautsCode: -1,
				StautsMsg:  "请求繁忙，请稍后重试！",
			})
			c.Abort()
			return
		}
		c.Next(ctx)
		e.Exit()
	}
}
