package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

type databaseHealth struct {
}

func UseReadiness(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health/readiness", func(ctx *Context.HttpContext) {
		status := "UP"

		var databases []Abstractions.IDataSource
		_ = ctx.RequiredServices.GetService(&databases)
		dbmap := make(map[string]interface{})
		for _, db := range databases {
			dbmap["name"] = db.GetName()
			dbmap["status"] = false
			if db.Ping() {
				dbmap["status"] = true
			}

		}

		//databases1 := &datasources.MySqlDataSource{}
		//_ = ctx.RequiredServices.GetServiceByName(&databases1,"db1")

		ctx.JSON(200, Context.H{
			"status":    status,
			"databases": dbmap,
		})
	})
}
