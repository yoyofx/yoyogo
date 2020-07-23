package contollers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
)

type TestActionFilter struct {
}

func (f *TestActionFilter) OnActionExecuting(context Mvc.ActionFilterContext) bool {
	fmt.Println("TestActionFilter OnActionExecuting")
	return false
}

func (f *TestActionFilter) OnActionExecuted(context Mvc.ActionFilterContext) {
	fmt.Println("TestActionFilter OnActionExecuted")
}
