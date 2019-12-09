package YoyoGo

//
//import (
//	"github.com/maxzhang1985/yoyogo/Context"
//	"github.com/maxzhang1985/yoyogo/Router"
//	"github.com/maxzhang1985/yoyogo/Standard"
//)
//
///*
//IRouterBuilder by applicationBuilder imp
//*/
//
//func (yoyo *ApplicationBuilder) Map(method string, path string, handler func(ctx *Context.HttpContext)) {
//	if len(path) < 1 || path[0] != '/' {
//		panic("Path should be like '/yoyo/go'")
//	}
//	yoyo.routerBuilder.Map(method, path, handler)
//}
//
//// GET register GET request handler
//func (yoyo *ApplicationBuilder) GET(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.GET, path, handler)
//}
//
//// HEAD register HEAD request handler
//func (yoyo *ApplicationBuilder) HEAD(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.HEAD, path, handler)
//}
//
//// OPTIONS register OPTIONS request handler
//func (yoyo *ApplicationBuilder) OPTIONS(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.OPTIONS, path, handler)
//}
//
//// POST register POST request handler
//func (yoyo *ApplicationBuilder) POST(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.POST, path, handler)
//}
//
//// PUT register PUT request handler
//func (yoyo *ApplicationBuilder) PUT(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.PUT, path, handler)
//}
//
//// PATCH register PATCH request HandlerFunc
//func (yoyo *ApplicationBuilder) PATCH(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.PATCH, path, handler)
//}
//
//// DELETE register DELETE request handler
//func (yoyo *ApplicationBuilder) DELETE(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.DELETE, path, handler)
//}
//
//// CONNECT register CONNECT request handler
//func (yoyo *ApplicationBuilder) CONNECT(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.CONNECT, path, handler)
//}
//
//// TRACE register TRACE request handler
//func (yoyo *ApplicationBuilder) TRACE(path string, handler func(ctx *Context.HttpContext)) {
//	yoyo.Map(Std.TRACE, path, handler)
//}
//
//// Any register any method handler
//func (yoyo *ApplicationBuilder) Any(path string, handler func(ctx *Context.HttpContext)) {
//	for _, m := range Std.Methods {
//		yoyo.Map(m, path, handler)
//	}
//}
//
//func (yoyo *ApplicationBuilder) Group(name string, routerBuilderFunc func(router *Router.RouterGroup)) {
//	group := &Router.RouterGroup{
//		Name:          name,
//		RouterBuilder: yoyo.routerBuilder,
//	}
//	if routerBuilderFunc == nil {
//		panic("routerBuilderFunc is nil")
//	}
//
//	routerBuilderFunc(group)
//}
