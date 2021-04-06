package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	yrpc "github.com/yoyofx/yoyogo/grpc"
	"google.golang.org/grpc"
	pb "grpc-demo/proto/helloworld"
	"grpc-demo/services"
)

func main() {
	configuration := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	hosting := yrpc.NewHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *yrpc.ApplicationBuilder) {
			//app.AddUnaryServerInterceptor( logger.UnaryServerInterceptor() )
			//app.AddStreamServerInterceptor( logger.StreamServerInterceptor() )
			app.AddGrpcService(func(server *grpc.Server, ctx *yrpc.ServiceContext) {
				pb.RegisterGreeterServer(server, services.NewGreeterServer())
			})

		}).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {

		}).Build()

	hosting.Run()
}
