package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
	"net/http"
	"net/http/pprof"
)

func pprofHandler(h http.HandlerFunc) YoyoGo.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *Context.HttpContext) {
		c.Output.SetStatus(200)
		if c.Input.Path() == "/debug/pprof/" {
			c.Output.Header(Context.HeaderContentType, Context.MIMETextHTML)
		}
		handler.ServeHTTP(c.Output.GetWriter(), c.Input.GetReader())
	}
}

func UsePprof(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded pprof endpoint.")

	router.Group("/debug/pprof", func(prefixRouter *Router.RouterGroup) {
		prefixRouter.GET("/", pprofHandler(pprof.Index))
		prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
		prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
		prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
		prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	})

}
