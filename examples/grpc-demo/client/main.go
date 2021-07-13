package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofx/yoyogo/console"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/nacos"
	"github.com/yoyofxteam/dependencyinjection"
)

func main() {
	configuration := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	console.NewHostBuilder().
		UseConfiguration(configuration).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			hosting.AddHostService(collection, NewClientService)
			collection.AddSingleton(NewHelloworldApi)
			//register sd for nacos
			nacos.UseServiceDiscovery(collection)
		}).
		Build().Run()
}
