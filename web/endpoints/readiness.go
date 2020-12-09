package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
)

func UseReadiness(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health/readiness", func(ctx *context.HttpContext) {
		status := "UP"

		var databases []abstractions.IDataSource
		_ = ctx.RequiredServices.GetService(&databases)
		dumpArrays := make([]map[string]interface{}, 0)
		for _, db := range databases {
			dump := make(map[string]interface{})
			dump["name"] = db.GetName()
			dump["status"] = false
			if db.Ping() {
				dump["status"] = true
			}
			dumpArrays = append(dumpArrays, dump)

		}

		//databases1 := &datasources.MySqlDataSource{}
		//_ = ctx.RequiredServices.GetServiceByName(&databases1,"db1")

		ctx.JSON(200, context.H{
			"status":    status,
			"databases": dumpArrays,
		})
	})
}
