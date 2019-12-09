package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
)

type MvcRouterHandler struct {
}

func (handler *MvcRouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {
	var controllers Controller.IController

	err := ctx.RequiredServices.GetServiceByName(&controllers, pathComponents[0])
	if err != nil {
		panic(err)
	} else {

	}

	return nil
}
