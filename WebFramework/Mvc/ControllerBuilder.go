package Mvc

import (
	"github.com/yoyofx/yoyogo/Utils/Reflect"
	"strings"
)

type ControllerBuilder struct {
	controllerDescriptors map[string]ControllerDescriptor
	//controllerDescriptors []ControllerDescriptor
	mvcRouterHandler *RouterHandler
	Options          Options
}

// NewControllerBuilder new controller builder
func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{
		Options:               NewMvcOptions(),
		mvcRouterHandler:      NewMvcRouterHandler(),
		controllerDescriptors: make(map[string]ControllerDescriptor),
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
	// Create Controller and Action descriptors
	descriptor := NewControllerDescriptor(controllerName, controllerType, controllerCtor)
	builder.controllerDescriptors[controllerName] = descriptor
}

// GetControllerDescriptorList is get controller descriptor array
func (builder *ControllerBuilder) GetControllerDescriptorList() []ControllerDescriptor {
	values := make([]ControllerDescriptor, 0, len(builder.controllerDescriptors))
	for _, value := range builder.controllerDescriptors {
		values = append(values, value)
	}
	return values
}

// GetControllerDescriptorByName get controller descriptor by controller name
func (builder *ControllerBuilder) GetControllerDescriptorByName(name string) ControllerDescriptor {
	return builder.controllerDescriptors[name]
}

// GetMvcOptions get mvc options
func (builder *ControllerBuilder) GetMvcOptions() Options {
	return builder.Options
}

// GetRouterHandler is get mvc router handler.
func (builder *ControllerBuilder) GetRouterHandler() *RouterHandler {
	return builder.mvcRouterHandler
}
