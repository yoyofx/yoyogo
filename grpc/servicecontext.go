package grpc

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

type ServiceContext struct {
	ApplicationServices dependencyinjection.IServiceProvider
	Configuration       abstractions.IConfiguration
}

func NewServiceContext(sp dependencyinjection.IServiceProvider, config abstractions.IConfiguration) *ServiceContext {
	return &ServiceContext{ApplicationServices: sp, Configuration: config}
}

func (ctx *ServiceContext) Register(fn interface{}) {
	err := ctx.ApplicationServices.InvokeService(fn)
	if err != nil {
		panic(err)
	}
}
