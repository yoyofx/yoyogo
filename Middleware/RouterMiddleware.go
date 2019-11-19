package Middleware

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Router"
	"strings"
)

//var ReqFuncMap = make(map[string]func(ctx * YoyoGo.HttpContext))

type RouterMiddleware struct {
	RouterHandler Router.IRouterHandler
}

func NewRouter(tree Router.IRouterHandler) *RouterMiddleware {
	return &RouterMiddleware{RouterHandler: tree}
}

func NotFoundHandler(ctx *Context.HttpContext) {
	//http.NotFound(ctx.Resp, ctx.Req)
	ctx.Resp.WriteHeader(400)
}

// MethodNotAllowedHandler .
func MethodNotAllowedHandler(ctx *Context.HttpContext) {
	//http.Error(ctx.Resp, "Method Not Allowed", 405)
	ctx.Resp.WriteHeader(405)
}

func (router *RouterMiddleware) Inovke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
	var handler func(ctx *Context.HttpContext)
	handler = router.RouterHandler.Search(ctx, strings.Split(ctx.Req.URL.Path, "/")[1:], ctx.RouterData)
	if handler != nil {
		handler(ctx)
	}
	next(ctx)

}
