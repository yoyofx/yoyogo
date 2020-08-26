package ServerDiscovery

import (
	"github.com/yoyofx/yoyogo/Abstractions/ServerDiscovery"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/Internal/ServerDiscovery/Nacos"
)

func UseNacos(serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddTransientByImplements(Nacos.NewServerDiscoveryWithDI, new(ServerDiscovery.IServerDiscovery))
}
