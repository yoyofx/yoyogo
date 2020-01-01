package Router

import (
	"github.com/maxzhang1985/yoyogo/ActionResult"
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
	"strings"
)

type MvcRouterHandler struct {
}

func (handler *MvcRouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {

	if pathComponents == nil || len(pathComponents) < 2 {
		return nil
	}
	controllerName := strings.ToLower(pathComponents[0])
	if !strings.Contains(controllerName, "controller") {
		controllerName += "controller"
	}
	actionName := pathComponents[1]

	controller := Controller.ActivateController(ctx.RequiredServices, controllerName)

	executorContext := &Controller.ActionExecutorContext{
		ControllerName: controllerName,
		Controller:     controller,
		ActionName:     actionName,
		Context:        ctx,
		In:             &Controller.ActionExecutorInParam{},
	}
	actionMethodExecutor := Controller.NewActionMethodExecutor()
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
