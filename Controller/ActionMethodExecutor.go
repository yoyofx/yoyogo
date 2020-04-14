package Controller

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Utils"
	"net/http"
	"reflect"
)

type ActionMethodExecutor struct {
}

func NewActionMethodExecutor() ActionMethodExecutor {
	return ActionMethodExecutor{}
}

func (actionExecutor ActionMethodExecutor) Execute(ctx *ActionExecutorContext) interface{} {
	if ctx.Controller != nil {
		if ctx.In.MethodInovker == nil {
			ctx.In.MethodInovker = Utils.NewMethodCaller(ctx.Controller, ctx.ActionName)
			if ctx.In.MethodInovker == nil {
				ctx.Context.Response.WriteHeader(http.StatusNotFound)
				panic(ctx.ActionName + " action is not found! at " + ctx.ControllerName)
			}
			ctx.In.ActionParamTypes = ctx.In.MethodInovker.GetParamTypes()
		}

		values := getParamValues(ctx.In.ActionParamTypes, ctx.Context)
		returns := ctx.In.MethodInovker.Invoke(values...)
		if len(returns) > 0 {
			responseData := returns[0]
			return responseData
		}

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
				if paramType.NumField() > 0 && paramType.Field(0).Name == "RequestBody" {
					reqBindingData := reflect.New(paramType).Interface()
					_ = ctx.Bind(reqBindingData)
					values[index] = reqBindingData
				}
			}
		}
	}

	return values
}

func RequestParamTypeConvertFunc(index int, paramType reflect.Type, ctx *Context.HttpContext) {

}
