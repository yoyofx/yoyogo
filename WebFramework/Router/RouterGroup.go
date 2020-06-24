package Router

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type RouterGroup struct {
	Name          string
	RouterHandler IRouterBuilder
}

func (group *RouterGroup) Map(method string, path string, handler func(ctx *Context.HttpContext)) {

	routerPath := group.Name + path
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}

	group.RouterHandler.Map(method, routerPath, handler)
}

// GET register GET request handler
func (group *RouterGroup) GET(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.GET, path, handler)
}

// HEAD register HEAD request handler
func (group *RouterGroup) HEAD(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (group *RouterGroup) OPTIONS(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.OPTIONS, path, handler)
}

// POST register POST request handler
func (group *RouterGroup) POST(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.POST, path, handler)
}

// PUT register PUT request handler
func (group *RouterGroup) PUT(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (group *RouterGroup) PATCH(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (group *RouterGroup) DELETE(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (group *RouterGroup) CONNECT(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (group *RouterGroup) TRACE(path string, handler func(ctx *Context.HttpContext)) {
	group.Map(Context.TRACE, path, handler)
}

// Any register any method handler
func (group *RouterGroup) Any(path string, handler func(ctx *Context.HttpContext)) {
	for _, m := range Context.Methods {
		group.Map(m, path, handler)
	}
}
