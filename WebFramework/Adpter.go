package YoyoGo

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net/http"
)

type HandlerFunc = func(ctx *Context.HttpContext)

func WarpHandlerFunc(h func(w http.ResponseWriter, r *http.Request)) HandlerFunc {
	return func(c *Context.HttpContext) {
		h(c.Output.GetWriter(), c.Input.GetReader())
	}
}

func WarpHttpHandlerFunc(h http.Handler) HandlerFunc {
	return func(c *Context.HttpContext) {
		h.ServeHTTP(c.Output.GetWriter(), c.Input.GetReader())
	}
}
