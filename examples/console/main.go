package main

import (
	"fmt"
	"github.com/yoyofx/yoyogo/console"
)

func main() {
	//configuration := abstractions.NewConfigurationBuilder().
	//	AddEnvironment().
	//	AddYamlFile("config").Build()

	//hosting := console.NewHostBuilder().
	//UseConfiguration(configuration).
	//Configure(func(app *console.ApplicationBuilder) {
	//}).
	//	ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
	//		hosting.AddHostService(collection, NewService)
	//	}).Build()
	//
	//hosting.Run()

	console.NewHostBuilder().
		UseStartup(Startup).
		//ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
		//	hosting.AddHostService(collection, NewService)
		//}).
		Build().Run()
}

type Service1 struct {
}

func NewService() *Service1 {
	return &Service1{}
}

func (s *Service1) Run() error {
	fmt.Println("host service Running")
	return nil
}

func (s *Service1) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}
