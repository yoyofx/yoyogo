package contollers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/web/mvc"
)

type TestActionFilter struct {
}

func (f *TestActionFilter) OnActionExecuting(context mvc.ActionFilterContext) bool {
	fmt.Println("TestActionFilter OnActionExecuting")
	return false
}

func (f *TestActionFilter) OnActionExecuted(context mvc.ActionFilterContext) {
	fmt.Println("TestActionFilter OnActionExecuted")
}
