package grpc

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

type HostBuilderDecorator struct {
}

func (h HostBuilderDecorator) OverrideConfigure(configureFunc interface{}, builder abstractions.IApplicationBuilder) {
	configureWebAppFunc := configureFunc.(func(applicationBuilder *ApplicationBuilder))
	configureWebAppFunc(builder.(*ApplicationBuilder))
}

func (h HostBuilderDecorator) OverrideNewApplicationBuilder(context *abstractions.HostBuilderContext) abstractions.IApplicationBuilder {
	applicationBuilder := NewApplicationBuilder()
	applicationBuilder.SetHostBuildContext(context)
	return applicationBuilder
}

func (h HostBuilderDecorator) OverrideNewHost(server abstractions.IServer, context *abstractions.HostBuilderContext) abstractions.IServiceHost {
	serverType := "grpc"
	if server == nil && context.HostConfiguration != nil {
		serverType = context.HostConfiguration.Server.ServerType
	}
	_ = context.ApplicationServices.GetServiceByName(&server, serverType)
	return NewHost(server, context)
}

func (h HostBuilderDecorator) OverrideIOCInnerConfigures(serviceCollection *dependencyinjection.ServiceCollection) {
	serviceCollection.AddSingletonByNameAndImplements("grpc", NewGrpcServerConfig, new(abstractions.IServer))
}

func NewHostBuilderDecorator() HostBuilderDecorator {
	return HostBuilderDecorator{}
}
