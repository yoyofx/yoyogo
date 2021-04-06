package servicediscovery

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/strategy"
)

func UseGeneralServiceDiscovery(serviceCollection *dependencyinjection.ServiceCollection) {
	// service discovery Client
	serviceCollection.AddSingletonByImplements(NewClient, new(servicediscovery.IServiceDiscoveryClient))
	// registration for Cache and options
	serviceCollection.AddSingletonByImplements(servicediscovery.NewCache, new(servicediscovery.Cache))

	// selector (LB) Strategy
	serviceCollection.AddSingletonByImplements(strategy.NewRound, new(servicediscovery.Strategy))
	//serviceCollection.AddSingletonByImplements(strategy.NewRandom, new(servicediscovery.Strategy))
	//serviceCollection.AddSingletonByImplements(strategy.NewWeightedResponseTime(), new(servicediscovery.Strategy))

	// selector for service discovery
	serviceCollection.AddSingletonByImplements(servicediscovery.NewSelector, new(servicediscovery.ISelector))
	// http client facotry

}
