package Controller

import "github.com/maxzhang1985/yoyogo/DependencyInjection"

func ActivateController(serviceProvider DependencyInjection.IServiceProvider, controllerName string) (IController, error) {
	var controller IController
	err := serviceProvider.GetServiceByName(&controller, controllerName)
	return controller, err
}
