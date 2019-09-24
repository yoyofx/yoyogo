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
        _ = ctx.Bind(userInfo)
        ctx.JSON(YoyoGo.M{"info": "hello world"})
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
* [X] JSON Serialization Struct （YoyoGo.M）
* [X] Get Request File And Save
* [X] Unite Get Post Data Away (form-data , x-www-form-urlencoded)
* [X] Binding Model for Http Request ( From, JSON ) 
### Response
* [X] JSON
* [X] JSONP
* [X] TEXT
* [X] Binary
* [ ] XML
* [ ] File
* [ ] Image
* [ ] Other Format
* [ ] Render View (Template)

## Middleware
* [X] Logger
* [X] StaticFile
* [X] Router
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
* [ ] GET、POST、HEAD、PUT、DELETE Method Support
* [ ] Router Tree
* [ ] Router Expression
* [ ] RouteData (var)
* [ ] Router Support Struct Refect Func Binded.
* [ ] Router Support Prefix and Group Such as "/api/v1/endpoint"

## Features
* [ ] swagger
* [ ] configtion
* [ ] WebSocket
* [ ] GRpc
