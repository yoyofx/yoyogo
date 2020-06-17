package Mvc

import (
	"github.com/maxzhang1985/yoyogo/Utils"
	"strings"
)

type ControllerBuilder struct {
	controllerDescriptors []ControllerDescriptor
	options               MvcOptions
}

func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{options: MvcOptions{}}
}

func (builder *ControllerBuilder) SetupOptions() {

}

func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	controllerName := Utils.GetCtorFuncName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	descriptor := NewControllerDescriptor(controllerName, controllerCtor)
	builder.controllerDescriptors = append(builder.controllerDescriptors, descriptor)
}

func (builder *ControllerBuilder) GetControllerDescriptorList() []ControllerDescriptor {
	return builder.controllerDescriptors
}
