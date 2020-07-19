package Mvc

import (
	"errors"
	"github.com/yoyofx/yoyogo/Utils/Reflect"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
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
		methodInfo, methodFounded := Reflect.GetObjectMethodInfoByName(ctx.Controller, ctx.ActionName)
		if methodFounded {
			values := getParamValues(methodInfo.Parameters, ctx.Context)
			returns := methodInfo.InvokeWithValue(values...)
			if len(returns) > 0 {
				responseData := returns[0]
				return responseData
			}
		} else {
			ctx.Context.Response.WriteHeader(http.StatusNotFound)
			panic(ctx.ActionName + " action is not found! at " + ctx.ControllerName)
		}
	}

	return nil
}

func getParamValues(paramList []Reflect.ParameterInfo, ctx *Context.HttpContext) []reflect.Value {
	if len(paramList) == 0 {
		return nil
	}
	values := make([]reflect.Value, len(paramList))
	for index, param := range paramList {
		val, err := requestParamTypeConvertFunc(index, param, ctx)
		if err == nil {
			values[index] = val
		}
	}

	return values
}

func getParamValues1(paramTypes []reflect.Type, ctx *Context.HttpContext) []interface{} {
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

func requestParamTypeConvertFunc(index int, parameter Reflect.ParameterInfo, ctx *Context.HttpContext) (reflect.Value, error) {
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
				_ = ctx.Bind(reqBindingData)
				value = reflect.ValueOf(reqBindingData)
			}
		}
		return value, err
	}
	return value, errors.New("the type not support")
}
