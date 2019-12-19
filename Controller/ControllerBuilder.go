package Controller

import (
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/Utils"
	"strings"
)

type ControllerBuilder struct {
	services *DependencyInjection.ServiceCollection
}

func NewControllerBuilder(sc *DependencyInjection.ServiceCollection) *ControllerBuilder {
	return &ControllerBuilder{services: sc}
}

func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	controllerName := Utils.GetCtorFuncName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	builder.services.AddSingletonByNameAndImplements(controllerName, controllerCtor, new(IController))
}
