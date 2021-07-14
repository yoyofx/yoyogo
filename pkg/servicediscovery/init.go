package servicediscovery

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/strategy"
	"github.com/yoyofxteam/dependencyinjection"
	"strconv"
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

			// selector (LB) Strategy
			sdStrategy := config.GetString("yoyogo.cloud.discovery.strategy")
			// round-robin  , weight-time ,  random
			switch sdStrategy {
			case "random":
				serviceCollection.AddSingletonByImplements(strategy.NewRandom, new(servicediscovery.Strategy))
			case "weight-time":
				serviceCollection.AddSingletonByImplements(strategy.NewWeightedResponseTime(), new(servicediscovery.Strategy))
			default:
				serviceCollection.AddSingletonByImplements(strategy.NewRound, new(servicediscovery.Strategy))
			}

			sdRegEnableStr := config.GetString("yoyogo.cloud.discovery.register-enable")
			sdRegEnable := true
			if sdRegEnableStr != "" {
				sdRegEnable, _ = strconv.ParseBool(sdRegEnableStr)
			}
			sdConfig := servicediscovery.NewConfig(sdRegEnable)
			serviceCollection.AddSingleton(func() *servicediscovery.Config { return sdConfig })

		})
}
