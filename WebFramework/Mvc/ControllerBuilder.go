package Mvc

import (
	"github.com/maxzhang1985/yoyogo/Utils"
	"strings"
)

type ControllerBuilder struct {
	controllerDescriptors []ControllerDescriptor
	mvcRouterHandler      *RouterHandler
	Options               MvcOptions
}

// NewControllerBuilder new controller builder
func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{
		Options:          MvcOptions{},
		mvcRouterHandler: &RouterHandler{},
	}
}

// SetupOptions , setup mvc builder options
func (builder *ControllerBuilder) SetupOptions() {

}

// AddController add controller (ctor) to ioc.
func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	controllerName := Utils.GetCtorFuncName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	descriptor := NewControllerDescriptor(controllerName, controllerCtor)
	builder.controllerDescriptors = append(builder.controllerDescriptors, descriptor)
}

// GetControllerDescriptorList is get controller descriptor array
func (builder *ControllerBuilder) GetControllerDescriptorList() []ControllerDescriptor {
	return builder.controllerDescriptors
}

// GetRouterHandler is get mvc router handler.
func (builder *ControllerBuilder) GetRouterHandler() *RouterHandler {
	return builder.mvcRouterHandler
}
