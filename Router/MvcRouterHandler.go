package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
	"github.com/maxzhang1985/yoyogo/Utils"
	"reflect"
)

type MvcRouterHandler struct {
}

func (handler *MvcRouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {
	controllerName := pathComponents[0]
	actionName := pathComponents[1]

	var controllers Controller.IController
	err := ctx.RequiredServices.GetServiceByName(&controllers, controllerName)
	if err != nil {
		panic("Controller not found! " + err.Error())
	} else {
		caller := Utils.NewMethodCaller(controllers, actionName)
		if caller != nil {
			_ = getParamValues(caller.GetParamTypes(), ctx)

		}
		//
	}

	return nil
}

func getParamValues(paramTypes []reflect.Type, ctx *Context.HttpContext) []reflect.Value {
	//_= types[0].Field(0).
	return nil
}
