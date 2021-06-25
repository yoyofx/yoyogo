package console

const ProjectItem_startup_go = `
package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

type AppStartup struct {
}

func Startup() abstractions.IStartup {
	return &AppStartup{}
}

func (s *AppStartup) ConfigureServices(collection *dependencyinjection.ServiceCollection) {
	hosting.AddHostService(collection, NewService)
}
`
