package Mvc

type IActionFilter interface {
	OnActionExecuting(context ActionExecutorContext) bool
	OnActionExecuted(context ActionExecutorContext)
}
