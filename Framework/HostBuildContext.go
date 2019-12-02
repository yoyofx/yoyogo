package YoyoGo

import "github.com/maxzhang1985/yoyogo/DependencyInjection"

type HostBuildContext struct {
	hostingEnvironment  *HostEnv
	applicationServices DependencyInjection.IServiceProvider
}
