package mvc

import (
	"github.com/yoyofx/yoyogo/web/context"
)

type ActionExecutorContext struct {
	ControllerName   string
	ActionName       string
	Controller       IController
	ActionDescriptor ActionDescriptor
	Context          *context.HttpContext
}
