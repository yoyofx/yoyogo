package main

import (
	"github.com/yoyofx/yoyogo/console"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func main() {
	//configuration := abstractions.NewConfigurationBuilder().
	//	AddEnvironment().
	//	AddYamlFile("config").Build()

	hosting := console.NewHostBuilder().
		//UseConfiguration(configuration).
		//Configure(func(app *console.ApplicationBuilder) {
		//}).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {

		}).Build()

	hosting.Run()

}
