package Abstractions

import (
	"github.com/yoyofx/yoyogo/Abstractions/configs"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type HostBuildContext struct {
	RequestDelegate        interface{}
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *Context.HostEnvironment
	Configuration          IConfiguration
	HostConfiguration      *configs.HostConfig
	ApplicationServicesDef *DependencyInjection.ServiceCollection
	ApplicationServices    DependencyInjection.IServiceProvider
	HostServices           DependencyInjection.IServiceProvider
}
