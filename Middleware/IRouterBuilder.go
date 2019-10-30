package Middleware

type IRouterBuilder interface {
	Map(method string, path string, handler func(ctx *HttpContext))

	// GET register GET request handler
	GET(path string, handler func(ctx *HttpContext))

	// HEAD register HEAD request handler
	HEAD(path string, handler func(ctx *HttpContext))

	// OPTIONS register OPTIONS request handler
	OPTIONS(path string, handler func(ctx *HttpContext))

	// POST register POST request handler
	POST(path string, handler func(ctx *HttpContext))

	// PUT register PUT request handler
	PUT(path string, handler func(ctx *HttpContext))

	// PATCH register PATCH request HandlerFunc
	PATCH(path string, handler func(ctx *HttpContext))

	// DELETE register DELETE request handler
	DELETE(path string, handler func(ctx *HttpContext))

	// CONNECT register CONNECT request handler
	CONNECT(path string, handler func(ctx *HttpContext))

	// TRACE register TRACE request handler
	TRACE(path string, handler func(ctx *HttpContext))

	// Any register any method handler
	Any(path string, handler func(ctx *HttpContext))

	Group(name string, routerBuilderFunc func(router *RouterGroup))
}
