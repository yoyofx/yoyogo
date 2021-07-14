package main

import (
	"context"
	grpconn "github.com/yoyofx/yoyogo/grpc/conn"
	pb "grpc-demo/proto/helloworld"
	"io"
	"log"
	"strconv"
)

type Api struct {
	client pb.GreeterClient
}

func NewHelloworldApi(factory *grpconn.Factory) *Api {
	clientConn, err := factory.CreateClientConn("grpc://public/[yoyogo_grpc_dev]")
	if err != nil {
		log.Println(err)
		return nil
	}

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
