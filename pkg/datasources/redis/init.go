package redis

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("redis-master", NewRedis, new(abstractions.IDataSource))
			serviceCollection.AddTransientByImplements(NewRedisClient, new(redis.IClient))
		})
}
