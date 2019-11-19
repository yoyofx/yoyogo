package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"net/url"
)

type IRouterHandler interface {
	Map(method, path string, handler func(ctx *Context.HttpContext))
	Search(ctx *Context.HttpContext, components []string, params url.Values) func(ctx *Context.HttpContext)
}
