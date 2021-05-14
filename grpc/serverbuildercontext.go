package grpc

import "google.golang.org/grpc"

type ServerBuilderContext struct {
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	serverOption       []grpc.ServerOption
	serviceConfigures  []func(server *grpc.Server, ctx *ServiceContext)
	context            *ServiceContext
	server             *grpc.Server
}
