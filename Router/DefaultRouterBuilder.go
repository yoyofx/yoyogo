package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	Std "github.com/maxzhang1985/yoyogo/Standard"
	"net/url"
	"strings"
)

type DefaultRouterBuilder struct {
	UseMvc     bool
	routerTree *Trie
}

func (router *DefaultRouterBuilder) SetMvc(used bool) {
	router.UseMvc = used
}

func (router *DefaultRouterBuilder) IsMvc() bool {
	return router.UseMvc
}

func (router *DefaultRouterBuilder) Search(ctx *Context.HttpContext, components []string, params url.Values) func(ctx *Context.HttpContext) {
	var handler func(ctx *Context.HttpContext)
	node := router.routerTree.Search(strings.Split(ctx.Req.URL.Path, "/")[1:], ctx.RouterData)
	if node != nil && node.Methods[ctx.Req.Method] != nil {
		handler = node.Methods[ctx.Req.Method]
	} else if node != nil && node.Methods[ctx.Req.Method] == nil {
		//handler = MethodNotAllowedHandler
	} else {
		//handler = NotFoundHandler
	}
	return handler
}

func (router *DefaultRouterBuilder) MapSet(method, path string, handler func(ctx *Context.HttpContext)) {
	router.routerTree.Insert(method, path, handler)
}

func (router *DefaultRouterBuilder) Map(method string, path string, handler func(ctx *Context.HttpContext)) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}
	router.MapSet(method, path, handler)
}

// GET register GET request handler
func (router *DefaultRouterBuilder) GET(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.GET, path, handler)
}

// HEAD register HEAD request handler
func (router *DefaultRouterBuilder) HEAD(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (router *DefaultRouterBuilder) OPTIONS(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.OPTIONS, path, handler)
}

// POST register POST request handler
func (router *DefaultRouterBuilder) POST(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.POST, path, handler)
}

// PUT register PUT request handler
func (router *DefaultRouterBuilder) PUT(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (router *DefaultRouterBuilder) PATCH(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (router *DefaultRouterBuilder) DELETE(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (router *DefaultRouterBuilder) CONNECT(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (router *DefaultRouterBuilder) TRACE(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Std.TRACE, path, handler)
}

// Any register any method handler
func (router *DefaultRouterBuilder) Any(path string, handler func(ctx *Context.HttpContext)) {
	for _, m := range Std.Methods {
		router.Map(m, path, handler)
	}
}

func (router *DefaultRouterBuilder) Group(name string, routerBuilderFunc func(router *RouterGroup)) {
	group := &RouterGroup{
		Name:          name,
		RouterHandler: router,
	}
	if routerBuilderFunc == nil {
		panic("routerBuilderFunc is nil")
	}

	routerBuilderFunc(group)
}