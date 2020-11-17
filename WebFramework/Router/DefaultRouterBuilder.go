package Router

import (
	"github.com/prometheus/common/log"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"net/url"
	"strings"
)

type DefaultRouterBuilder struct {
	mvcControllerBuilder  *Mvc.ControllerBuilder
	endPointRouterHandler *EndPointRouterHandler
	configuration         Abstractions.IConfiguration
	log                   XLog.ILogger
}

func NewRouterBuilder() IRouterBuilder {

	endpoint := &EndPointRouterHandler{
		Component: "/",
		Methods:   make(map[string]func(ctx *Context.HttpContext)),
	}

	defaultRouterHandler := &DefaultRouterBuilder{endPointRouterHandler: endpoint}
	defaultRouterHandler.log = XLog.GetXLogger("DefaultRouterBuilder")
	return defaultRouterHandler
}

func (router *DefaultRouterBuilder) SetConfiguration(config Abstractions.IConfiguration) {
	router.configuration = config
	// server.path
	serverPath, hasPath := config.Get("yoyogo.application.server.path").(string)
	if hasPath {
		router.endPointRouterHandler.Component = serverPath
		log.Infof("server.path:%s", serverPath)
	}
	// mvc.template
	mvcTemplate, hasTemplate := config.Get("yoyogo.application.server.mvc.template").(string)
	if hasTemplate {
		if hasPath {
			mvcTemplate = serverPath + mvcTemplate
		}
		router.mvcControllerBuilder.GetMvcOptions().MapRoute(mvcTemplate)
		log.Infof("mvc.template:%s", mvcTemplate)

	}

}

func (router *DefaultRouterBuilder) GetConfiguration() Abstractions.IConfiguration {
	return router.configuration
}

func (router *DefaultRouterBuilder) UseMvc(used bool) {
	if used {
		router.mvcControllerBuilder = Mvc.NewControllerBuilder()
	} else {
		router.mvcControllerBuilder = nil
	}
}

func (router *DefaultRouterBuilder) IsMvc() bool {
	return router.mvcControllerBuilder != nil
}

func (router *DefaultRouterBuilder) GetMvcBuilder() *Mvc.ControllerBuilder {
	return router.mvcControllerBuilder
}

func (router *DefaultRouterBuilder) Search(ctx *Context.HttpContext, components []string, params url.Values) func(ctx *Context.HttpContext) {
	var handler func(ctx *Context.HttpContext) = nil
	pathComponents := strings.Split(ctx.Input.Request.URL.Path, "/")[1:]
	handler = router.endPointRouterHandler.Invoke(ctx, pathComponents)

	if handler == nil && router.IsMvc() {
		handler = router.mvcControllerBuilder.GetRouterHandler().Invoke(ctx, pathComponents)
	}

	return handler
}

func (router *DefaultRouterBuilder) MapSet(method, path string, handler func(ctx *Context.HttpContext)) {
	router.endPointRouterHandler.Insert(method, path, handler)
}

func (router *DefaultRouterBuilder) Map(method string, path string, handler func(ctx *Context.HttpContext)) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}
	router.MapSet(method, path, handler)
}

// GET register GET request handler
func (router *DefaultRouterBuilder) GET(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.GET, path, handler)
}

// HEAD register HEAD request handler
func (router *DefaultRouterBuilder) HEAD(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (router *DefaultRouterBuilder) OPTIONS(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.OPTIONS, path, handler)
}

// POST register POST request handler
func (router *DefaultRouterBuilder) POST(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.POST, path, handler)
}

// PUT register PUT request handler
func (router *DefaultRouterBuilder) PUT(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (router *DefaultRouterBuilder) PATCH(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (router *DefaultRouterBuilder) DELETE(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (router *DefaultRouterBuilder) CONNECT(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (router *DefaultRouterBuilder) TRACE(path string, handler func(ctx *Context.HttpContext)) {
	router.Map(Context.TRACE, path, handler)
}

// Any register any method handler
func (router *DefaultRouterBuilder) Any(path string, handler func(ctx *Context.HttpContext)) {
	for _, m := range Context.Methods {
		router.Map(m, path, handler)
	}
}

func (router *DefaultRouterBuilder) Group(name string, routerBuilderFunc func(router *RouterGroup)) {
	group := &RouterGroup{
		Name:          name,
		RouterHandler: router,
	}
	if routerBuilderFunc == nil {
		panic("routerBuilderFunc is nil")
	}

	routerBuilderFunc(group)
}
