package servicediscovery

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/pkg/httpclient"
)

func UseGeneralServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	// service discovery Client
	serviceCollection.AddSingletonByImplements(NewClient, new(servicediscovery.IServiceDiscoveryClient))
	// registration for Cache and options
	serviceCollection.AddSingletonByImplements(servicediscovery.NewCache, new(servicediscovery.Cache))
	// selector for service discovery
	serviceCollection.AddSingletonByImplements(servicediscovery.NewSelector, new(servicediscovery.ISelector))
	// http client facotry
	serviceCollection.AddSingleton(httpclient.NewDiscoveryClientFactory)
}
