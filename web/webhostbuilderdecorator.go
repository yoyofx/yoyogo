package web

import (
	"github.com/yoyofx/yoyogo/abstractions"
)

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type WebHostBuilderDecorator struct {
}

// OverrideConfigure is configure function by web application builder.
func (decorator WebHostBuilderDecorator) OverrideConfigure(configureFunc interface{}, builder abstractions.IApplicationBuilder) {
	configureWebAppFunc := configureFunc.(func(applicationBuilder *WebApplicationBuilder))
	configureWebAppFunc(builder.(*WebApplicationBuilder))
}

// OverrideNewApplicationBuilder create web application builder.
func (decorator WebHostBuilderDecorator) OverrideNewApplicationBuilder(context *abstractions.HostBuilderContext) abstractions.IApplicationBuilder {
	applicationBuilder := NewWebApplicationBuilder()
	applicationBuilder.SetHostBuildContext(context)
	return applicationBuilder
}

// OverrideNewHost Create WebHost.
func (decorator WebHostBuilderDecorator) OverrideNewHost(server abstractions.IServer, context *abstractions.HostBuilderContext) abstractions.IServiceHost {
	if server == nil && context.HostConfiguration != nil {
		//section := context.Configuration.GetSection("yoyogo.application.server")
		serverType := context.HostConfiguration.Server.ServerType
		address := context.HostConfiguration.Server.Address
		if serverType == "fasthttp" {
			server = NewFastHttp(address)
		} else if serverType == "http" {
			server = DefaultHttpServer(address)
		}

	} else {
		server = NewFastHttp(DefaultAddress)
	}
	return NewWebHost(server, context)
}

// NewWebHostBuilderDecorator WebHostBuilderDecorator.
func NewWebHostBuilderDecorator() WebHostBuilderDecorator {
	return WebHostBuilderDecorator{}
}
