package main

import (
	"github.com/maxzhang1985/yoyogo/Framework"
)

func main() {

	app := YoyoGo.Classic()

	app.Map("/info", func(ctx *YoyoGo.HttpContext) {
		ctx.JSON(YoyoGo.M{"info": "hello world"})
	})

	app.Run(":8080")

}
