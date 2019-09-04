# YoyoGo
YoyoGo is a simple, light and fast Web framework written in Go. 

# Features
- Pretty and fast router - based on radix tree
- Middleware Support
- Friendly to REST API
- No regexp or reflect
- Inspired by many excellent Go Web framework

# ToDo
## Standard
* [X] Unified program entry (YoyoGo)
* [X] Simple router binded handle func
* [X] HttpContext (request,response)
* [X] Static File EndPoint（Static File Server）
### Response
* [X] JSON
* [X] JSONP
* [ ] TEXT
* [ ] File
* [ ] Image
* [ ] Binary
* [ ] Other Format

## Middleware

## Router
* [ ] Router Tree
* [ ] Router Expression
* [ ] GET、POST、HEAD、PUT、DELETE Method Support

## Features



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
		ctx.JSON(YoyoGo.M{"info": "hello world"})
	})

	app.Run(":8080")

}
```
