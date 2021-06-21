package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	nacosconfig "github.com/yoyofx/yoyogo/pkg/configuration/nacos"
	_ "github.com/yoyofx/yoyogo/pkg/datasources/mysql"
	_ "github.com/yoyofx/yoyogo/pkg/datasources/redis"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/nacos"
	web "github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/endpoints"
	"github.com/yoyofx/yoyogo/web/middlewares"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"github.com/yoyofx/yoyogo/web/session"
	"github.com/yoyofx/yoyogo/web/session/identity"
	"github.com/yoyofx/yoyogo/web/session/store"
	"simpleweb/contollers"
	"simpleweb/hubs"
	"simpleweb/models"
)

func SimpleDemo() {
	web.CreateHttpBuilder(func(router router.IRouterBuilder) {
		endpoints.UsePrometheus(router)
		registerEndpointRouterConfig(router)

		router.GET("/info", func(ctx *context.HttpContext) {
			ctx.JSON(200, context.H{"info": "ok"})
		})
	}).Build().Run()
}

func main() {
	//SimpleDemo()

	webHost := CreateCustomBuilder().Build()
	webHost.Run()
}

//* Create the builder of Web host
func CreateCustomBuilder() *abstractions.HostBuilder {
	config := nacosconfig.RemoteConfig("config")
	//config := apollo.RemoteConfig("config")
	return web.NewWebHostBuilder().
		UseConfiguration(config).
		Configure(func(app *web.ApplicationBuilder) {
			app.Use(middlewares.NewSessionWith)
			app.UseMiddleware(middlewares.NewCORS())
			//web.UseMiddleware(middlewares.NewRequestTracker())
			app.UseStaticAssets()
			app.UseEndpoints(registerEndpointRouterConfig)
			//app.SetJsonSerializer(extension.CamelJson())
			app.UseMvc(func(builder *mvc.ControllerBuilder) {
				//builder.AddViews(&view.Option{Path: "./static/templates"})
				builder.AddViewsByConfig()
				builder.AddController(contollers.NewUserController)
				builder.AddController(contollers.NewHubController)
				builder.AddController(contollers.NewDbController)
				builder.AddController(contollers.NewSDController)
				builder.AddFilter("/v1/user/info", &contollers.TestActionFilter{})
			})
		}).
		ConfigureServices(func(serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddTransientByImplements(models.NewUserAction, new(models.IUserAction))
			serviceCollection.AddSingleton(hubs.NewHub) // add websocket hubs

			//eureka.UseServiceDiscovery(serviceCollection)
			//consul.UseServiceDiscovery(serviceCollection)
			nacos.UseServiceDiscovery(serviceCollection)
			//etcd.UseServiceDiscovery(serviceCollection)
			//serviceCollection.AddSingletonByImplements(strategy.NewRandom, new(servicediscovery.Strategy))

			session.UseSession(serviceCollection, func(options *session.Options) {
				options.AddSessionStoreFactory(store.NewRedis)
				//options.AddSessionMemoryStore(store.NewMemory())
				options.AddSessionIdentity(identity.NewCookie())
			})
		}).
		OnApplicationLifeEvent(getApplicationLifeEvent)
}

//*/

//region router config function
func registerEndpointRouterConfig(rb router.IRouterBuilder) {
	endpoints.UseHealth(rb)
	endpoints.UseViz(rb)
	endpoints.UsePrometheus(rb)
	endpoints.UsePprof(rb)
	endpoints.UseReadiness(rb)
	endpoints.UseLiveness(rb)
	endpoints.UseJwt(rb)
	endpoints.UseRouteInfo(rb)

	rb.GET("/error", func(ctx *context.HttpContext) {
		panic("http get error")
	})

	rb.POST("/info/:id", PostInfo)

	rb.Group("/v1/api", func(rg *router.RouterGroup) {
		rg.GET("/info", GetInfo)
	})

	rb.GET("/", GetInfo)

	rb.GET("/info", GetInfo)
	rb.GET("/ioc", GetInfoByIOC)
	rb.GET("/session", TestSession)
	rb.GET("/newsession", SetSession)

}

//endregion

func SetSession(ctx *context.HttpContext) {
	ctx.GetSession().SetValue("user", "yoyogo")
	ctx.JSON(200, context.H{"ok": true})
}

func TestSession(ctx *context.HttpContext) {
	ret := ctx.GetSession().GetString("user")
	ctx.JSON(200, context.H{"user": ret})
}

//region Http Request Methods

type UserInfo struct {
	UserName string `param:"username"`
	Number   string `param:"q1"`
	Id       string `param:"id"`
}

//HttpGet request: /info  or /v1/api/info
//bind UserInfo for id,q1,username
func GetInfo(ctx *context.HttpContext) {
	ctx.JSON(200, context.H{"info": "ok"})
}

func GetInfoByIOC(ctx *context.HttpContext) {
	var userAction models.IUserAction
	_ = ctx.RequiredServices.GetService(&userAction)
	ctx.JSON(200, context.H{"info": "ok " + userAction.Login("zhang")})
}

//bootstrap binding
//HttpPost request: /info/:id ?q1=abc&username=123
func PostInfo(ctx *context.HttpContext) {
	qs_q1 := ctx.Input.Query("q1")
	pd_name := ctx.Input.Param("username")
	id := ctx.Input.Param("id")
	userInfo := &UserInfo{}
	_ = ctx.Bind(userInfo)

	strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s , routeData id:%s", pd_name, qs_q1, userInfo, id)

	ctx.JSON(200, context.H{"info": "hello world", "result": strResult})
}

func getApplicationLifeEvent(life *abstractions.ApplicationLife) {
	printDataEvent := func(event abstractions.ApplicationEvent) {
		xlog.GetXLogger("Application Life Event:").Debug("Topic: %s; Event: %v", event.Topic, event.Data)
		//fmt.Printf("[yoyogo] Topic: %s; Event: %v\n", event.Topic, event.Data)
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
