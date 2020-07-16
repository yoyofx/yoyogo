package Mvc

import (
	"github.com/yoyofx/yoyogo/Utils/Reflect"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"reflect"
)

type ActionExecutorInParam struct {
	ActionParamTypes []reflect.Type
	MethodInovker    *Reflect.MethodCaller
}

type ActionExecutorContext struct {
	ControllerName string
	ActionName     string
	Controller     IController
	Context        *Context.HttpContext
}
