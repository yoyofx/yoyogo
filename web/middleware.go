package web

import (
	"github.com/yoyofx/yoyogo/web/context"
	"net/http"
)

type MiddlewareHandler interface {
	Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext))
}

type NextFunc func(ctx *context.HttpContext)

type MiddlewareHandlerFunc func(ctx *context.HttpContext, next func(ctx *context.HttpContext))

func (h MiddlewareHandlerFunc) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	h(ctx, next)
}

type middleware struct {
	handler MiddlewareHandler

	// nextfn stores the next.ServeHTTP to reduce memory allocate
	nextfn func(ctx *context.HttpContext)
}

func newMiddleware(handler MiddlewareHandler, next *middleware) middleware {
	return middleware{
		handler: handler,
		nextfn:  next.Invoke,
	}
}

func (m middleware) Invoke(ctx *context.HttpContext) {
	m.handler.Inovke(ctx, m.nextfn)
}

func wrap(handler http.Handler) MiddlewareHandler {
	return MiddlewareHandlerFunc(func(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
		handler.ServeHTTP(ctx.Output.GetWriter(), ctx.Input.GetReader())
		next(ctx)
	})
}

func wrapFunc(handlerFunc http.HandlerFunc) MiddlewareHandler {
	return MiddlewareHandlerFunc(func(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
		handlerFunc(ctx.Output.GetWriter(), ctx.Input.GetReader())
		next(ctx)
	})
}

func voidMiddleware() middleware {
	return newMiddleware(
		MiddlewareHandlerFunc(func(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
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
