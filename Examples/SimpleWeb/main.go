package main

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/Examples/SimpleWeb/contollers"
	"github.com/maxzhang1985/yoyogo/Examples/SimpleWeb/models"
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Mvc"
	"github.com/maxzhang1985/yoyogo/Router"
)

func main() {

	//webHost := YoyoGo.CreateDefaultBuilder(registerEndpointRouterConfig).Build()
	webHost := CreateCustomBuilder().Build()
	webHost.Run()
}

//* Create the builder of Web host
func CreateCustomBuilder() *YoyoGo.HostBuilder {
	return YoyoGo.NewWebHostBuilder().
		//UseHttp().
		UseFastHttp().
		//UseServer(YoyoGo.DefaultHttps(":8080", "./Certificate/server.pem", "./Certificate/server.key")).
		Configure(func(app *YoyoGo.ApplicationBuilder) {
			app.SetEnvironment(Context.Dev)
			app.UseStatic("Static")
			app.UseEndpoints(registerEndpointRouterConfig)

			app.UseMvc()
			app.ConfigureMvcParts(func(builder *Mvc.ControllerBuilder) {
				builder.AddController(contollers.NewUserController)
			})
		}).
		ConfigureServices(func(serviceCollection *DependencyInjection.ServiceCollection) {
			serviceCollection.AddTransientByImplements(models.NewUserAction, new(models.IUserAction))
		}).
		OnApplicationLifeEvent(getApplicationLifeEvent)
}

//*/

//region router config function
func registerEndpointRouterConfig(router Router.IRouterBuilder) {
	router.GET("/error", func(ctx *Context.HttpContext) {
		panic("http get error")
	})

	router.POST("/info/:id", PostInfo)

	router.Group("/v1/api", func(router *Router.RouterGroup) {
		router.GET("/info", GetInfo)
	})

	router.GET("/info", GetInfo)
	router.GET("/ioc", GetInfoByIOC)
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
	ctx.JSON(200, Context.M{"info": "ok"})
}

func GetInfoByIOC(ctx *Context.HttpContext) {
	var userAction models.IUserAction
	_ = ctx.RequiredServices.GetService(&userAction)
	ctx.JSON(200, Context.M{"info": "ok " + userAction.Login("zhang")})
}

//HttpPost request: /info/:id ?q1=abc&username=123
func PostInfo(ctx *Context.HttpContext) {
	qs_q1 := ctx.Query("q1")
	pd_name := ctx.Param("username")

	userInfo := &UserInfo{}
	_ = ctx.Bind(userInfo)

	strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)

	ctx.JSON(200, Context.M{"info": "hello world", "result": strResult})
}

func getApplicationLifeEvent(life *YoyoGo.ApplicationLife) {
	printDataEvent := func(event YoyoGo.ApplicationEvent) {
		fmt.Printf("[yoyogo] Topic: %s; Event: %v\n", event.Topic, event.Data)
	}

	for {
		select {
		case ev := <-life.ApplicationStarted:
			go printDataEvent(ev)
		case ev := <-life.ApplicationStopped:
			go printDataEvent(ev)
			break
		}
	}
}

//endregion
