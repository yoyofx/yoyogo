package YoyoGo

import "github.com/maxzhang1985/yoyogo/DependencyInjection"

type HostBuildContext struct {
	RequestDelegate     IRequestDelegate
	ApplicationCycle    *ApplicationLife
	hostingEnvironment  *HostEnv
	applicationServices DependencyInjection.IServiceProvider
}
