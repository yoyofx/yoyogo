package health

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddTransientByImplements(NewDiskHealthIndicator, new(Indicator))
		})
}
