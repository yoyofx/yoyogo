package abstractions

import (
	"github.com/yoyofx/yoyogo/abstractions/hostenv"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/web/context"
)

type HostBuilderContext struct {
	RequestDelegate        interface{}
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *context.HostEnvironment
	Configuration          IConfiguration
	HostConfiguration      *hostenv.HostConfig
	ApplicationServicesDef *dependencyinjection.ServiceCollection
	ApplicationServices    dependencyinjection.IServiceProvider
	HostServices           dependencyinjection.IServiceProvider
}
