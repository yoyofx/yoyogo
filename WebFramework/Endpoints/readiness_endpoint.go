package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func UseReadiness(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health/readiness", func(ctx *Context.HttpContext) {
		var database Abstractions.IDataSource
		_ = ctx.RequiredServices.GetService(&database)

		ctx.JSON(200, Context.H{
			"status": "UP",
		})
	})
}
