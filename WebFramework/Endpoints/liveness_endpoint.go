package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func UseLiveness(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health/liveness", func(ctx *Context.HttpContext) {
		ctx.JSON(200, Context.H{
			"status": "UP",
		})
	})
}
