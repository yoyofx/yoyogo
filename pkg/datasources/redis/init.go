package redis

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/health"
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
	"github.com/yoyofxteam/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("redis-master", NewRedis, new(abstractions.IDataSource))
			serviceCollection.AddTransientByImplements(NewRedisClient, new(redis.IClient))
			serviceCollection.AddTransientByImplements(NewRedisHealthIndicator, new(health.Indicator))
		})
}
