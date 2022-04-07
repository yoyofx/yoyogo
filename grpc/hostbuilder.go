package grpc

import "github.com/yoyofx/yoyogo/abstractions"

type HostBuilder struct {
	abstractions.HostBuilder
}

func NewHostBuilder() *HostBuilder {
	builder := &HostBuilder{
		abstractions.HostBuilder{
			Context:   &abstractions.HostBuilderContext{HostingEnvironment: &abstractions.HostEnvironment{}},
			Decorator: NewHostBuilderDecorator(),
		},
	}

	return builder
}

func (host *HostBuilder) UseConfiguration(configuration abstractions.IConfiguration) *HostBuilder {
	host.HostBuilder.UseConfiguration(configuration)
	return host
}

func (host *HostBuilder) Configure(configure func(*ApplicationBuilder)) *HostBuilder {
	host.HostBuilder.Configure(configure)
	return host
}
