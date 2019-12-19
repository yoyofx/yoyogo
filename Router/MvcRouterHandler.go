package Router

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
	"reflect"
	"strings"
)

type MvcRouterHandler struct {
}

func (handler *MvcRouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {
	if pathComponents == nil || len(pathComponents) < 2 {
		return nil
	}
	controllerName := strings.ToLower(pathComponents[0])
	actionName := pathComponents[1]

	controller := Controller.ActivateController(ctx.RequiredServices, controllerName)

	executorContext := &Controller.ActionExecutorContext{
		ControllerName: controllerName,
		Controller:     controller,
		ActionName:     actionName,
		Context:        ctx,
	}
	actionMehtodExecutor := Controller.NewActionMethodExecutor()
	actionResult := actionMehtodExecutor.Execute(executorContext)

	return func(ctx *Context.HttpContext) {
		ctx.JSON(200, actionResult)
	}

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
				if paramType.NumField() > 0 && paramType.Field(0).Name == "RequestBody" {
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
