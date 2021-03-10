package contollers

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/web/mvc"
)

type SDController struct {
	mvc.ApiController
	discoveryCache  servicediscovery.Cache
	discoveryClient servicediscovery.IServiceDiscoveryClient
}

func NewSDController(sd servicediscovery.IServiceDiscoveryClient, cache servicediscovery.Cache) *SDController {
	return &SDController{discoveryClient: sd, discoveryCache: cache}
}

func (controller SDController) GetSD() mvc.ApiResult {
	serviceList := controller.discoveryClient.GetAllInstances("yoyogo_demo_dev")
	return controller.OK(serviceList)
}

func (controller SDController) GetServices() mvc.ApiResult {
	serviceList, _ := controller.discoveryClient.GetAllServices()
	return controller.OK(serviceList)
}

func (controller SDController) GetCacheServices() mvc.ApiResult {
	serviceList, _ := controller.discoveryCache.GetService("yoyogo_demo_dev")
	return controller.OK(serviceList)
}
