package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"net/url"
	"strings"
)

type DefaultRouterHandler struct {
	routerTree *Trie
}

func (d DefaultRouterHandler) Map(method, path string, handler func(ctx *Context.HttpContext)) {
	d.routerTree.Insert(method, path, handler)
}

func (d DefaultRouterHandler) Search(ctx *Context.HttpContext, components []string, params url.Values) func(ctx *Context.HttpContext) {
	var handler func(ctx *Context.HttpContext)
	node := d.routerTree.Search(strings.Split(ctx.Req.URL.Path, "/")[1:], ctx.RouterData)
	if node != nil && node.Methods[ctx.Req.Method] != nil {
		handler = node.Methods[ctx.Req.Method]
	} else if node != nil && node.Methods[ctx.Req.Method] == nil {
		//handler = MethodNotAllowedHandler
	} else {
		//handler = NotFoundHandler
	}
	return handler
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
