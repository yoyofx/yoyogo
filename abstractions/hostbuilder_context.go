package abstractions

import (
	"github.com/yoyofx/yoyogo/abstractions/hostenv"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

type HostBuilderContext struct {
	RequestDelegate        interface{}
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *HostEnvironment
	Configuration          IConfiguration
	HostConfiguration      *hostenv.HostConfig
	ApplicationServicesDef *dependencyinjection.ServiceCollection
	ApplicationServices    dependencyinjection.IServiceProvider
	HostServices           dependencyinjection.IServiceProvider
}
