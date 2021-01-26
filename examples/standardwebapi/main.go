package main

import (
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/endpoints"
	"github.com/yoyofx/yoyogo/web/router"
)

func main() {
	web.CreateHttpBuilder(func(rb router.IRouterBuilder) {
		// 运维特性
		endpoints.UseHealth(rb)
		endpoints.UseViz(rb)
		endpoints.UsePrometheus(rb)
		endpoints.UsePprof(rb)
		endpoints.UseReadiness(rb)
		endpoints.UseLiveness(rb)

		// 标准接口
		rb.GET("/info", func(ctx *context.HttpContext) {
			ctx.JSON(200, context.H{"info": "ok"})
		})

		rb.Group("/api", func(rg *router.RouterGroup) {
			rg.GET("/info", func(ctx *context.HttpContext) {
				ctx.JSON(200, context.H{"api/info": "ok"})
			})
		})

	}).Build().Run()
}
