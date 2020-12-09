package router

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"net/url"
)

type IRouterBuilder interface {
	IsMvc() bool

	UseMvc(used bool)

	GetMvcBuilder() *mvc.ControllerBuilder

	Map(method string, path string, handler func(ctx *context.HttpContext))

	Search(ctx *context.HttpContext, components []string, params url.Values) func(ctx *context.HttpContext)

	// GET register GET request handler
	GET(path string, handler func(ctx *context.HttpContext))

	// HEAD register HEAD request handler
	HEAD(path string, handler func(ctx *context.HttpContext))

	// OPTIONS register OPTIONS request handler
	OPTIONS(path string, handler func(ctx *context.HttpContext))

	// POST register POST request handler
	POST(path string, handler func(ctx *context.HttpContext))

	// PUT register PUT request handler
	PUT(path string, handler func(ctx *context.HttpContext))

	// PATCH register PATCH request HandlerFunc
	PATCH(path string, handler func(ctx *context.HttpContext))

	// DELETE register DELETE request handler
	DELETE(path string, handler func(ctx *context.HttpContext))

	// CONNECT register CONNECT request handler
	CONNECT(path string, handler func(ctx *context.HttpContext))

	// TRACE register TRACE request handler
	TRACE(path string, handler func(ctx *context.HttpContext))

	// Any register any method handler
	Any(path string, handler func(ctx *context.HttpContext))

	Group(name string, routerBuilderFunc func(router *RouterGroup))

	SetConfiguration(config abstractions.IConfiguration)

	GetConfiguration() abstractions.IConfiguration
}
