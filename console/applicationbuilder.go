package console

import "github.com/yoyofx/yoyogo/abstractions"

type ApplicationBuilder struct {
	hostBuilderContext *abstractions.HostBuilderContext
}

func (builder *ApplicationBuilder) Build() interface{} {
	return builder
}

func (builder *ApplicationBuilder) SetHostBuildContext(context *abstractions.HostBuilderContext) {
	builder.hostBuilderContext = context
}

func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{}
}
