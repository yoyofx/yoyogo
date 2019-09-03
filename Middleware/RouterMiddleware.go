package Middleware

import (
	"net/http"
)

//var ReqFuncMap = make(map[string]func(ctx * YoyoGo.HttpContext))

type RouterMiddleware struct {
	ReqFuncMap map[string]func(ctx *HttpContext)
}

func NewRouter() *RouterMiddleware {
	return &RouterMiddleware{ReqFuncMap: make(map[string]func(ctx *HttpContext))}
}

func NotFoundHandler(ctx *HttpContext) error {
	http.NotFound(ctx.Resp, ctx.Req)
	return nil
}

// MethodNotAllowedHandler .
func MethodNotAllowedHandler(ctx *HttpContext) error {
	http.Error(ctx.Resp, "Method Not Allowed", 405)
	return nil
}

func (router *RouterMiddleware) Inovke(ctx *HttpContext, next func(ctx *HttpContext)) {

	fun, ok := router.ReqFuncMap[ctx.Req.URL.Path]
	if ok {
		fun(ctx)
		next(ctx)
		return
	}
	//else {
	//	_ = NotFoundHandler(ctx)
	//}
	next(ctx)
}
