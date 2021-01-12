package middlewares

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

type Middleware struct {
	handler MiddlewareHandler

	// nextfn stores the next.ServeHTTP to reduce memory allocate
	nextfn func(ctx *context.HttpContext)
}

func newMiddleware(handler MiddlewareHandler, next *Middleware) Middleware {
	return Middleware{
		handler: handler,
		nextfn:  next.Invoke,
	}
}

func (m Middleware) Invoke(ctx *context.HttpContext) {
	m.handler.Inovke(ctx, m.nextfn)
}

func Wrap(handler http.Handler) MiddlewareHandler {
	return MiddlewareHandlerFunc(func(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
		handler.ServeHTTP(ctx.Output.GetWriter(), ctx.Input.GetReader())
		next(ctx)
	})
}

func WrapFunc(handlerFunc http.HandlerFunc) MiddlewareHandler {
	return MiddlewareHandlerFunc(func(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
		handlerFunc(ctx.Output.GetWriter(), ctx.Input.GetReader())
		next(ctx)
	})
}

func voidMiddleware() Middleware {
	return newMiddleware(
		MiddlewareHandlerFunc(func(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
			if ctx.Output.Status() == 0 {
				ctx.Output.SetStatus(404)
			}
		}),
		&Middleware{},
	)
}

func Build(handlers []MiddlewareHandler) Middleware {
	var next Middleware

	switch {
	case len(handlers) == 0:
		return voidMiddleware()
	case len(handlers) > 1:
		next = Build(handlers[1:])
	default:
		next = voidMiddleware()
	}

	return newMiddleware(handlers[0], &next)
}
