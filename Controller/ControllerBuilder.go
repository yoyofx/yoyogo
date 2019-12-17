package Controller

import (
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
)

type ControllerBuilder struct {
	services DependencyInjection.ServiceCollection
}

func NewControllerBuilder(sc DependencyInjection.ServiceCollection) *ControllerBuilder {
	return &ControllerBuilder{services: sc}
}

func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	//"usercontroller", contollers.NewUserController, new(Controller.IController)
	builder.services.AddSingletonByNameAndImplements("usercontroller", controllerCtor, new(IController))
}
