package Mvc

import (
	"fmt"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/View"
	"net/http"
	"strings"
)

type RouterHandler struct {
	ControllerFilters     []ActionFilterChain
	ControllerDescriptors map[string]ControllerDescriptor
	Options               *Options
	ViewEngine            View.IViewEngine
}

func NewMvcRouterHandler() *RouterHandler {
	return &RouterHandler{
		Options:               NewMvcOptions(),
		ControllerDescriptors: make(map[string]ControllerDescriptor),
	}
}

func (handler *RouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {
	if !handler.Options.Template.Match(pathComponents) {
		return nil
	}

	controllerName := handler.Options.Template.ControllerName
	controller, err := ActivateController(ctx.RequiredServices, controllerName)
	if err != nil {
		ctx.Output.SetStatus(http.StatusNotFound)
		panic(controllerName + " controller is not found! " + err.Error())
		return nil
	}

	actionName := handler.Options.Template.ActionName
	controllerDescriptor := handler.ControllerDescriptors[controllerName]
	actionDescriptor, foundAction := controllerDescriptor.GetActionDescriptorByName(actionName)

	if foundAction && actionDescriptor.ActionMethod != "any" && strings.ToLower(ctx.Input.Method()) != actionDescriptor.ActionMethod {
		ctx.Output.SetStatus(http.StatusMethodNotAllowed)
		panic(fmt.Sprintf("Status method not allowed ! Request method is %s ,that define with %s .", ctx.Input.Method(),
			strings.ToUpper(actionDescriptor.ActionMethod)))
		return nil
	}

	if foundAction {
		actionName = actionDescriptor.ActionName
	} else {
		ctx.Output.SetStatus(http.StatusMethodNotAllowed)
		panic(actionName + " action is not found! ")
		return nil
	}

	if handler.ViewEngine != nil {
		controller.SetViewEngine(handler.ViewEngine)
	}

	actionMethodExecutor := NewActionMethodExecutor()
	executorContext := &ActionExecutorContext{
		ControllerName:   controllerName,
		Controller:       controller,
		ActionName:       actionName,
		ActionDescriptor: actionDescriptor,
		Context:          ctx,
	}

	actionFilterContext := ActionFilterContext{*executorContext, nil}
	filterPassed := true
	actionFilters := handler.MatchFilters(ctx)
	if len(actionFilters) > 0 {
	FilterLoop:
		for _, filter := range actionFilters {
			filterPassed = filter.OnActionExecuting(actionFilterContext)
			if !filterPassed {
				break FilterLoop
			}
		}
	}

	var actionResult interface{}
	if filterPassed {
		//Execute Action
		actionResult = actionMethodExecutor.Execute(executorContext)
		actionFilterContext.Result = actionResult
		for _, filter := range actionFilters {
			filter.OnActionExecuted(actionFilterContext)
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, Context.H{"Message": "Unauthorized"})
	}

	response := &RouterHandlerResponse{Result: actionResult}
	return response.Callback

}

func (handler RouterHandler) MatchFilters(ctx *Context.HttpContext) []IActionFilter {
	var filterList []IActionFilter
	for _, filterChain := range handler.ControllerFilters {
		actionFilter := filterChain.MatchFilter(ctx.Input.Request.URL.Path)
		if actionFilter != nil {
			filterList = append(filterList, actionFilter)
		}
	}
	return filterList
}
