package abstractions

import "github.com/yoyofxteam/dependencyinjection"

var (
	configurationProcessors []func(config IConfiguration, serviceCollection *dependencyinjection.ServiceCollection)
)

func RegisterConfigurationProcessor(configure func(config IConfiguration, serviceCollection *dependencyinjection.ServiceCollection)) {
	configurationProcessors = append(configurationProcessors, configure)
}
