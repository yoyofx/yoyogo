package main

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Middleware"
	Std "github.com/maxzhang1985/yoyogo/Standard"
)

type UserInfo struct {
	UserName string `json:"username"`
}

func main() {

	app := YoyoGo.UseMvc()
	app.UseStatic("Static")

	app.Map("/info", func(ctx *Middleware.HttpContext) {

		qs_q1 := ctx.Query("q1")
		pd_name := ctx.Param("username")

		userInfo := &UserInfo{}
		_ = ctx.Bind(userInfo)

		strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo.UserName)

		ctx.JSON(Std.M{"info": "hello world", "result": strResult})
	})

	app.Run(":8080")

}
