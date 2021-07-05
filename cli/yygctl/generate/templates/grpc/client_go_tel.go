package grpc

const Client_Main_Tel = `
package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofx/yoyogo/console"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/nacos"
)

func main() {
	configuration := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	console.NewHostBuilder().
		UseConfiguration(configuration).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			hosting.AddHostService(collection, NewClientService)
			collection.AddSingleton(NewHelloworldApi)
			//register sd for nacos
			nacos.UseServiceDiscovery(collection)
		}).
		Build().Run()
}
`

const Client_Service_Tel = `
package main

import (
	"fmt"
	pb "{{.ModelName}}/proto/helloworld"
)

type ClientService struct {
	helloworldApi *Api
}

func NewClientService(api *Api) *ClientService {
	return &ClientService{helloworldApi: api}
}

func (s *ClientService) Run() error {
	fmt.Println("host service Running")
	err := s.helloworldApi.SayRecord(&pb.HelloRequest{})
	if err != nil {
		return err
	}

	return nil
}

func (s *ClientService) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}

`

const Client_Api_Tel = `
package {{.CurrentModelName}}

import (
	"context"
	grpconn "github.com/yoyofx/yoyogo/grpc/conn"
	pb "{{.ModelName}}/proto/helloworld"
	"io"
	"log"
	"strconv"
)

type Api struct {
	client pb.GreeterClient
}

func NewHelloworldApi(factory *grpconn.Factory) *Api {
	clientConn, _ := factory.CreateClientConn("grpc://public/[yoyogo_grpc_dev]")
	client := pb.NewGreeterClient(clientConn)
	return &Api{client: client}
}

func (hw *Api) SayHello() error {
	resp, _ := hw.client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func (hw *Api) SayList(r *pb.HelloRequest) error {
	stream, _ := hw.client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}

func (hw *Api) SayRecord(r *pb.HelloRequest) error {
	stream, _ := hw.client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		r.Name = strconv.Itoa(n)
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("client call resp: %v", resp)
	return nil
}

func (hw *Api) SayRoute(r *pb.HelloRequest) error {
	stream, _ := hw.client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		r.Name = "client" + strconv.Itoa(n)
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("client callback resp: %v", resp)
	}

	_ = stream.CloseSend()

	return nil
}

`
