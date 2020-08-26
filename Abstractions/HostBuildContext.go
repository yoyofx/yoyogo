package Abstractions

import (
	"github.com/yoyofx/yoyogo/Abstractions/Configs"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type HostBuildContext struct {
	RequestDelegate        interface{}
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *Context.HostEnvironment
	Configuration          IConfiguration
	HostConfiguration      *Configs.HostConfig
	ApplicationServicesDef *DependencyInjection.ServiceCollection
	ApplicationServices    DependencyInjection.IServiceProvider
	HostServices           DependencyInjection.IServiceProvider
}
