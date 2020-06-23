package Abstract

import (
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
)

type HostBuildContext struct {
	RequestDelegate        IRequestDelegate
	ApplicationCycle       *ApplicationLife
	HostingEnvironment     *Context.HostEnvironment
	ApplicationServicesDef *DependencyInjection.ServiceCollection
	ApplicationServices    DependencyInjection.IServiceProvider
}
