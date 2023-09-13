package main

// 该示例应用程序公开了my_app_processed_ops_total 计数器，该计数器对到目前为止已处理的操作数量进行计数。每2秒，计数器增加1。
import (
	"monitor/helper"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	// 每2秒，计数器增加1。
	go func() {
		for {
			opsProcessed.Inc()
			counter := helper.GetRandom(10)
			opsProcessed.Add(float64(counter))
			time.Sleep(2 * time.Second)
		}
	}()
}

// 公开了my_app_processed_ops_total计数器
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_app_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
