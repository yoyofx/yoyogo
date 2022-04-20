package mvc

import (
	"errors"
	"github.com/yoyofx/yoyogo/web/binding"
	"github.com/yoyofx/yoyogo/web/context"
	"reflect"
)

var RequestMppingFuncs map[string]ActionParametersMappingFunc = map[string]ActionParametersMappingFunc{
	"HttpContext": httpContextMappingMapping,
	"Default":     requestBodyMappingMapping,
}

type ActionParametersMappingFunc func(paramName string, paramTypeName string, paramType reflect.Type, sourceContext *context.HttpContext) (reflect.Value, error)

func httpContextMappingMapping(paramName string, paramTypeName string, paramType reflect.Type, sourceContext *context.HttpContext) (reflect.Value, error) {
	var value reflect.Value
	var err error
	if paramTypeName == "HttpContext" {
		value = reflect.ValueOf(sourceContext)
	} else {
		err = errors.New("not HttpContext ")
	}
	return value, err
}

/**
 绑定
form-data/multipart-form , json , uri , header
tags: json , form , uri ,header
*/
func requestBodyMappingMapping(paramName string, paramTypeName string, paramType reflect.Type, sourceContext *context.HttpContext) (reflect.Value, error) {
	var value reflect.Value
	var err error
	reqBindingData := reflect.New(paramType).Interface()

	fmTags := map[string]bool{"header": false, "uri": false, "path": false}
	for fi := 0; fi < paramType.NumField(); fi++ {
		for key, _ := range fmTags {
			_, inTag := paramType.Field(fi).Tag.Lookup(key)
			if inTag {
				fmTags[key] = inTag
				continue
			}
		}
	}

	if paramType.NumField() > 0 {
		paramTypeName := paramType.Field(0).Name
		if paramTypeName == "RequestBody" || paramTypeName == "RequestGET" || paramTypeName == "RequestPOST" {
			err = sourceContext.Bind(reqBindingData)
			if fmTags["uri"] {
				_ = sourceContext.BindWith(reqBindingData, binding.Query)
			}
			if fmTags["header"] {
				_ = sourceContext.BindWith(reqBindingData, binding.Header)
			}
			if fmTags["path"] {
				_ = sourceContext.BindWithRouteData(reqBindingData)
			}
		} else {
			err = errors.New("Can't bind non mvc.RequestBody!")
		}

	} else {
		err = errors.New("Can't bind non mvc.RequestBody!")
	}
	value = reflect.ValueOf(reqBindingData)
	return value, err
}
