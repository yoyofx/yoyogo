package grpc

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hostenv"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	IsTLS         bool
	Addr          string
	CertFile      string `mapstructure:"cert"`
	KeyFile       string `mapstructure:"key"`
	serverContext *ServerBuilderContext
}

func NewGrpcServerConfig(config hostenv.HttpServerConfig) *Server {
	return &Server{
		IsTLS:    config.CertFile != "",
		Addr:     config.Addr,
		CertFile: config.CertFile,
		KeyFile:  config.KeyFile,
	}
}

func (server *Server) GetAddr() string {
	return server.Addr
}

func (server *Server) Run(context *abstractions.HostBuilderContext) (e error) {
	addr := server.Addr
	if server.Addr == "" {
		addr = context.HostingEnvironment.Addr
	}
	_ = addr

	server.serverContext = context.RequestDelegate.(*ServerBuilderContext)

	for _, configure := range server.serverContext.serviceConfigures {
		configure(server.serverContext.server, server.serverContext.context)
	}

	reflection.Register(server.serverContext.server)

	lis, _ := net.Listen("tcp", addr)
	return server.serverContext.server.Serve(lis)
}

func (server *Server) Shutdown() {
	server.serverContext.server.GracefulStop()

	log.Fatal("Shutdown HTTP server...")
}
