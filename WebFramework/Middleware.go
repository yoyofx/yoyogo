package YoyoGo

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net/http"
)

type MiddlewareHandler interface {
	Invoke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext))
}

type NextFunc func(ctx *Context.HttpContext)

type MiddlewareHandlerFunc func(ctx *Context.HttpContext, next func(ctx *Context.HttpContext))

func (h MiddlewareHandlerFunc) Invoke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
	h(ctx, next)
}

type middleware struct {
	handler MiddlewareHandler

	// nextfn stores the next.ServeHTTP to reduce memory allocate
	nextfn func(ctx *Context.HttpContext)
}

func newMiddleware(handler MiddlewareHandler, next *middleware) middleware {
	return middleware{
		handler: handler,
		nextfn:  next.Invoke,
	}
}

func (m middleware) Invoke(ctx *Context.HttpContext) {
	m.handler.Invoke(ctx, m.nextfn)
}

func wrap(handler http.Handler) MiddlewareHandler {
	return MiddlewareHandlerFunc(func(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
		handler.ServeHTTP(ctx.Output.GetWriter(), ctx.Input.GetReader())
		next(ctx)
	})
}

func wrapFunc(handlerFunc http.HandlerFunc) MiddlewareHandler {
	return MiddlewareHandlerFunc(func(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
		handlerFunc(ctx.Output.GetWriter(), ctx.Input.GetReader())
		next(ctx)
	})
}

func voidMiddleware() middleware {
	return newMiddleware(
		MiddlewareHandlerFunc(func(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
			if ctx.Output.Status() == 0 {
				ctx.Output.SetStatus(404)
			}
		}),
		&middleware{},
	)
}

func build(handlers []MiddlewareHandler) middleware {
	var next middleware

	switch {
	case len(handlers) == 0:
		return voidMiddleware()
	case len(handlers) > 1:
		next = build(handlers[1:])
	default:
		next = voidMiddleware()
	}

	return newMiddleware(handlers[0], &next)
}
