package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
)

type HostBuildContext struct {
	RequestDelegate        IRequestDelegate
	ApplicationCycle       *ApplicationLife
	hostingEnvironment     *Context.HostEnvironment
	applicationServicesDef *DependencyInjection.ServiceCollection
	applicationServices    DependencyInjection.IServiceProvider
}
