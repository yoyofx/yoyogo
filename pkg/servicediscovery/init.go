package servicediscovery

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"time"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			ttl, _ := config.Get("yoyogo.cloud.discovery.cache.ttl").(int64)
			ttlDuration := servicediscovery.DefaultTTL // 30 * seconds
			if ttl > 0 {
				ttlDuration = time.Duration(ttl) * time.Second
			}
			serviceCollection.AddSingleton(func() *servicediscovery.CacheOptions {
				return &servicediscovery.CacheOptions{TTL: ttlDuration}
			})
		})
}
