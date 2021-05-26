package scheduler

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/console"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func NewXxlJobBuilder(config abstractions.IConfiguration) *abstractions.HostBuilder {
	return console.NewHostBuilder().
		UseConfiguration(config).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			UseXxlJob(collection)
		})
}
