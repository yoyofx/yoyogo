package Controller

import "github.com/maxzhang1985/yoyogo/DependencyInjection"

func ActivateController(serviceProvider DependencyInjection.IServiceProvider, controllerName string) IController {
	var controller IController
	err := serviceProvider.GetServiceByName(&controller, controllerName)
	if err != nil {
		panic("Controller not found! " + err.Error())
	}
	return controller
}
