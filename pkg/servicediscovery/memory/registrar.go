package memory

import (
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
)

type Registrar struct {
	services []servicediscovery.ServiceInstance
}

func NewServerDiscovery(serviceName string, serviceList []string) servicediscovery.IServiceDiscovery {
	var services []servicediscovery.ServiceInstance
	for _, service := range serviceList {
		instance := &servicediscovery.DefaultServiceInstance{
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

func (registrar *Registrar) GetName() string {
	return "memory"
}

func (registrar *Registrar) Register() error {
	panic("implement me")
}

func (registrar *Registrar) Update() error {
	panic("implement me")
}

func (registrar *Registrar) Unregister() error {
	panic("implement me")
}

func (registrar *Registrar) GetHealthyInstances(serviceName string) []servicediscovery.ServiceInstance {
	return registrar.services
}

func (registrar *Registrar) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	return registrar.services
}

func (registrar *Registrar) Destroy() error {
	panic("implement me")
}

func (registrar *Registrar) Watch(opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	return nil, nil
}
