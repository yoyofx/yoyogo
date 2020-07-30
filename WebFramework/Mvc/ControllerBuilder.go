package Mvc

import (
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

// add filter to mvc
func (builder *ControllerBuilder) AddFilter(pattern string, actionFilter IActionFilter) {
	chain := NewActionFilterChain(pattern, actionFilter)
	builder.mvcRouterHandler.ControllerFilters = append(builder.mvcRouterHandler.ControllerFilters, chain)
}

// SetupOptions , setup mvc builder options
func (builder *ControllerBuilder) SetupOptions(configOption func(options Options)) {
	configOption(builder.mvcRouterHandler.Options)
}

// AddController add controller (ctor) to ioc.
func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	controllerName, controllerType := reflectx.GetCtorFuncOutTypeName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	// Create Controller and Action descriptors
	descriptor := NewControllerDescriptor(controllerName, controllerType, controllerCtor)
	builder.mvcRouterHandler.ControllerDescriptors[controllerName] = descriptor
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
