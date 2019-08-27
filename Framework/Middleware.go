package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/Middleware"
	"net/http"
)

type Handler interface {
	Inovke(ctx *Middleware.HttpContext, next func(ctx *Middleware.HttpContext))
}

type NextFunc func(ctx *Middleware.HttpContext)

type HandlerFunc func(ctx *Middleware.HttpContext, next func(ctx *Middleware.HttpContext))

func (h HandlerFunc) Inovke(ctx *Middleware.HttpContext, next func(ctx *Middleware.HttpContext)) {
	h(ctx, next)
}

type middleware struct {
	handler Handler

	// nextfn stores the next.ServeHTTP to reduce memory allocate
	nextfn func(ctx *Middleware.HttpContext)
}

func newMiddleware(handler Handler, next *middleware) middleware {
	return middleware{
		handler: handler,
		nextfn:  next.Invoke,
	}
}

func (m middleware) Invoke(ctx *Middleware.HttpContext) {
	m.handler.Inovke(ctx, m.nextfn)
}

func wrap(handler http.Handler) Handler {
	return HandlerFunc(func(ctx *Middleware.HttpContext, next func(ctx *Middleware.HttpContext)) {
		handler.ServeHTTP(ctx.Resp, ctx.Req)
		next(ctx)
	})
}

func wrapFunc(handlerFunc http.HandlerFunc) Handler {
	return HandlerFunc(func(ctx *Middleware.HttpContext, next func(ctx *Middleware.HttpContext)) {
		handlerFunc(ctx.Resp, ctx.Req)
		next(ctx)
	})
}

func voidMiddleware() middleware {
	return newMiddleware(
		HandlerFunc(func(ctx *Middleware.HttpContext, next func(ctx *Middleware.HttpContext)) {}),
		&middleware{},
	)
}

func build(handlers []Handler) middleware {
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
