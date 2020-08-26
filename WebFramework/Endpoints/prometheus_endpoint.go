package Endpoints

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func UsePrometheus(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded prometheus endpoint.")

	router.GET("/actuator/metrics", YoyoGo.WarpHttpHandlerFunc(promhttp.Handler()))
}
