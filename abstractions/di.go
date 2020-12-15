package abstractions

import "github.com/yoyofx/yoyogo/dependencyinjection"

var (
	configurationProcessors []func(config IConfiguration, serviceCollection *dependencyinjection.ServiceCollection)
)

func RegisterConfigurationProcessor(configure func(config IConfiguration, serviceCollection *dependencyinjection.ServiceCollection)) {
	configurationProcessors = append(configurationProcessors, configure)
}
