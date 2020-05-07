package Controller

type IActionFilter interface {
	OnActionExecuting(context ActionExecutorContext) bool
	OnActionExecuted(context ActionExecutorContext)
}
