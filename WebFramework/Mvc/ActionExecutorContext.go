package Mvc

import (
	"github.com/maxzhang1985/yoyogo/Utils"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
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
