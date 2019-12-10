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
			values := getParamValues(caller.GetParamTypes(), ctx)
			returns := caller.Invoke(values...)
			if len(returns) > 0 {
				responseData := returns[0]
				return func(ctx *Context.HttpContext) {
					ctx.JSON(200, responseData)
				}
			}

		}
		//
	}

	return nil
}

func getParamValues(paramTypes []reflect.Type, ctx *Context.HttpContext) []interface{} {
	if len(paramTypes) == 0 {
		return nil
	}
	values := make([]interface{}, len(paramTypes))
	for index, paramType := range paramTypes {
		if paramType.Kind() == reflect.Ptr {
			paramType = paramType.Elem()
		}
		if paramType.Kind() == reflect.Struct {
			switch paramType.Name() {
			case "HttpContext":
				values[index] = ctx
			default:
				if paramType.NumField() > 0 && paramType.Field(0).Name == "RequestParam" {
					reqBindingData := reflect.New(paramType).Interface()
					_ = ctx.Bind(reqBindingData)
					values[index] = reqBindingData
				}
			}

		}

	}

	//type1 := paramTypes[1].Elem()
	//d := reflect.New(type1).Interface()
	//_ = ctx.Bind(d)
	return values
}

func RequestParamTypeConvertFunc(index int, paramType reflect.Type, ctx *Context.HttpContext) {

}
