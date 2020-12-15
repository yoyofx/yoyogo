package router

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"net/url"
	"path"
	"strings"
)

type DefaultRouterBuilder struct {
	mvcControllerBuilder  *mvc.ControllerBuilder
	endPointRouterHandler *EndPointRouterHandler
	configuration         abstractions.IConfiguration
	log                   xlog.ILogger
}

func NewRouterBuilder() IRouterBuilder {

	endpoint := &EndPointRouterHandler{
		Component: "/",
		Methods:   make(map[string]func(ctx *context.HttpContext)),
	}

	defaultRouterHandler := &DefaultRouterBuilder{endPointRouterHandler: endpoint}
	defaultRouterHandler.log = xlog.GetXLogger("DefaultRouterBuilder")
	defaultRouterHandler.log.SetCustomLogFormat(nil)
	return defaultRouterHandler
}

func (router *DefaultRouterBuilder) SetConfiguration(config abstractions.IConfiguration) {
	router.configuration = config
	if config == nil {
		return
	}
	// server.path
	serverPath, hasPath := config.Get("yoyogo.application.server.path").(string)
	if hasPath {
		router.endPointRouterHandler.Component = serverPath
		router.log.Info("server.path:  %s", consolecolors.Green(serverPath))
	}
	// mvc.template
	mvcTemplate, hasTemplate := config.Get("yoyogo.application.server.mvc.template").(string)
	if !hasTemplate {
		mvcTemplate = mvc.DefaultMvcTemplate
	}
	if hasPath {
		mvcTemplate = path.Join(serverPath, mvcTemplate)
	}
	router.mvcControllerBuilder.GetMvcOptions().MapRoute(mvcTemplate)
	router.log.Info("mvc.template:  %s", consolecolors.Green(mvcTemplate))
	// mvc.serializer.json.camecase
	cameCase := config.GetBool("yoyogo.application.server.mvc.serializer.json.camecase")
	router.mvcControllerBuilder.GetMvcOptions().Serializer = &mvc.SerializerOption{JsonCameCase: cameCase}

}

func (router *DefaultRouterBuilder) GetConfiguration() abstractions.IConfiguration {
	return router.configuration
}

func (router *DefaultRouterBuilder) UseMvc(used bool) {
	if used {
		router.mvcControllerBuilder = mvc.NewControllerBuilder()
	} else {
		router.mvcControllerBuilder = nil
	}
}

func (router *DefaultRouterBuilder) IsMvc() bool {
	return router.mvcControllerBuilder != nil
}

func (router *DefaultRouterBuilder) GetMvcBuilder() *mvc.ControllerBuilder {
	return router.mvcControllerBuilder
}

func (router *DefaultRouterBuilder) Search(ctx *context.HttpContext, components []string, params url.Values) func(ctx *context.HttpContext) {
	var handler func(ctx *context.HttpContext) = nil
	pathComponents := strings.Split(ctx.Input.Request.URL.Path, "/")[1:]
	handler = router.endPointRouterHandler.Invoke(ctx, pathComponents)

	if handler == nil && router.IsMvc() {
		handler = router.mvcControllerBuilder.GetRouterHandler().Invoke(ctx, pathComponents)
	}

	return handler
}

func (router *DefaultRouterBuilder) MapSet(method, path string, handler func(ctx *context.HttpContext)) {
	router.endPointRouterHandler.Insert(method, path, handler)
}

func (router *DefaultRouterBuilder) Map(method string, path string, handler func(ctx *context.HttpContext)) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}
	router.MapSet(method, path, handler)
}

// GET register GET request handler
func (router *DefaultRouterBuilder) GET(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.GET, path, handler)
}

// HEAD register HEAD request handler
func (router *DefaultRouterBuilder) HEAD(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (router *DefaultRouterBuilder) OPTIONS(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.OPTIONS, path, handler)
}

// POST register POST request handler
func (router *DefaultRouterBuilder) POST(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.POST, path, handler)
}

// PUT register PUT request handler
func (router *DefaultRouterBuilder) PUT(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (router *DefaultRouterBuilder) PATCH(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (router *DefaultRouterBuilder) DELETE(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (router *DefaultRouterBuilder) CONNECT(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (router *DefaultRouterBuilder) TRACE(path string, handler func(ctx *context.HttpContext)) {
	router.Map(context.TRACE, path, handler)
}

// Any register any method handler
func (router *DefaultRouterBuilder) Any(path string, handler func(ctx *context.HttpContext)) {
	for _, m := range context.Methods {
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
