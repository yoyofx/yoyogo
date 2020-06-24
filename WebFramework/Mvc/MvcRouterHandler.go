package Mvc

import (
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net/http"
	"strings"
)

type RouterHandler struct {
}

func (handler *RouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {

	if pathComponents == nil || len(pathComponents) < 2 {
		return nil
	}
	controllerName := strings.ToLower(pathComponents[0])
	if !strings.Contains(controllerName, "controller") {
		controllerName += "controller"
	}
	actionName := pathComponents[1]

	controller, err := ActivateController(ctx.RequiredServices, controllerName)
	if err != nil {
		ctx.Response.WriteHeader(http.StatusNotFound)
		panic(controllerName + " controller is not found! " + err.Error())
	}

	actionMethodExecutor := NewActionMethodExecutor()
	executorContext := &ActionExecutorContext{
		ControllerName: controllerName,
		Controller:     controller,
		ActionName:     actionName,
		Context:        ctx,
		In:             nil,
	}
	executorContext.In = &ActionExecutorInParam{}

	actionResult := actionMethodExecutor.Execute(executorContext)

	ctx.SetItem("actionResult", actionResult)

	return func(ctx *Context.HttpContext) {
		result := ctx.GetItem("actionResult")

		if actionResult, ok := result.(ActionResult.IActionResult); ok {
			ctx.Render(200, actionResult)
		} else {
			contentType := ctx.Request.Header.Get(Context.HeaderContentType)
			switch {
			case strings.HasPrefix(contentType, Context.MIMEApplicationXML):
				ctx.XML(200, result)
			case strings.HasPrefix(contentType, Context.MIMEApplicationYAML):
				ctx.YAML(200, result)
			case strings.HasPrefix(contentType, Context.MIMEApplicationJSON):
				fallthrough
			default:
				ctx.JSON(200, result)

			}

		}

	}

}

//func findControllerAction() {
//	t := reflect.ValueOf(method.Object)
//	method.methodInfo = t.MethodByName(method.MethodName)
//	if !method.methodInfo.IsValid() {
//		return false
//	}
//}
