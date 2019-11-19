package main

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Router"
	"github.com/maxzhang1985/yoyogo/Standard"
	"os"
)

func main() {

	webHost := CreateWebHostBuilder(os.Args).Build()
	webHost.Run()

}

//region Create the builder of Web host
func CreateWebHostBuilder(args []string) YoyoGo.HostBuilder {
	return YoyoGo.NewWebHostBuilder().
		UseServer(YoyoGo.DefaultHttpServer(":8080")).
		Configure(func(app *YoyoGo.ApplicationBuilder) {
			//app.UseMvc()
			app.UseStatic("Static")
		}).
		UseRouter(func(router Router.IRouterBuilder) {

			router.GET("/error", func(ctx *Context.HttpContext) {
				panic("http get error")
			})

			router.POST("/info/:id", PostInfo)
			router.Group("/v1/api", func(router *Router.RouterGroup) {
				router.GET("/info", GetInfo)
			})

			router.GET("/info", GetInfo)
		})
}

//endregion

//region Http Request Methods

type UserInfo struct {
	UserName string `param:"username"`
	Number   string `param:"q1"`
	Id       string `param:"id"`
}

//HttpGet request: /info  or /v1/api/info
//bind UserInfo for id,q1,username
func GetInfo(ctx *Context.HttpContext) {
	ctx.JSON(200, Std.M{"info": "ok"})
}

//HttpPost request: /info/:id ?q1=abc&username=123
func PostInfo(ctx *Context.HttpContext) {
	qs_q1 := ctx.Query("q1")
	pd_name := ctx.Param("username")

	userInfo := &UserInfo{}
	_ = ctx.Bind(userInfo)

	strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)

	ctx.JSON(200, Std.M{"info": "hello world", "result": strResult})
}

//endregion
