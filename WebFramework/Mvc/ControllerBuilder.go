package Mvc

import (
	"github.com/yoyofx/yoyogo/Utils/Reflect"
	"strings"
)

type ControllerBuilder struct {
	controllerDescriptors []ControllerDescriptor
	mvcRouterHandler      *RouterHandler
	Options               Options
}

// NewControllerBuilder new controller builder
func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{
		Options:          NewMvcOptions(),
		mvcRouterHandler: NewMvcRouterHandler(),
	}
}

// SetupOptions , setup mvc builder options
func (builder *ControllerBuilder) SetupOptions(configOption func(options Options)) {
	configOption(builder.Options)
}

// AddController add controller (ctor) to ioc.
func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	controllerName, controllerType := Reflect.GetCtorFuncOutTypeName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	descriptor := NewControllerDescriptor(controllerName, controllerCtor)
	instance := Reflect.CreateInstance(controllerType)
	ms := Reflect.GetObjectMehtodInfoList(instance)
	_ = ms
	builder.controllerDescriptors = append(builder.controllerDescriptors, descriptor)
}

// GetControllerDescriptorList is get controller descriptor array
func (builder *ControllerBuilder) GetControllerDescriptorList() []ControllerDescriptor {
	return builder.controllerDescriptors
}

// GetMvcOptions get mvc options
func (builder *ControllerBuilder) GetMvcOptions() Options {
	return builder.Options
}

// GetRouterHandler is get mvc router handler.
func (builder *ControllerBuilder) GetRouterHandler() *RouterHandler {
	return builder.mvcRouterHandler
}
