package main

import (
	"net/http"
	"time"

	"monitor/helper"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 用来模拟 CPU 涨幅
	var count = 0
	go func() {
		for {
			count += helper.GetRandom(10)
			count -= helper.GetRandom(10)
			// 随机睡眠时间
			time.Sleep(time.Nanosecond * time.Duration(helper.GetRandom(1000)))
		}
	}()

	// 模拟 goroutine 涨幅与下降代码
	go func() {
		for {
			go func() {
				time.Sleep(time.Second * time.Duration(helper.GetRandom(5)))
			}()
			time.Sleep(time.Second * time.Duration(helper.GetRandom(5)))
		}
	}()
	// 注册 prometheus 监控指标
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe("0.0.0.0:2112", nil)
}
