package Mvc

import (
	"fmt"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net/http"
	"strings"
)

type RouterHandler struct {
	ControllerFilters     []ActionFilterChain
	ControllerDescriptors map[string]ControllerDescriptor
	Options               Options
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
		ctx.Response.WriteHeader(http.StatusNotFound)
		panic(controllerName + " controller is not found! " + err.Error())
		return nil
	}

	actionName := handler.Options.Template.ActionName
	controllerDescriptor := handler.ControllerDescriptors[controllerName]
	actionDescriptor, foundAction := controllerDescriptor.GetActionDescriptorByName(actionName)
	if !foundAction {
		actionName = fmt.Sprintf("%s%s", strings.ToLower(ctx.Request.Method), actionName)
		actionDescriptor, foundAction = controllerDescriptor.GetActionDescriptorByName(actionName)
	}
	if foundAction {
		actionName = actionDescriptor.ActionName
	} else {
		ctx.Response.WriteHeader(http.StatusNotFound)
		panic(actionName + " action is not found! ")
		return nil
	}

	actionMethodExecutor := NewActionMethodExecutor()
	executorContext := &ActionExecutorContext{
		ControllerName: controllerName,
		Controller:     controller,
		ActionName:     actionName,
		Context:        ctx,
	}

	//actionFilters := handler.MatchFilters(ctx)

	actionResult := actionMethodExecutor.Execute(executorContext)

	response := &RouterHandlerResponse{Result: actionResult}
	return response.Callback

}

func (handler RouterHandler) MatchFilters(ctx *Context.HttpContext) []IActionFilter {
	var filterList []IActionFilter
	for _, filterChain := range handler.ControllerFilters {
		actionFilter := filterChain.MatchFilter(ctx.Request.URL.Path)
		if actionFilter != nil {
			filterList = append(filterList, actionFilter)
		}
	}
	return filterList
}
