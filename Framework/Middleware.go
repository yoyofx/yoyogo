package YoyoGo

import "net/http"

type Handler interface {
	Inovke(ctx *HttpContext, next func(ctx *HttpContext))
}

type HandlerFunc func(ctx *HttpContext, next func(ctx *HttpContext))

func (h HandlerFunc) Inovke(ctx *HttpContext, next func(ctx *HttpContext)) {
	h(ctx, next)
}

type middleware struct {
	handler Handler

	// nextfn stores the next.ServeHTTP to reduce memory allocate
	nextfn func(ctx *HttpContext)
}

func newMiddleware(handler Handler, next *middleware) middleware {
	return middleware{
		handler: handler,
		nextfn:  next.Invoke,
	}
}

func (m middleware) Invoke(ctx *HttpContext) {
	m.handler.Inovke(ctx, m.nextfn)
}

func wrap(handler http.Handler) Handler {
	return HandlerFunc(func(ctx *HttpContext, next func(ctx *HttpContext)) {
		handler.ServeHTTP(ctx.Resp, ctx.Req)
		next(ctx)
	})
}

func wrapFunc(handlerFunc http.HandlerFunc) Handler {
	return HandlerFunc(func(ctx *HttpContext, next func(ctx *HttpContext)) {
		handlerFunc(ctx.Resp, ctx.Req)
		next(ctx)
	})
}

func voidMiddleware() middleware {
	return newMiddleware(
		HandlerFunc(func(ctx *HttpContext, next func(ctx *HttpContext)) {}),
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
