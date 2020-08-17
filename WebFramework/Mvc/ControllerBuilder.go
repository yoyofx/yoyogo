package Mvc

import (
	"github.com/yoyofx/yoyogo/Abstractions/xlog"
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofxteam/reflectx"
	"strings"
)

// ControllerBuilder: controller builder
type ControllerBuilder struct {
	mvcRouterHandler *RouterHandler
}

// NewControllerBuilder new controller builder
func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{mvcRouterHandler: NewMvcRouterHandler()}
}

// add views to mvc
func (builder *ControllerBuilder) AddViews(option ViewOption) {
	xlog.GetXLogger("ControllerBuilder").Debug("add mvc views: %s", option.Pattern)
	builder.mvcRouterHandler.ViewEngine = &ActionResult.HTMLDebug{Files: option.Files,
		Glob:    option.Pattern,
		Delims:  ActionResult.Delims{Left: "{[{", Right: "}]}"},
		FuncMap: option.FuncMap,
	}
}

// add filter to mvc
func (builder *ControllerBuilder) AddFilter(pattern string, actionFilter IActionFilter) {
	xlog.GetXLogger("ControllerBuilder").Debug("add mvc filter: %s", pattern)
	chain := NewActionFilterChain(pattern, actionFilter)
	builder.mvcRouterHandler.ControllerFilters = append(builder.mvcRouterHandler.ControllerFilters, chain)
}

// SetupOptions , setup mvc builder options
func (builder *ControllerBuilder) SetupOptions(configOption func(options Options)) {
	configOption(builder.mvcRouterHandler.Options)
}

// AddController add controller (ctor) to ioc.
func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	logger := xlog.GetXLogger("ControllerBuilder")

	controllerName, controllerType := reflectx.GetCtorFuncOutTypeName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	// Create Controller and Action descriptors
	descriptor := NewControllerDescriptor(controllerName, controllerType, controllerCtor)
	builder.mvcRouterHandler.ControllerDescriptors[controllerName] = descriptor

	logger.Debug("add mvc controller: %s", controllerName)
	for _, desc := range descriptor.GetActionDescriptors() {
		logger.Debug("add mvc controller action: %s", desc.ActionName)
	}
}

// GetControllerDescriptorList is get controller descriptor array
func (builder *ControllerBuilder) GetControllerDescriptorList() []ControllerDescriptor {
	values := make([]ControllerDescriptor, 0, len(builder.mvcRouterHandler.ControllerDescriptors))
	for _, value := range builder.mvcRouterHandler.ControllerDescriptors {
		values = append(values, value)
	}
	return values
}

// GetControllerDescriptorByName get controller descriptor by controller name
func (builder *ControllerBuilder) GetControllerDescriptorByName(name string) ControllerDescriptor {
	return builder.mvcRouterHandler.ControllerDescriptors[name]
}

// GetMvcOptions get mvc options
func (builder *ControllerBuilder) GetMvcOptions() Options {
	return builder.mvcRouterHandler.Options
}

// GetRouterHandler is get mvc router handler.
func (builder *ControllerBuilder) GetRouterHandler() *RouterHandler {
	return builder.mvcRouterHandler
}
