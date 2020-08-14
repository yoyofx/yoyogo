package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions/xlog"
	"github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
	"net/http"
	"net/http/pprof"
)

type pprofHandler struct {
	Path        string
	HandlerFunc YoyoGo.HandlerFunc
}

var debupApi = []pprofHandler{
	{"/debug/pprof/", WarpHandlerFunc(pprof.Index)},
	{"/debug/pprof/cmdline", WarpHandlerFunc(pprof.Cmdline)},
	{"/debug/pprof/profile", WarpHandlerFunc(pprof.Profile)},
	{"/debug/pprof/symbol", WarpHandlerFunc(pprof.Symbol)},
	{"/debug/pprof/trace", WarpHandlerFunc(pprof.Trace)},
}

func WarpHandlerFunc(h func(w http.ResponseWriter, r *http.Request)) YoyoGo.HandlerFunc {
	return func(c *Context.HttpContext) {
		if c.Input.Path() == "/debug/pprof/" {
			c.Output.Header(Context.HeaderContentType, Context.MIMETextHTML)
		}
		c.Output.SetStatus(200)
		h(c.Output.GetWriter(), c.Input.GetReader())
	}
}

func UsePprof(router Router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded pprof endpoint.")

	for _, item := range debupApi {
		router.GET(item.Path, item.HandlerFunc)
	}
}
