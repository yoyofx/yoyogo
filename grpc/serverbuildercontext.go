package grpc

import "google.golang.org/grpc"

type ServerBuilderContext struct {
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	serverOption       []grpc.ServerOption
	server             *grpc.Server
}
