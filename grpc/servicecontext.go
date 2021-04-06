package grpc

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

type ServiceContext struct {
	ApplicationServices dependencyinjection.IServiceProvider
	Configuration       abstractions.IConfiguration
}

func NewServiceContext(sp dependencyinjection.IServiceProvider, config abstractions.IConfiguration) *ServiceContext {
	return &ServiceContext{ApplicationServices: sp, Configuration: config}
}
