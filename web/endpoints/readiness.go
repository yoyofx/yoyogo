package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
)

func UseReadiness(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded health-readiness endpoint.")

	router.GET("/actuator/health/readiness", func(ctx *context.HttpContext) {
		var appLife *abstractions.ApplicationLife
		_ = ctx.RequiredServices.GetService(&appLife)
		statusCode := 500
		if appLife.State == "up" {
			statusCode = 200
		}

		ctx.JSON(statusCode, context.H{"status": appLife.State})
	})
}
