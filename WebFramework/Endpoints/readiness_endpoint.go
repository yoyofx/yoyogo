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
		status := "UP"

		var databases []Abstractions.IDataSource
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

		ctx.JSON(200, Context.H{
			"status":    status,
			"databases": dumpArrays,
		})
	})
}
