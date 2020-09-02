package Nacos

import (
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"github.com/yoyofx/yoyogo/DependencyInjection"
)

func UseServiceDiscovery(serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddTransientByImplements(NewServerDiscoveryWithDI, new(ServiceDiscovery.IServiceDiscovery))
}
