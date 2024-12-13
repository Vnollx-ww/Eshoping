package middlerware

import (
	"Eshop/kitex_gen/base"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sony/gobreaker"
	"log"
	"net/http"
	"sync"
	"time"
)

type res struct {
	StautsCode int32
	StautsMsg  string
}

var circuitBreakers sync.Map // 使用 sync.Map 存储每个路由的熔断器

// 创建一个新的熔断器
func createCircuitBreaker(route string) *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:        route,
		MaxRequests: 10,               // 允许最多5个请求
		Interval:    10 * time.Second, // 每10秒进行健康检查
		Timeout:     30 * time.Second, // 超过30秒未响应则认为熔断
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			totalRequests := float64(counts.TotalSuccesses + counts.TotalFailures)
			if totalRequests <= 10 {
				return false // 如果没有请求，则不触发熔断
			}
			if counts.ConsecutiveFailures > 5 {
				return true
			}
			failureRate := float64(counts.TotalFailures) / totalRequests
			if failureRate > 0.5 {
				log.Println(failureRate)
			}
			return failureRate > 0.5 // 当失败率大于50%时触发熔断
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("熔断器%s:状态从%s改变到%s\n", name, from, to)
		},
		IsSuccessful: func(err error) bool {
			return err == nil // 请求成功时返回true
		},
	}

	return gobreaker.NewCircuitBreaker(settings)
}

// 获取指定路由的熔断器
func getCircuitBreaker(route string) *gobreaker.CircuitBreaker {
	if cb, ok := circuitBreakers.Load(route); ok {
		return cb.(*gobreaker.CircuitBreaker)
	}

	// 如果没有对应的熔断器，则创建一个新的
	cb := createCircuitBreaker(route)
	circuitBreakers.Store(route, cb)
	return cb
}

// 中间件，熔断器
func CrcuitBreakerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		route := c.Request.Path() // 获取当前请求的路由（路径）
		// 获取当前路由的熔断器
		cb := getCircuitBreaker(string(route))
		state := cb.State()
		if state == gobreaker.StateOpen {
			// 如果熔断器是开启状态，直接返回服务不可用错误
			log.Println("服务已被熔断!")
			c.JSON(http.StatusBadRequest, base.Base{
				StatusCode: -1,
				StatusMsg:  "服务暂不可用，请稍后重试",
			})
			c.Abort()
			return
		}

		// 执行请求，使用熔断器包装
		_, err := cb.Execute(func() (interface{}, error) {
			c.Next(ctx) // 继续处理请求
			if err, exists := c.Get("error"); err != nil && exists == true {
				// 如果在 handler 中设置了错误，捕获并返回
				return nil, err.(error) // 将错误返回给熔断器
			}
			//log.Println("请求成功！qaq")
			return nil, nil
		})
		if err != nil {
			log.Println("服务通过！")
			c.JSON(http.StatusBadRequest, base.Base{
				StatusCode: -1,
				StatusMsg:  "服务处理失败，请重试",
			})
		}
	}
}
