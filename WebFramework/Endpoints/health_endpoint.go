package Endpoints

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func UseHealth(router Router.IRouterBuilder) {
	router.GET("/actuator/health", func(ctx *Context.HttpContext) {
		ctx.JSON(200, Context.M{
			"status": "UP",
		})
	})
}
