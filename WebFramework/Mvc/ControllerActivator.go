package Mvc

import "github.com/yoyofx/yoyogo/DependencyInjection"

func ActivateController(serviceProvider DependencyInjection.IServiceProvider, controllerName string) (IController, error) {
	var controller IController
	err := serviceProvider.GetServiceByName(&controller, controllerName)
	return controller, err
}
