package redis

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("redis1", NewRedis, new(abstractions.IDataSource))
		})
}
