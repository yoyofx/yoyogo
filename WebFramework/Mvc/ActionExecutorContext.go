package Mvc

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type ActionExecutorContext struct {
	ControllerName string
	ActionName     string
	Controller     IController
	Context        *Context.HttpContext
}
