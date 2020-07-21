package Mvc

type IActionFilter interface {
	OnActionExecuting(context ActionFilterContext) bool
	OnActionExecuted(context ActionFilterContext)
}
