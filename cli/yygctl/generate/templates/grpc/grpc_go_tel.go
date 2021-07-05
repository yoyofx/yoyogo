package grpc

const Main_Tel = `
package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	yrpc "github.com/yoyofx/yoyogo/grpc"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/nacos"
	"google.golang.org/grpc"
	pb "{{.ModelName}}/proto/helloworld"
	"{{.ModelName}}/services"
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

`

const Mod_Tel = `
module {{.ModelName}}

go 1.16

require (
	github.com/golang/protobuf v1.5.1 
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/yoyofx/yoyogo {{.Version}}
	golang.org/x/net v0.0.0-20210326220855-61e056675ecf 
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181 
	golang.org/x/text v0.3.5 
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.26.0
)
`

const Config_Tel = `
yoyogo:
  application:
    name: yoyogo_grpc_clientdemo    # go build grpc-demo/client --profile=test
    metadata: "grpc client demo"
    server:
      type: "console"
`
