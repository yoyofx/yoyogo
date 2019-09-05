package main

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Middleware"
)

func main() {

	app := YoyoGo.UseMvc()
	app.Use(Middleware.NewStatic("Static"))

	app.Map("/info", func(ctx *Middleware.HttpContext) {

		qs_q1 := ctx.Query("q1")
		pd_name := ctx.Param("username")

		strResult := fmt.Sprintf("Name:%s , Q1:%s", pd_name, qs_q1)

		ctx.JSON(YoyoGo.M{"info": "hello world", "result": strResult})
	})

	app.Run(":8080")

}
