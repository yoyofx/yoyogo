package mvc

import (
	"github.com/yoyofx/yoyogo/web/context"
)

type IRouterAttributeBuilder interface {
	Match(ctx *context.HttpContext, pathComponents []string) (string, bool)
	Insert(method, path string, handler func(ctx *context.HttpContext))
}
