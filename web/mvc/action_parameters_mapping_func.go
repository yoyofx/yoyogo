package mvc

import (
	"errors"
	"github.com/yoyofx/yoyogo/web/context"
	"reflect"
)

type ActionParametersMappingFunc interface {
	Mapping(paramName string, paramTypeName string, paramType reflect.Type, sourceContext *context.HttpContext) (reflect.Value, error)
}

type httpContextMapping struct{}

func (mapping httpContextMapping) Mapping(paramName string, paramTypeName string, paramType reflect.Type, sourceContext *context.HttpContext) (reflect.Value, error) {
	var value reflect.Value
	var err error
	if paramTypeName == "HttpContext" {
		value = reflect.ValueOf(sourceContext)
	} else {
		err = errors.New("not HttpContext ")
	}
	return value, err
}

type requestBodyMapping struct{}

func (mapping requestBodyMapping) Mapping(paramName string, paramTypeName string, paramType reflect.Type, sourceContext *context.HttpContext) (reflect.Value, error) {
	var value reflect.Value
	var err error
	reqBindingData := reflect.New(paramType).Interface()
	if paramType.NumField() > 0 && paramType.Field(0).Name == "RequestBody" {
		err = sourceContext.Bind(reqBindingData)
	} else {
		err = errors.New("Can't bind non mvc.RequestBody!")
	}
	value = reflect.ValueOf(reqBindingData)
	return value, err
}
