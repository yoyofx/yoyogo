package endpoints

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/router"
)

func UsePrometheus(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded prometheus endpoint.")

	router.GET("/actuator/metrics", web.WarpHttpHandlerFunc(promhttp.Handler()))
}
