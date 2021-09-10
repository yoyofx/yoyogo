package mvc

import (
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofxteam/reflectx"
	"reflect"
)

type ActionMethodExecutor struct {
	logger xlog.ILogger
}

func NewActionMethodExecutor() ActionMethodExecutor {
	return ActionMethodExecutor{logger: xlog.GetXLogger("ActionMethodExecutor")}
}

func (actionExecutor ActionMethodExecutor) Execute(ctx *ActionExecutorContext) interface{} {
	if ctx.Controller != nil {
		methodInfo := ctx.ActionDescriptor.MethodInfo
		values := getParamValues(methodInfo.Parameters, ctx.Context)
		returns := methodInfo.InvokeWithValue(reflect.ValueOf(ctx.Controller), values...)
		if len(returns) > 0 {
			responseData := returns[0]
			return responseData
		}
	}

	return nil
}

func getParamValues(paramList []reflectx.MethodParameterInfo, ctx *context.HttpContext) []reflect.Value {
	if len(paramList) == 0 {
		return nil
	}
	values := make([]reflect.Value, len(paramList)-1)
	for index, param := range paramList {
		if index == 0 {
			continue
		}
		val, err := requestParamTypeConvertFunc(index, param, ctx)
		if err == nil {
			values[index-1] = val
		}
	}

	return values
}

func requestParamTypeConvertFunc(index int, parameter reflectx.MethodParameterInfo, ctx *context.HttpContext) (reflect.Value, error) {
	var value reflect.Value
	var err error = nil
	paramType := parameter.ParameterType
	if paramType.Kind() == reflect.Ptr {
		paramType = paramType.Elem()
	}
	if paramType.Kind() == reflect.Struct {
		switch paramType.Name() {
		case "HttpContext":
			value = reflect.ValueOf(ctx)
		default:
			if paramType.NumField() > 0 && paramType.Field(0).Name == "RequestBody" {
				reqBindingData := reflect.New(paramType).Interface()
				bindErr := ctx.Bind(reqBindingData)
				if bindErr != nil {
					fmt.Println(bindErr)
				}
				value = reflect.ValueOf(reqBindingData)
			}
		}
		return value, err
	}
	return value, errors.New("the type not support")
}
