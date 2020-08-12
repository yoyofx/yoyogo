package Endpoints

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	YoyoGo "github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func UsePrometheus(router Router.IRouterBuilder) {
	router.GET("/actuator/metrics", YoyoGo.WarpHttpHandlerFunc(promhttp.Handler()))
}
