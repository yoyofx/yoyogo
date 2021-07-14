package web

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type WebHostBuilderDecorator struct {
}

// OverrideConfigure is configure function by web application builder.
func (decorator WebHostBuilderDecorator) OverrideConfigure(configureFunc interface{}, builder abstractions.IApplicationBuilder) {
	configureWebAppFunc := configureFunc.(func(applicationBuilder *ApplicationBuilder))
	configureWebAppFunc(builder.(*ApplicationBuilder))
}

// OverrideNewApplicationBuilder create web application builder.
func (decorator WebHostBuilderDecorator) OverrideNewApplicationBuilder(context *abstractions.HostBuilderContext) abstractions.IApplicationBuilder {
	applicationBuilder := NewWebApplicationBuilder()
	applicationBuilder.SetHostBuildContext(context)
	return applicationBuilder
}

// OverrideNewHost Create WebHost.
func (decorator WebHostBuilderDecorator) OverrideNewHost(server abstractions.IServer, context *abstractions.HostBuilderContext) abstractions.IServiceHost {
	serverType := "fasthttp"
	if server == nil && context.HostConfiguration != nil {
		serverType = context.HostConfiguration.Server.ServerType
	}
	_ = context.ApplicationServices.GetServiceByName(&server, serverType)
	return NewWebHost(server, context)
}

// OverrideInnerConfigures inner configures for IOC
func (decorator WebHostBuilderDecorator) OverrideIOCInnerConfigures(serviceCollection *dependencyinjection.ServiceCollection) {
	serviceCollection.AddSingletonByNameAndImplements("fasthttp", NewFastHttpByConfig, new(abstractions.IServer))
	serviceCollection.AddSingletonByNameAndImplements("http", NewDefaultHttpByConfig, new(abstractions.IServer))
}

// NewWebHostBuilderDecorator WebHostBuilderDecorator.
func NewWebHostBuilderDecorator() WebHostBuilderDecorator {
	return WebHostBuilderDecorator{}
}
