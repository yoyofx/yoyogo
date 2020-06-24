package Abstractions

import (
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type HostBuildContext struct {
	RequestDelegate        interface{}
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *Context.HostEnvironment
	ApplicationServicesDef *DependencyInjection.ServiceCollection
	ApplicationServices    DependencyInjection.IServiceProvider
}
