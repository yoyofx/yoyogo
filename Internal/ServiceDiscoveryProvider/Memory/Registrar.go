package Memory

import (
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
)

type Registrar struct {
	services []ServiceDiscovery.ServiceInstance
}

func NewServerDiscovery(serviceName string, serviceList []string) ServiceDiscovery.IServiceDiscovery {
	var services []ServiceDiscovery.ServiceInstance
	for _, service := range serviceList {
		instance := &ServiceDiscovery.DefaultServiceInstance{
			Id:          uuid.New().String(),
			ServiceName: serviceName,
			Host:        service,
			Port:        8080,
			Enable:      true,
			Weight:      0,
			Healthy:     true,
		}
		services = append(services, instance)
	}
	return &Registrar{services: services}
}

func (r Registrar) GetName() string {
	return "memory"
}

func (r Registrar) Register() error {
	panic("implement me")
}

func (r Registrar) Update() error {
	panic("implement me")
}

func (r Registrar) Unregister() error {
	panic("implement me")
}

func (r Registrar) GetHealthyInstances(serviceName string) []ServiceDiscovery.ServiceInstance {
	return r.services
}

func (r Registrar) GetAllInstances(serviceName string) []ServiceDiscovery.ServiceInstance {
	return r.services
}

func (r Registrar) Destroy() error {
	panic("implement me")
}
