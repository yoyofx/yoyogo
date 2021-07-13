package consul

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	sd "github.com/yoyofx/yoyogo/pkg/servicediscovery"
	"github.com/yoyofxteam/dependencyinjection"
)

func UseServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	sd.UseGeneralServiceDiscovery(serviceCollection)
	serviceCollection.AddSingletonByImplements(NewServerDiscoveryWithDI, new(servicediscovery.IServiceDiscovery))
	//serviceCollection.AddSingletonByImplements(sd.NewClient, new(servicediscovery.IServiceDiscoveryClient))

}
