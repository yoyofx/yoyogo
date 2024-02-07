package contollers

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/pkg/httpclient"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
)

type SDController struct {
	mvc.ApiController `doc:"服务发现接口Controller"`
	discoveryCache    servicediscovery.Cache
	discoveryClient   servicediscovery.IServiceDiscoveryClient
	discoverySelector servicediscovery.ISelector
	httpFactory       httpclient.IDiscoveryClientFactory
}

func NewSDController(sd servicediscovery.IServiceDiscoveryClient, cache servicediscovery.Cache, selector servicediscovery.ISelector, factory httpclient.IDiscoveryClientFactory) *SDController {
	return &SDController{discoveryClient: sd, discoveryCache: cache, discoverySelector: selector, httpFactory: factory}
}

func (controller *SDController) GetSD() mvc.ApiResult {
	serviceList := controller.discoveryClient.GetAllInstances("yoyogo_demo_dev")
	return controller.OK(serviceList)
}

func (controller *SDController) GetServices() mvc.ApiResult {
	serviceList, _ := controller.discoveryClient.GetAllServices()
	return controller.OK(serviceList)
}

func (controller *SDController) GetCacheServices() mvc.ApiResult {
	serviceList, _ := controller.discoveryCache.GetService("yoyogo_demo_dev")
	return controller.OK(serviceList)
}

func (controller *SDController) GetOne() mvc.ApiResult {
	serviceList, _ := controller.discoverySelector.Select("yoyogo_demo_dev")
	return controller.OK(serviceList)
}

func (controller *SDController) HttpTest() mvc.ApiResult {
	client, err := controller.httpFactory.Create("http://[yoyogo_demo_dev]")
	if err != nil {
		panic(err)
	}
	request := httpclient.WithRequest().GET("/app/v1/user/getinfo").SetTimeout(2)
	response, err := client.Do(request)
	if err != nil {
		return controller.OK(context.H{"request_url": request.GetUrl(), "response_body": "", "err": err.Error()})
	} else {
		return controller.OK(context.H{"request_url": request.GetUrl(), "response_body": string(response.Body)})
	}
}
