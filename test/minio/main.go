package main

import (
	"errors"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/util"
	"log"
	"math/rand"
)

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %d, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func main() {
	//基于连接数的降级模式
	conf := config.NewDefaultConfig()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})
	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Resource:         "abc",
			Strategy:         circuitbreaker.ErrorCount,
			RetryTimeoutMs:   3000, //3s只有尝试回复
			MinRequestAmount: 10,   //静默数
			StatIntervalMs:   5000,
			Threshold:        50,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				fmt.Println("协程熔断了")
			} else {
				if rand.Uint64()%20 > 9 {
					sentinel.TraceError(e, errors.New("biz error"))
				}
				e.Exit()
			}
		}
	}()
	select {}
}
