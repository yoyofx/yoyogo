<img src="https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/yoyogo.png" width = "380px" height = "120px" alt="" align=center />[中文](https://github.com/yoyofx/yoyogo/blob/master/README.md)  / [English](https://github.com/yoyofx/yoyogo/blob/master/README_En.md)

YoyoGo 简单、轻量、快速、基于依赖注入的微服务框架

* 文档： https://dev.yoyogo.run

![Release](https://img.shields.io/github/v/tag/yoyofx/yoyogo.svg?color=24B898&label=release&logo=github&sort=semver)
![Go](https://github.com/yoyofx/yoyogo/workflows/Go/badge.svg)
![GoVersion](https://img.shields.io/github/go-mod/go-version/yoyofx/yoyogo)
[![Report](https://goreportcard.com/badge/github.com/yoyofx/yoyogo)](https://goreportcard.com/report/github.com/yoyofx/yoyogo)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&logo=go&logoColor=ffffff)](https://godoc.org/github.com/yoyofx/yoyogo)
![Contributors](https://img.shields.io/github/contributors/yoyofx/yoyogo.svg)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

# YoyoGo 特色
- 漂亮又快速的路由器 & MVC 模式 .
- 丰富的中间件支持 (handler func & custom middleware) .
- 微服务框架抽象了分层，在一个框架体系兼容各种server实现，如 rest,grpc等 .
- 充分运用依赖注入DI，管理运行时生命周期，为框架提供了强大的扩展性 .
- 功能强大的微服务集成能力 (Nacos, Eureka, Consul, ETCD) .
- 受到许多出色的 Go Web 框架的启发，并实现了多种 server : **fasthttp** 和 **net.http** 和 **grpc** .

![framework desgin](https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/framework-desgin.jpg)

QQ交流群： [780385870](https://qm.qq.com/cgi-bin/qm/qr?k=xP5ZSGZaLLgJIjK0P89gen8V-p5b1cHg&jump_from=webapi) (Go浪文学院) , 在这里感谢**贾国锦**帮忙设计的logo很漂亮。

也可以加入我的公众号，通过公众号入群菜单进入微信群，主要还是在微信上活动啦。

<a href="https://sourcerer.io/yoyofx"><img src="https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/20201209184239.png" width = "180px" height = "180px" alt="" align=center /></a>


# 框架安装
```bash
go get github.com/yoyofx/yoyogo
```
# 安装依赖 (由于某些原因国内下载不了依赖)
##  go version < 1.20
```bash
window 下在 cmd 中执行：
set GO111MODULE=on
set GOPROXY=https://goproxy.cn,direct
linux  下执行：
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
```
##  go version >= 1.13
```
go env -w GOPROXY=https://goproxy.cn,direct
```
### vendor
```
go mod vendor       // 将依赖包拷贝到项目目录中去
```
# 简单的例子
```golang
package main
import ...

func main() {
	WebApplication.CreateDefaultBuilder(func(rb router.IRouterBuilder) {
        rb.GET("/info",func (ctx *context.HttpContext) {    // 支持Group方式
            ctx.JSON(200, context.H{"info": "ok"})
        })
    }).Build().Run()       //默认端口号 :8080
}
```
![](https://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/20201209182340.png)

## 依赖框架模块
只涉及框架使用或贡献过的库
```go
github.com/yoyofxteam/dependencyinjection

github.com/yoyofxteam/nacos-viper-remote  

github.com/yoyofxteam/reflectx

github.com/shima-park/agollo
```
# 实现进度
## 标准功能
* [X] 打印Logo和日志（YoyoGo）
* [X] 统一程序输入参数和环境变量 (YoyoGo)
* [X] 简单路由器绑定句柄功能
* [X] HttpContext 上下文封装(请求，响应)
* [X] 静态文件端点（静态文件服务器）
* [X] JSON 序列化结构（Context.H）
* [X] 获取请求文件并保存
* [X] 获取请求数据（form-data，x-www-form-urlencoded，Json ，XML，Protobuf 等）
* [X] Http 请求的绑定模型（Url, From，JSON，XML，Protobuf）
### 响应渲染功能
* [X] Render Interface
* [X] JSON Render
* [X] JSONP Render
* [X] Indented Json Render
* [X] Secure Json Render
* [X] Ascii Json Render
* [X] Pure Json Render
* [X] Binary Data Render
* [X] TEXT
* [X] Protobuf
* [X] MessagePack
* [X] XML
* [X] YAML
* [X] File
* [X] Image
* [X] Template
* [X] Auto formater Render

## 中间件
* [X] Logger
* [X] StaticFile
* [X] Router Middleware
* [X] CORS	
* [X] Binding
* [X] JWT
* [X] RequestId And Tracker for SkyWorking

## 路由
* [x] GET，POST，HEAD，PUT，DELETE 方法支持
* [x] 路由解析树与表达式支持
* [x] RouteData路由数据 (/api/:version/) 与 Binding的集成 
* [x] 路由组功能
* [x] MVC默认模板功能
* [X] MVC 自定义路由
* [x] 路由过滤器 Filter

## MVC
* [x] 路由请求触发Controller&Action
* [x] Action方法参数绑定
* [x] 内部对象的DI化
* [x] 关键对象的参数传递

## Dependency injection
* [X] 抽象集成第三方DI框架
* [X] MVC模式集成
* [X] 框架级的DI支持功能

## 扩展
* [X] 配置
* [X] WebSocket
* [X] JWT 
* [X] swagger
* [X] GRpc	 
* [X] Prometheus 


# 进阶范例

```golang
package main

import
...

func main() {
	webHost := CreateCustomWebHostBuilder().Build()
	webHost.Run()
}

// 自定义HostBuilder并支持 MVC 和 自动参数绑定功能，简单情况也可以直接使用CreateDefaultBuilder 。
func CreateCustomBuilder() *abstractions.HostBuilder {

	configuration := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	return WebApplication.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *WebApplication.WebApplicationBuilder) {
			app.UseMiddleware(middlewares.NewCORS())
			//WebApplication.UseMiddleware(middlewares.NewRequestTracker())
			app.UseStaticAssets()
			app.UseEndpoints(registerEndpointRouterConfig)
			app.UseMvc(func(builder *mvc.ControllerBuilder) {
				//builder.AddViews(&view.Option{Path: "./static/templates"})
				builder.AddViewsByConfig()
				builder.AddController(contollers.NewUserController)
				builder.AddFilter("/v1/user/info", &contollers.TestActionFilter{})
			})
		}).
		ConfigureServices(func(serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddTransientByImplements(models.NewUserAction, new(models.IUserAction))
			//eureka.UseServiceDiscovery(serviceCollection)
			//consul.UseServiceDiscovery(serviceCollection)
			nacos.UseServiceDiscovery(serviceCollection)
		}).
		OnApplicationLifeEvent(getApplicationLifeEvent)
}

//region endpoint 路由绑定函数
func registerEndpoints(rb router.IRouterBuilder) {
	Endpoints.UseHealth(rb)
	Endpoints.UseViz(rb)
	Endpoints.UsePrometheus(rb)
	Endpoints.UsePprof(rb)
	Endpoints.UseJwt(rb)
	
	//swagger api document
	endpoints.UseSwaggerDoc(rb,
		swagger.Info{
			Title:          "YoyoGO 框架文档演示",
			Version:        "v1.0.0",
			Description:    "框架文档演示swagger文档 v1.0 [ #yoyogo](https://github.com/yoyofx/yoyogo).",
			TermsOfService: "https://dev.yoyogo.run",
			Contact: swagger.Contact{
				Email: "zl.hxd@hotmail.com",
				Name:  "yoyogo",
			},
			License: swagger.License{
				Name: "MIT",
				Url:  "https://opensource.org/licenses/MIT",
			},
		},
		func(openapi *swagger.OpenApi) {
			openapi.AddSecurityBearerAuth()
		})

	rb.GET("/error", func(ctx *context.HttpContext) {
		panic("http get error")
	})

	//POST 请求: /info/:id ?q1=abc&username=123
	rb.POST("/info/:id", func(ctx *context.HttpContext) {
		qs_q1 := ctx.Query("q1")
		pd_name := ctx.Param("username")

		userInfo := &UserInfo{}

		_ = ctx.Bind(userInfo) // 手动绑定请求对象

		strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)

		ctx.JSON(200, context.H{"info": "hello world", "result": strResult})
	})

	// 路由组功能实现绑定 GET 请求:  /v1/api/info
	rb.Group("/v1/api", func(router *router.RouterGroup) {
		router.GET("/info", func(ctx *context.HttpContext) {
			ctx.JSON(200, context.H{"info": "ok"})
		})
	})

	// GET 请求: HttpContext.RequiredServices获取IOC对象
	rb.GET("/ioc", func(ctx *context.HttpContext) {
		var userAction models.IUserAction
		_ = ctx.RequiredServices.GetService(&userAction)
		ctx.JSON(200, context.H{"info": "ok " + userAction.Login("zhang")})
	})
}

//endregion

//region 请求对象
type UserInfo struct {
	UserName string `param:"username"`
	Number   string `param:"q1"`
	Id       string `param:"id"`
}

// ----------------------------------------- MVC 定义 ------------------------------------------------------

// 定义Controller
type UserController struct {
	*mvc.ApiController
	userAction models.IUserAction // IOC 对象参数
}

// 构造器依赖注入
func NewUserController(userAction models.IUserAction) *UserController {
	return &UserController{userAction: userAction}
}

// 请求对象的参数化绑定 , 使用 doc属性标注 支持swagger文档
type RegisterRequest struct {
	mvc.RequestBody `route:"/api/users/register" doc:"用户注册"`
	UserName   string `uri:"userName" doc:"用户名"`
	Password   string `uri:"password" doc:"密码"`
	TestNumber uint64 `uri:"num" doc:"数字"`
}

// Register函数自动绑定参数
func (this *UserController) Register(ctx *context.HttpContext, request *RegiserRequest) actionresult.IActionResult {
	result := mvc.ApiResult{Success: true, Message: "ok", Data: request}
	return actionresult.Json{Data: result}
}

// use userAction interface by ioc  
func (this *UserController) GetInfo() mvc.ApiResult {
	return this.OK(this.userAction.Login("zhang"))
}

// DocumentResponse custom document response , use doc tag for swagger
type DocumentResponse struct {
	Message string        `json:"message" doc:"消息"`
	List    []DocumentDto `json:"list" doc:"文档列表"`
	Success bool          `json:"success" doc:"是否成功"`
}

// Swagger API 文档支持
func (controller UserController) GetDocumentList(request *struct {
	mvc.RequestGET `route:"/v1/user/doc/list" doc:"获取全部文档列表"`
}) DocumentResponse {

	return DocumentResponse{Message: "GetDocumentList", List: []DocumentDto{
		{Id: 1, Name: "test1", Time: time.Now()}, {Id: 2, Name: "test2", Time: time.Now()},
		{Id: 3, Name: "test3", Time: time.Now()}, {Id: 4, Name: "test4", Time: time.Now()},
		{Id: 5, Name: "test5", Time: time.Now()}, {Id: 6, Name: "test6", Time: time.Now()},
	}, Success: true}
}


// Web程序的开始与停止事件
func fireApplicationLifeEvent(life *abstractions.ApplicationLife) {
	printDataEvent := func(event abstractions.ApplicationEvent) {
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

```
