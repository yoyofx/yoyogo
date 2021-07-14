package mvc

import "github.com/yoyofxteam/dependencyinjection"

func ActivateController(serviceProvider dependencyinjection.IServiceProvider, controllerName string) (IController, error) {
	var controller IController
	err := serviceProvider.GetServiceByName(&controller, controllerName)
	return controller, err
}
