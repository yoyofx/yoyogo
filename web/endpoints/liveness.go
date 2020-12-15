package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
)

func UseLiveness(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health/liveness", func(ctx *context.HttpContext) {
		ctx.JSON(200, context.H{
			"status": "UP",
		})
	})
}
