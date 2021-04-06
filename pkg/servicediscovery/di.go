package servicediscovery

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func UseGeneralServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	serviceCollection.AddSingletonByImplements(NewClient, new(servicediscovery.IServiceDiscoveryClient))
	// registration for Cache and options
	serviceCollection.AddSingletonByImplements(servicediscovery.NewCache, new(servicediscovery.Cache))
	//serviceCollection.AddSingleton(func() interface{} {
	//	return servicediscovery.CacheOptions{TTL: servicediscovery.DefaultTTL}
	//})
	//----------------------------------------------------------------------------
}
