package router

import (
	"github.com/yoyofx/yoyogo/web/context"
)

type RouterGroup struct {
	Name          string
	RouterHandler IRouterBuilder
}

func (group *RouterGroup) Map(method string, path string, handler func(ctx *context.HttpContext)) {

	routerPath := group.Name + path
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}

	group.RouterHandler.Map(method, routerPath, handler)
}

// GET register GET request handler
func (group *RouterGroup) GET(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.GET, path, handler)
}

// HEAD register HEAD request handler
func (group *RouterGroup) HEAD(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (group *RouterGroup) OPTIONS(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.OPTIONS, path, handler)
}

// POST register POST request handler
func (group *RouterGroup) POST(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.POST, path, handler)
}

// PUT register PUT request handler
func (group *RouterGroup) PUT(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (group *RouterGroup) PATCH(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (group *RouterGroup) DELETE(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (group *RouterGroup) CONNECT(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (group *RouterGroup) TRACE(path string, handler func(ctx *context.HttpContext)) {
	group.Map(context.TRACE, path, handler)
}

// Any register any method handler
func (group *RouterGroup) Any(path string, handler func(ctx *context.HttpContext)) {
	for _, m := range context.Methods {
		group.Map(m, path, handler)
	}
}
