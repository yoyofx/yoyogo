package main

import (
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Middleware"
)

func main() {

	app := YoyoGo.UseMvc()
	//http.Handle("/Static", http.StripPrefix("/Static", http.FileServer(http.Dir("/Users/yoyofx/Documents/Src/yoyogo/Static/"))))

	app.Use(Middleware.NewStatic("Static"))
	app.Map("/info", func(ctx *Middleware.HttpContext) {
		ctx.JSON(YoyoGo.M{"info": "hello world"})
	})

	app.Run(":8080")

}
