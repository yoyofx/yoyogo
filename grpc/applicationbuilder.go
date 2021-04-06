package grpc

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/grpc/interceptors"
	"google.golang.org/grpc"
)

type ApplicationBuilder struct {
	hostBuilderContext *abstractions.HostBuilderContext
	serviceConfigures  []func(server *grpc.Server, ctx *ServiceContext)
	serverContext      *ServerBuilderContext
}

func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{serverContext: &ServerBuilderContext{}}
}

func (builder *ApplicationBuilder) AddServerOption(serverOption grpc.ServerOption) *ApplicationBuilder {
	builder.serverContext.serverOption = append(builder.serverContext.serverOption, serverOption)
	return builder
}

func (builder *ApplicationBuilder) AddUnaryServerInterceptor(interceptor ...grpc.UnaryServerInterceptor) *ApplicationBuilder {
	builder.serverContext.unaryInterceptors = append(builder.serverContext.unaryInterceptors, interceptor...)
	return builder
}

func (builder *ApplicationBuilder) AddStreamServerInterceptor(interceptor ...grpc.StreamServerInterceptor) *ApplicationBuilder {
	builder.serverContext.streamInterceptors = append(builder.serverContext.streamInterceptors, interceptor...)
	return builder
}

func (builder *ApplicationBuilder) AddGrpcService(configure func(server *grpc.Server, ctx *ServiceContext)) *ApplicationBuilder {
	builder.serviceConfigures = append(builder.serviceConfigures, configure)
	return builder
}

func (builder *ApplicationBuilder) Build() interface{} {

	logger := interceptors.NewLogger()
	builder.AddUnaryServerInterceptor(logger.UnaryServerInterceptor())
	builder.AddUnaryServerInterceptor(grpc_recovery.UnaryServerInterceptor())
	builder.AddStreamServerInterceptor(logger.StreamServerInterceptor())
	builder.AddStreamServerInterceptor(grpc_recovery.StreamServerInterceptor())

	// add interceptors
	builder.AddServerOption(grpc.ChainUnaryInterceptor(builder.serverContext.unaryInterceptors...))
	builder.AddServerOption(grpc.ChainStreamInterceptor(builder.serverContext.streamInterceptors...))
	// add server options
	opts := builder.serverContext.serverOption

	server := grpc.NewServer(opts...)
	svrCtx := NewServiceContext(builder.hostBuilderContext.HostServices, builder.hostBuilderContext.Configuration)
	for _, configure := range builder.serviceConfigures {
		configure(server, svrCtx)
	}

	builder.serverContext.server = server
	return builder.serverContext
}

func (builder *ApplicationBuilder) SetHostBuildContext(context *abstractions.HostBuilderContext) {
	builder.hostBuilderContext = context
	if builder.hostBuilderContext.ApplicationServicesDef != nil {
		builder.innerConfigures()
	}
}

//  this time is not build host.context.HostServices , that add services define
func (builder *ApplicationBuilder) innerConfigures() {
	//builder.hostBuilderContext.
	//	ApplicationServicesDef.
	//	AddSingletonByNameAndImplements("grpcService", f, new(f))
}
