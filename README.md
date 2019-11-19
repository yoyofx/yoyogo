# YoyoGo
YoyoGo is a simple, light and fast Web framework written in Go. 

[![Report](https://goreportcard.com/badge/github.com/maxzhang1985/yoyogo)](https://goreportcard.com/report/github.com/maxzhang1985/yoyogo)
![Release](https://img.shields.io/github/v/tag/maxzhang1985/yoyogo.svg?color=24B898&label=release&logo=github&sort=semver)
[![Build Status](https://img.shields.io/travis/maxzhang1985/yoyogo.svg)](https://travis-ci.org/maxzhang1985/yoyogo)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&logo=go&logoColor=ffffff)](https://godoc.org/github.com/maxzhang1985/yoyogo)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
![Contributors](https://img.shields.io/github/contributors/maxzhang1985/yoyogo.svg)



# Features
- Pretty and fast router 
- Middleware Support
- Friendly to REST API
- No regexp or reflect
- Inspired by many excellent Go Web framework

# Installation

`go get github.com/maxzhang1985/yoyogo`


# Example
```golang
func main() {

	webHost := CreateWebHostBuilder(os.Args).Build()
	webHost.Run()

}
//region Create the builder of Web host
func CreateWebHostBuilder(args []string) YoyoGo.HostBuilder {
	return YoyoGo.NewWebHostBuilder().
		UseServer(YoyoGo.DefaultHttpServer(":8080")).
		Configure(func(app *YoyoGo.ApplicationBuilder) {
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
```


# ToDo
## Standard
* [X] Print Logo (YoyoGo)
* [X] Unified program entry (YoyoGo)
* [X] Simple router binded handle func
* [X] HttpContext (request,response)
* [X] Static File EndPoint（Static File Server）
* [X] JSON Serialization Struct （Std.M）
* [X] Get Request File And Save
* [X] Unite Get Post Data Away (form-data , x-www-form-urlencoded)
* [X] Binding Model for Http Request ( From, JSON ) 
### Response Render
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
* [ ] Template
* [ ] Auto formater Render

## Middleware
* [X] Logger
* [X] StaticFile
* [X] Router
* [X] Router Middleware
* [ ] Session
* [ ] CORS
* [X] Binding
* [ ] GZip	


## Router
* [x] GET、POST、HEAD、PUT、DELETE Method Support
* [x] Router Tree
* [x] Router Expression
* [x] RouteData (var)
* [x] Router Support Struct Refect Func Binded.
* [x] Router Support Prefix and Group Such as "/api/v1/endpoint"
* [ ] Router Filter
* [ ] Controller Router And Router Tempalte (Default)

## Dependency injection
* [ ] Dependency injection Framework
* [ ] Dependency injection Integration

## Features
* [ ] swagger
* [ ] configtion
* [ ] WebSocket
* [ ] GRpc
* [ ] JWT 
* [ ] OAuth2	 
* [ ] Prometheus 
* [ ] Secure
* [ ] JWT 
