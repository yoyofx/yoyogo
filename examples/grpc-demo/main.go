package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	yrpc "github.com/yoyofx/yoyogo/grpc"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/nacos"
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
				ctx.Register(pb.RegisterGreeterServer) // register grpc service
			})

		}).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			collection.AddSingleton(services.NewGreeterServer) // add grpc service
			collection.AddSingleton(services.NewIOCDemo)
			nacos.UseServiceDiscovery(collection)
		}).Build()

	hosting.Run()
}
