# YoyoGo [英文介绍](https://github.com/yoyofx/yoyogo/blob/master/README_En.md "中文介绍")
YoyoGo 一个简单、轻量、快速、基于依赖注入的微服务框架

* 文档： https://github.com/yoyofx/yoyogo/wiki

![Release](https://img.shields.io/github/v/tag/yoyofx/yoyogo.svg?color=24B898&label=release&logo=github&sort=semver)
![Go](https://github.com/yoyofx/yoyogo/workflows/Go/badge.svg)
![GoVersion](https://img.shields.io/github/go-mod/go-version/maxzhang1985/yoyogo)
![DockerPull](https://img.shields.io/docker/pulls/maxzhang1985/yoyogo)
[![Report](https://goreportcard.com/badge/github.com/yoyofx/yoyogo)](https://goreportcard.com/report/github.com/maxzhang1985/yoyogo)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&logo=go&logoColor=ffffff)](https://godoc.org/github.com/yoyofx/yoyogo)
![Contributors](https://img.shields.io/github/contributors/yoyofx/yoyogo.svg)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

# 特色
- 漂亮又快速的路由器
- 中间件支持 (handler func & custom middleware)
- 微服务框架抽象了分层，在一个框架体系兼容各种server实现，如 rest,grpc等
- 受到许多出色的 Go Web 框架的启发，server可替换，目前实现了 **fasthttp** 和 **net.http**

[![](Resources/dingdingQR.jpg)](https://sourcerer.io/yoyofx)

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
![](Resources/yoyorun.jpg)


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
* [x] MVC默认模板功能
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
* [ ] WebSocket
* [ ] JWT 
* [ ] swagger
* [ ] GRpc
* [ ] OAuth2	 
* [X] Prometheus 
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
func CreateCustomBuilder() *Abstractions.HostBuilder {
    return YoyoGo.NewWebHostBuilder().
        SetEnvironment(Context.Prod).
        UseFastHttp().
        //UseServer(YoyoGo.DefaultHttps(":8080", "./Certificate/server.pem", "./Certificate/server.key")).
        Configure(func(app *YoyoGo.WebApplicationBuilder) {
            app.UseStatic("Static")
            app.UseEndpoints(registerEndpointRouterConfig)
            app.UseMvc(func(builder *Mvc.ControllerBuilder) {
                builder.AddController(contollers.NewUserController)
            })
        }).
        ConfigureServices(func(serviceCollection *DependencyInjection.ServiceCollection) {
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