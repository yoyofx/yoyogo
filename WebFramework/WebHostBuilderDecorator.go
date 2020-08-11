package YoyoGo

import "github.com/yoyofx/yoyogo/Abstractions"

type WebHostBuilderDecorator struct {
}

// OverrideConfigure is configure function by web application builder.
func (decorator WebHostBuilderDecorator) OverrideConfigure(configureFunc interface{}, builder Abstractions.IApplicationBuilder) {
	configureWebAppFunc := configureFunc.(func(applicationBuilder *WebApplicationBuilder))
	configureWebAppFunc(builder.(*WebApplicationBuilder))
}

// OverrideNewApplicationBuilder create web application builder.
func (decorator WebHostBuilderDecorator) OverrideNewApplicationBuilder(context *Abstractions.HostBuildContext) Abstractions.IApplicationBuilder {
	applicationBuilder := NewWebApplicationBuilder()
	applicationBuilder.SetHostBuildContext(context)
	return applicationBuilder
}

// OverrideNewHost Create WebHost.
func (decorator WebHostBuilderDecorator) OverrideNewHost(server Abstractions.IServer, context *Abstractions.HostBuildContext) Abstractions.IServiceHost {
	if server == nil && context.HostConfiguration != nil {
		section := context.HostConfiguration.GetSection("application.server.web")
		if section != nil {
			serverType := section.Get("type").(string)
			address := section.Get("address").(string)
			if serverType == "fasthttp" {
				server = NewFastHttp(address)
			} else if serverType == "http" {
				server = DefaultHttpServer(address)
			}
		}
	}
	return NewWebHost(server, context)
}

// NewWebHostBuilderDecorator WebHostBuilderDecorator.
func NewWebHostBuilderDecorator() WebHostBuilderDecorator {
	return WebHostBuilderDecorator{}
}
