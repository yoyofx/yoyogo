package Controller

import (
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/Utils"
)

type ControllerBuilder struct {
	services *DependencyInjection.ServiceCollection
}

func NewControllerBuilder(sc *DependencyInjection.ServiceCollection) *ControllerBuilder {
	return &ControllerBuilder{services: sc}
}

func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	controllerName := Utils.GetCtorFuncName(controllerCtor)
	controllerName = Utils.LowercaseFirst(controllerName)
	builder.services.AddSingletonByNameAndImplements(controllerName, controllerCtor, new(IController))
}
