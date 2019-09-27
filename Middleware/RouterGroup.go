package Middleware

import Std "github.com/maxzhang1985/yoyogo/Standard"

type RouterGroup struct {
	Name       string
	RouterTree *Trie
}

func (group *RouterGroup) Map(method string, path string, handler func(ctx *HttpContext)) {

	routerPath := group.Name + path
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}

	group.RouterTree.Insert(method, routerPath, handler)
}

// GET register GET request handler
func (group *RouterGroup) GET(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.GET, path, handler)
}

// HEAD register HEAD request handler
func (group *RouterGroup) HEAD(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (group *RouterGroup) OPTIONS(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.OPTIONS, path, handler)
}

// POST register POST request handler
func (group *RouterGroup) POST(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.POST, path, handler)
}

// PUT register PUT request handler
func (group *RouterGroup) PUT(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (group *RouterGroup) PATCH(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (group *RouterGroup) DELETE(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (group *RouterGroup) CONNECT(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (group *RouterGroup) TRACE(path string, handler func(ctx *HttpContext)) {
	group.Map(Std.TRACE, path, handler)
}

// Any register any method handler
func (group *RouterGroup) Any(path string, handler func(ctx *HttpContext)) {
	for _, m := range Std.Methods {
		group.Map(m, path, handler)
	}
}
