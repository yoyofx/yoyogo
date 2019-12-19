package Controller

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Utils"
	"reflect"
)

type ActionExecutorInParam struct {
	ActionParamTypes []reflect.Type
	MethodInovker    *Utils.MethodCaller
}

type ActionExecutorContext struct {
	ControllerName string
	ActionName     string

	Controller IController
	Context    *Context.HttpContext
	In         *ActionExecutorInParam
}
