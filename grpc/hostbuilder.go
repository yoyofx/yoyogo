package grpc

import "github.com/yoyofx/yoyogo/abstractions"

type HostBuilder struct {
	abstractions.HostBuilder
}

func NewWebHostBuilder() *HostBuilder {
	builder := &HostBuilder{
		abstractions.HostBuilder{
			Context:   &abstractions.HostBuilderContext{HostingEnvironment: &abstractions.HostEnvironment{}},
			Decorator: NewHostBuilderDecorator(),
		},
	}

	return builder
}
