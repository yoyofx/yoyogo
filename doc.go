// YoyoGo is a simple, light and fast Web framework written in Go.
// As Features:
// 		Pretty and fast router
// 		Middleware Support
// 		Friendly to REST API
// 		No regexp or reflect
// 		Inspired by many excellent Go Web framework
// As a quick start:
//		package main
//		import ...
//		func main() {
//			//webHost := YoyoGo.CreateDefaultWebHostBuilder(os.Args,RouterConfigFunc).Build()
//			webHost := CreateCustomWebHostBuilder(os.Args).Build()
//			webHost.Run()
//		}
//
//		//* Create the builder of Web host
//		func CreateCustomWebHostBuilder(args []string) *YoyoGo.HostBuilder {
//			return YoyoGo.NewWebHostBuilder().
//				UseFastHttp().   //default port:8080
//				//UseServer(YoyoGo.DefaultHttps(":8080", "./Certificate/server.pem", "./Certificate/server.key")).
//				Configure(func(app *YoyoGo.ApplicationBuilder) {
//					app.SetEnvironment(Context.Prod)   //set prod to environment
//					app.UseStatic("Static")
//				}).
//				UseEndpoints(RouterConfigFunc).
//				ConfigureServices(func(serviceCollection *DependencyInjection.ServiceCollection) {
//					serviceCollection.AddTransientByImplements(models.NewUserAction, new(models.IUserAction))
//				}).
//				OnApplicationLifeEvent(fireApplicationLifeEvent)
//		}
//
//		//*/
//		//region router config function
//		func RouterConfigFunc(router Router.IRouterBuilder) {
//			router.GET("/error", func(ctx *Context.HttpContext) {
//				panic("http get error")
//			})
//
//			router.POST("/info/:id", PostInfo)
//
//			router.Group("/v1/api", func(router *Router.RouterGroup) {
//				router.GET("/info", GetInfo)
//			})
//
//			router.GET("/info", GetInfo)
//			router.GET("/ioc", GetInfoByIOC)
//		}
//
//		//endregion
//
//		//region Http Request Methods
//
//		type UserInfo struct {
//			UserName string `param:"username"`
//			Number   string `param:"q1"`
//			Id       string `param:"id"`
//		}
//
//		//HttpGet request: /info  or /v1/api/info
//		//bind UserInfo for id,q1,username
//		func GetInfo(ctx *Context.HttpContext) {
//			ctx.JSON(200, Std.M{"info": "ok"})
//		}
//
//		func GetInfoByIOC(ctx *Context.HttpContext) {
//			var userAction models.IUserAction
//			_ = ctx.RequiredServices.GetService(&userAction)
//			ctx.JSON(200, Std.M{"info": "ok " + userAction.Login("zhang")})
//		}
//
//		//HttpPost request: /info/:id ?q1=abc&username=123
//		func PostInfo(ctx *Context.HttpContext) {
//			qs_q1 := ctx.Query("q1")
//			pd_name := ctx.Param("username")
//
//			userInfo := &UserInfo{}
//			_ = ctx.Bind(userInfo)
//
//			strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)
//
//			ctx.JSON(200, Std.M{"info": "hello world", "result": strResult})
//		}
//
//		func fireApplicationLifeEvent(life *YoyoGo.ApplicationLife) {
//			printDataEvent := func(event YoyoGo.ApplicationEvent) {
//				fmt.Printf("[yoyogo] Topic: %s; Event: %v\n", event.Topic, event.Data)
//			}
//			for {
//				select {
//				case ev := <-life.ApplicationStarted:
//					go printDataEvent(ev)
//				case ev := <-life.ApplicationStopped:
//					go printDataEvent(ev)
//					break
//				}
//			}
//		}
//

package YoyoGo
