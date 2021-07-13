package console

const ProjectItem_startup_go = `
package {{.CurrentModelName}}

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofxteam/dependencyinjection"
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
