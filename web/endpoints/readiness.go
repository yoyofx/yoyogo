package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions/health"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
)

func UseReadiness(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health/readiness", func(ctx *context.HttpContext) {
		var indicatorList []health.Indicator
		_ = ctx.RequiredServices.GetService(&indicatorList)
		builder := health.NewHealthIndicator(indicatorList)
		root := builder.Build()
		statusCode := 200
		if root["status"] != "up" {
			statusCode = 500
		}

		ctx.JSON(statusCode, builder.Build())
	})
}
