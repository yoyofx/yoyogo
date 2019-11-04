package main

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Middleware"
	"github.com/maxzhang1985/yoyogo/Standard"
	"os"
)

type UserInfo struct {
	UserName string `param:"username"`
	Number   string `param:"q1"`
	Id       string `param:"id"`
}

func main() {
	//region PanicHandlerFunc

	//app.Recovery.PanicHandlerFunc = func(information *Middleware.PanicInformation) {
	//	fmt.Println("----------------------------------------------ERROR----------------------------------------------------")
	//	fmt.Println("********************************  Global Recovery Display Errors  *************************************")
	//	fmt.Println("-----------------------------------------------END-----------------------------------------------------")
	//
	//}
	//endregion
	webHost := CreateWebHostBuilder(os.Args).Build()
	webHost.Run()

}

func CreateWebHostBuilder(args []string) YoyoGo.HostBuilder {
	return YoyoGo.NewWebHostBuilder().
		UseServer(YoyoGo.DefaultHttpServer(":8080")).
		Configure(func(app *YoyoGo.ApplicationBuilder) {
			//app.UseMvc()
			app.UseStatic("Static")
		}).
		UseRouter(func(router Middleware.IRouterBuilder) {

			router.GET("/error", func(ctx *Middleware.HttpContext) {
				panic("http get error")
			})

			router.POST("/info/:id", PostInfo)
			router.Group("/v1/api", func(router *Middleware.RouterGroup) {
				router.GET("/info", GetInfo)
			})

			router.GET("/info", GetInfo)
		})
}

func GetInfo(ctx *Middleware.HttpContext) {
	ctx.JSON(200, Std.M{"info": "ok"})
}

func PostInfo(ctx *Middleware.HttpContext) {

	qs_q1 := ctx.Query("q1")
	pd_name := ctx.Param("username")

	userInfo := &UserInfo{}
	_ = ctx.Bind(userInfo)

	strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)

	ctx.JSON(200, Std.M{"info": "hello world", "result": strResult})
}
