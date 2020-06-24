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
func (decorator WebHostBuilderDecorator) OverrideNewApplicationBuilder() Abstractions.IApplicationBuilder {
	return NewWebApplicationBuilder()
}

// OverrideNewHost Create WebHost.
func (decorator WebHostBuilderDecorator) OverrideNewHost(server Abstractions.IServer, context *Abstractions.HostBuildContext) Abstractions.IServiceHost {
	return NewWebHost(server, context)
}

// NewWebHostBuilderDecorator WebHostBuilderDecorator.
func NewWebHostBuilderDecorator() WebHostBuilderDecorator {
	return WebHostBuilderDecorator{}
}
