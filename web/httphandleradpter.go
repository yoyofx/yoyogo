package web

import (
	"github.com/yoyofx/yoyogo/web/context"
	"net/http"
)

type HandlerFunc = func(ctx *context.HttpContext)

func WarpHandlerFunc(h func(w http.ResponseWriter, r *http.Request)) HandlerFunc {
	return func(c *context.HttpContext) {
		h(c.Output.GetWriter(), c.Input.GetReader())
	}
}

func WarpHttpHandlerFunc(h http.Handler) HandlerFunc {
	return func(c *context.HttpContext) {
		h.ServeHTTP(c.Output.GetWriter(), c.Input.GetReader())
	}
}
