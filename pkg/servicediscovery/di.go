package servicediscovery

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	grpconn "github.com/yoyofx/yoyogo/grpc/conn"
	"github.com/yoyofx/yoyogo/pkg/httpclient"
	"github.com/yoyofxteam/dependencyinjection"
)

func UseGeneralServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	// service discovery Client
	serviceCollection.AddSingletonByImplements(NewClient, new(servicediscovery.IServiceDiscoveryClient))
	// registration for Cache and options
	serviceCollection.AddSingletonByImplements(servicediscovery.NewCache, new(servicediscovery.Cache))
	// selector for service discovery
	serviceCollection.AddSingletonByImplements(servicediscovery.NewSelector, new(servicediscovery.ISelector))
	// http client factory
	serviceCollection.AddSingleton(httpclient.NewDiscoveryClientFactory)
	// grpc client factory
	serviceCollection.AddSingleton(grpconn.NewFactory)

}
