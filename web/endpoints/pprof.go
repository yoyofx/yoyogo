package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
	"net/http"
	"net/http/pprof"
)

func pprofHandler(h http.HandlerFunc) web.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *context.HttpContext) {
		c.Output.SetStatus(200)
		if c.Input.Path() == "/debug/pprof/" {
			c.Output.Header(context.HeaderContentType, context.MIMETextHTML)
		}
		handler.ServeHTTP(c.Output.GetWriter(), c.Input.GetReader())
	}
}

func UsePprof(routerBuilder router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded pprof endpoint.")

	routerBuilder.Group("/debug/pprof", func(prefixRouter *router.RouterGroup) {
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
