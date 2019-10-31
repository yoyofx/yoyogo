package Middleware

import (
	"strings"
)

//var ReqFuncMap = make(map[string]func(ctx * YoyoGo.HttpContext))

type RouterMiddleware struct {
	ReqFuncMap map[string]func(ctx *HttpContext)
	Tree       *Trie
}

func NewRouter() *RouterMiddleware {
	tree := &Trie{
		Component: "/",
		Methods:   make(map[string]func(ctx *HttpContext)),
	}

	return &RouterMiddleware{ReqFuncMap: make(map[string]func(ctx *HttpContext)), Tree: tree}
}

func NotFoundHandler(ctx *HttpContext) {
	//http.NotFound(ctx.Resp, ctx.Req)
	ctx.Resp.WriteHeader(400)
}

// MethodNotAllowedHandler .
func MethodNotAllowedHandler(ctx *HttpContext) {
	//http.Error(ctx.Resp, "Method Not Allowed", 405)
	ctx.Resp.WriteHeader(405)
}

func (router *RouterMiddleware) Inovke(ctx *HttpContext, next func(ctx *HttpContext)) {
	var handler func(ctx *HttpContext)
	node := router.Tree.Search(strings.Split(ctx.Req.URL.Path, "/")[1:], ctx.RouterData)
	if node != nil && node.Methods[ctx.Req.Method] != nil {
		handler = node.Methods[ctx.Req.Method]
		handler(ctx)
		return
	} else if node != nil && node.Methods[ctx.Req.Method] == nil {
		handler = MethodNotAllowedHandler
	} else {
		handler = NotFoundHandler
	}
	next(ctx)

}
