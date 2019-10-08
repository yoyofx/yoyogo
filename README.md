# YoyoGo
YoyoGo is a simple, light and fast Web framework written in Go. 

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
package main

import (
	"github.com/maxzhang1985/yoyogo/Framework"
	"github.com/maxzhang1985/yoyogo/Middleware"
)

func main() {

	app := YoyoGo.UseMvc()

	app.Use(Middleware.NewStatic("Static"))
	app.Map("/info", func(ctx *Middleware.HttpContext) {
	    userInfo := &UserInfo{}
            ctx.Bind(userInfo)
            strResult := fmt.Sprintf("Name:%s , Q1:%s , bind: %s", pd_name, qs_q1, userInfo.UserName)
            ctx.JSON(Std.M{"info": "hello world", "result": strResult})
	})

	app.Run(":8080")

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
* [ ] TEXT
* [ ] Protobuf
* [ ] MessagePack
* [ ] XML
* [ ] YAML
* [ ] File
* [ ] Image
* [ ] Template

## Middleware
* [X] Logger
* [X] StaticFile
* [X] Router
* [ ] Router Filter or Middleware (first ！！！)
* [ ] Session
* [ ] CORS
* [X] Binding
* [ ] GZip	
* [ ] JWT 
* [ ] OAuth2	 
* [ ] Prometheus 
* [ ] Secure
* [ ] JWT 

## Router
* [x] GET、POST、HEAD、PUT、DELETE Method Support
* [x] Router Tree
* [x] Router Expression
* [x] RouteData (var)
* [x] Router Support Struct Refect Func Binded.
* [x] Router Support Prefix and Group Such as "/api/v1/endpoint"
* [ ] Controller Router And Router Tempalte (Default)

## Features
* [ ] swagger
* [ ] configtion
* [ ] WebSocket
* [ ] GRpc
