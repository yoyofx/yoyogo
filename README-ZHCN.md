# YoyoGo
YoyoGo 是一个用 Go 编写的简单，轻便，快速的 Web 框架。

![Release](https://img.shields.io/github/v/tag/yoyofx/yoyogo.svg?color=24B898&label=release&logo=github&sort=semver)
[![Build Status](https://img.shields.io/travis/yoyofx/yoyogo.svg)](https://travis-ci.org/yoyofx/yoyogo)
[![Goproxy](https://goproxy.cn/stats/github.com/yoyofx/yoyogo/badges/download-count.svg)](https://goproxy.cn/stats/github.com/yoyofx/yoyogo)
![GoVersion](https://img.shields.io/github/go-mod/go-version/maxzhang1985/yoyogo)
[![GoBadge](https://img.shields.io/badge/Go-197%20commits-orange.svg)](https://sourcerer.io/yoyofx)
![DockerPull](https://img.shields.io/docker/pulls/maxzhang1985/yoyogo)
[![Report](https://goreportcard.com/badge/github.com/yoyofx/yoyogo)](https://goreportcard.com/report/github.com/maxzhang1985/yoyogo)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&logo=go&logoColor=ffffff)](https://godoc.org/github.com/yoyofx/yoyogo)
![Contributors](https://img.shields.io/github/contributors/yoyofx/yoyogo.svg)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

# 特色
- 漂亮又快速的路由器
- 中间件支持 (handler func & custom middleware)
- 对 REST API 友好
- 没有正则表达式
- 受到许多出色的 Go Web 框架的启发

[![](https://avatars3.githubusercontent.com/u/4504853?v=4)](https://sourcerer.io/yoyofx)
# 框架安装
```bash
go get github.com/yoyofx/yoyogo
```
# 安装依赖 (由于某些原因国内下载不了依赖)
##  go version < 1.13
```bash
window 下在 cmd 中执行：
set GO111MODULE=on
set  GOPROXY=https://goproxy.cn

linux  下执行：
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
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
    YoyoGo.CreateDefaultBuilder(func(router Router.IRouterBuilder) {
        router.GET("/info",func (ctx *Context.HttpContext) {    // 支持Group方式
            ctx.JSON(200, Context.M{"info": "ok"})
        })
    }).Build().Run()       //默认端口号 :8080
}
```
![](./yoyorun.jpg)


# 实现进度
## 标准功能
* [X] 打印Logo和日志（YoyoGo）
* [X] 统一程序输入参数和环境变量 (YoyoGo)
* [X] 简单路由器绑定句柄功能
* [X] HttpContext 上下文封装(请求，响应)
* [X] 静态文件端点（静态文件服务器）
* [X] JSON 序列化结构（Context.M）
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
* [ ] Session
* [ ] CORS
* [ ] GZip	
* [X] Binding
* [ ] Binding Valateion



## 路由
* [x] GET，POST，HEAD，PUT，DELETE 方法支持
* [x] 路由解析树与表达式支持
* [x] RouteData路由数据 (/api/:version/) 与 Binding的集成 
* [x] 路由组功能
* [ ] MVC默认模板功能
* [ ] 路由过滤器 Filter

## MVC
* [x] 路由请求触发Controller&Action
* [X] Action方法参数绑定
* [ ] 内部对象的DI化
* [ ] 关键对象的参数传递

## Dependency injection
* [X] 抽象集成第三方DI框架
* [X] MVC模式集成
* [ ] 框架级的DI支持功能

## 扩展
* [ ] 配置
* [ ] WebSocket
* [ ] JWT 
* [ ] swagger
* [ ] GRpc
* [ ] OAuth2	 
* [ ] Prometheus 
* [ ] 安全


# 进阶范例
```golang
package main
import ...

func main() {
	webHost := CreateCustomWebHostBuilder().Build()
	webHost.Run()
}

// 自定义HostBuilder并支持 MVC 和 自动参数绑定功能，简单情况也可以直接使用CreateDefaultBuilder 。
func CreateCustomBuilder() *YoyoGo.HostBuilder {
	return YoyoGo.NewWebHostBuilder().
		UseFastHttp().         //Server可以指定多种，这里使用FastHttp作为Server，后期也会实现gRPC和WebSocket
		                       //使用默认Server并指定协议和端口 UseServer(YoyoGo.DefaultHttps(":8080", "./Certificate/server.pem", "./Certificate/server.key")).
		Configure(func(app *YoyoGo.ApplicationBuilder) {
			app.SetEnvironment(Context.Dev)
			app.UseStatic("Static")
			app.UseEndpoints(registerEndpoints)             // endpoint 路由绑定函数
			app.UseMvc()                                    // 开启MVC功能
			app.ConfigureMvcParts(func(builder *Controller.ControllerBuilder) {
				builder.AddController(contollers.NewUserController)
			})
		}).
		ConfigureServices(func(serviceCollection *DependencyInjection.ServiceCollection) {      // 依赖注入方法
			serviceCollection.AddTransientByImplements(models.NewUserAction, new(models.IUserAction))
		}).
		OnApplicationLifeEvent(getApplicationLifeEvent)
}

//region endpoint 路由绑定函数
func registerEndpoints(router Router.IRouterBuilder) {
	router.GET("/error", func(ctx *Context.HttpContext) {
		panic("http get error")
	})

    //POST 请求: /info/:id ?q1=abc&username=123
	router.POST("/info/:id", func (ctx *Context.HttpContext) {
        qs_q1 := ctx.Query("q1")
        pd_name := ctx.Param("username")

        userInfo := &UserInfo{}
        
        _ = ctx.Bind(userInfo)    // 手动绑定请求对象

        strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo)

        ctx.JSON(200, Std.M{"info": "hello world", "result": strResult})
    })

    // 路由组功能实现绑定 GET 请求:  /v1/api/info
	router.Group("/v1/api", func(router *Router.RouterGroup) {
		router.GET("/info", func (ctx *Context.HttpContext) {
	        ctx.JSON(200, Std.M{"info": "ok"})
        })
	})
    
    // GET 请求: HttpContext.RequiredServices获取IOC对象
	router.GET("/ioc", func (ctx *Context.HttpContext) {
        var userAction models.IUserAction
        _ = ctx.RequiredServices.GetService(&userAction)
        ctx.JSON(200, Std.M{"info": "ok " + userAction.Login("zhang")})
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
	*Controller.ApiController
	userAction models.IUserAction    // IOC 对象参数
}

// 构造器依赖注入
func NewUserController(userAction models.IUserAction) *UserController {
	return &UserController{userAction: userAction}
}

// 请求对象的参数化绑定
type RegiserRequest struct {
	Controller.RequestParam
	UserName string `param:"username"`
	Password string `param:"password"`
}

// Register函数自动绑定参数
func (this *UserController) Register(ctx *Context.HttpContext, request *RegiserRequest) ActionResult.IActionResult {
	result := Controller.ApiResult{Success: true, Message: "ok", Data: request}
	return ActionResult.Json{Data: result}
}

// use userAction interface by ioc  
func (this *UserController) GetInfo() Controller.ApiResult {
	return this.OK(this.userAction.Login("zhang"))
}


// Web程序的开始与停止事件
func fireApplicationLifeEvent(life *YoyoGo.ApplicationLife) {
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

```