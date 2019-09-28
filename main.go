package main

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Middleware"
	"github.com/maxzhang1985/yoyogo/Standard"
)

type UserInfo struct {
	UserName string `param:"username"`
	Number   string `param:"q1"`
	Id       string `param:"id"`
}

func main() {

	app := YoyoGo.UseMvc()

	app.Recovery.PanicHandlerFunc = func(information *Middleware.PanicInformation) {
		fmt.Println("----------------------------------------------ERROR----------------------------------------------------")
		fmt.Println("********************************  Global Recovery Display Errors  *************************************")
		fmt.Println("-----------------------------------------------END-----------------------------------------------------")

	}

	app.UseStatic("Static")

	app.GET("/error", func(ctx *Middleware.HttpContext) {
		panic("http get error")
	})

	app.POST("/info/:id", PostInfo)

	app.Group("/v1/api", func(router *Middleware.RouterGroup) {
		router.POST("/info/:id", PostInfo)
	})

	app.Run(":8080")

}

func PostInfo(ctx *Middleware.HttpContext) {

	qs_q1 := ctx.Query("q1")
	pd_name := ctx.Param("username")

	userInfo := &UserInfo{}
	_ = ctx.Bind(userInfo)

	strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)

	ctx.JSON(200, Std.M{"info": "hello world", "result": strResult})
}
