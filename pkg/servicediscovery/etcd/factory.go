package etcd

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	sd "github.com/yoyofx/yoyogo/pkg/servicediscovery"
)

func UseServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	sd.UseGeneralServiceDiscovery(serviceCollection)
	serviceCollection.AddSingletonByImplements(NewServerDiscoveryWithDI, new(servicediscovery.IServiceDiscovery))
	//serviceCollection.AddSingletonByImplements(sd.NewClient, new(servicediscovery.IServiceDiscoveryClient))
}
