package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"net/url"
	"strings"
)

type DefaultRouterBuilder struct {
	mvcRouterHandler      *MvcRouterHandler
	endPointRouterHandler *EndPointRouterHandler
}

func NewRouterBuilder() IRouterBuilder {
	endpoint := &EndPointRouterHandler{
		Component: "/",
		Methods:   make(map[string]func(ctx *Context.HttpContext)),
	}

	defaultRouterHandler := &DefaultRouterBuilder{endPointRouterHandler: endpoint}

	return defaultRouterHandler
}

func (router *DefaultRouterBuilder) SetMvc(used bool) {
	if used {
		router.mvcRouterHandler = &MvcRouterHandler{}
	} else {
		router.mvcRouterHandler = nil
	}
}

func (router *DefaultRouterBuilder) IsMvc() bool {
	return router.mvcRouterHandler != nil
}

func (router *DefaultRouterBuilder) Search(ctx *Context.HttpContext, components []string, params url.Values) func(ctx *Context.HttpContext) {
	var handler func(ctx *Context.HttpContext) = nil

	handler = router.endPointRouterHandler.Invoke(ctx, strings.Split(ctx.Req.URL.Path, "/")[1:])

	if handler == nil && router.IsMvc() {
		handler = router.mvcRouterHandler.Invoke(ctx, strings.Split(ctx.Req.URL.Path, "/")[1:])
	}

	return handler
}

func (router *DefaultRouterBuilder) MapSet(method, path string, handler func(ctx *Context.HttpContext)) {
	router.endPointRouterHandler.Insert(method, path, handler)
}

func (router *DefaultRouterBuilder) Map(method string, path string, handler func(ctx *Context.HttpContext)) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}
	router.MapSet(method, path, handler)
}

// GET register GET request handler
func (router *DefaultRouterBuilder) GET(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.GET, path, handler)
}

// HEAD register HEAD request handler
func (router *DefaultRouterBuilder) HEAD(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (router *DefaultRouterBuilder) OPTIONS(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.OPTIONS, path, handler)
}

// POST register POST request handler
func (router *DefaultRouterBuilder) POST(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.POST, path, handler)
}

// PUT register PUT request handler
func (router *DefaultRouterBuilder) PUT(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (router *DefaultRouterBuilder) PATCH(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (router *DefaultRouterBuilder) DELETE(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (router *DefaultRouterBuilder) CONNECT(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (router *DefaultRouterBuilder) TRACE(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.TRACE, path, handler)
}

// Any register any method handler
func (router *DefaultRouterBuilder) Any(path string, handler func(ctx *Context.HttpContext)) {
	for _, m := range Context.Methods {
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
