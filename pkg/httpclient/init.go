package httpclient

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
		serviceCollection.AddSingleton(NewFactory)
	})
}