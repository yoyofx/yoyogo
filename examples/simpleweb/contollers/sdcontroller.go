package contollers

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/web/mvc"
)

type SDController struct {
	mvc.ApiController

	discoveryClient servicediscovery.IServiceDiscovery
}

func NewSDController(sd servicediscovery.IServiceDiscovery) *SDController {
	return &SDController{discoveryClient: sd}
}

func (controller SDController) GetSD() mvc.ApiResult {
	serviceList := controller.discoveryClient.GetAllInstances("yoyogo_demo_dev")
	return controller.OK(serviceList)
}

func (controller SDController) GetServices() mvc.ApiResult {
	serviceList, _ := controller.discoveryClient.GetAllServices()
	for _, service := range serviceList {
		service.Nodes = controller.discoveryClient.GetAllInstances(service.Name)
	}
	return controller.OK(serviceList)
}
