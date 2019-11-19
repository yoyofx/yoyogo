package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
)

type IRouterBuilder interface {
	Map(method string, path string, handler func(ctx *Context.HttpContext))

	// GET register GET request handler
	GET(path string, handler func(ctx *Context.HttpContext))

	// HEAD register HEAD request handler
	HEAD(path string, handler func(ctx *Context.HttpContext))

	// OPTIONS register OPTIONS request handler
	OPTIONS(path string, handler func(ctx *Context.HttpContext))

	// POST register POST request handler
	POST(path string, handler func(ctx *Context.HttpContext))

	// PUT register PUT request handler
	PUT(path string, handler func(ctx *Context.HttpContext))

	// PATCH register PATCH request HandlerFunc
	PATCH(path string, handler func(ctx *Context.HttpContext))

	// DELETE register DELETE request handler
	DELETE(path string, handler func(ctx *Context.HttpContext))

	// CONNECT register CONNECT request handler
	CONNECT(path string, handler func(ctx *Context.HttpContext))

	// TRACE register TRACE request handler
	TRACE(path string, handler func(ctx *Context.HttpContext))

	// Any register any method handler
	Any(path string, handler func(ctx *Context.HttpContext))

	Group(name string, routerBuilderFunc func(router *RouterGroup))
}
