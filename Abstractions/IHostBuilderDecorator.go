package Abstractions

// IHostBuilderDecorator Host Builder decorator or extension
type IHostBuilderDecorator interface {

	// OverrideConfigure is configure function by application builder.
	OverrideConfigure(configureFunc interface{}, builder IApplicationBuilder)
	// OverrideNewApplicationBuilder create application builder.
	OverrideNewApplicationBuilder() IApplicationBuilder
	// OverrideNewHost Create IServiceHost.
	OverrideNewHost(server IServer, context *HostBuildContext) IServiceHost
}
