package main

import (
	"context"
	"google.golang.org/grpc"
	//_ "google.golang.org/grpc/balancer/roundrobin"
	pb "grpc-demo/proto/helloworld"
	"io"
	"log"
	"strconv"
)

func main() {

	conn, err := grpc.Dial("sd://public/demo1",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithResolvers(NewResolverBuilder()),
	)
	// call java service
	//conn, err := grpc.Dial("localhost:6565",
	//	grpc.WithInsecure(),
	//)
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewGreeterClient(conn)
	//_ = SayHello(client)
	//
	//_ = SayList(client, &pb.HelloRequest{})

	_ = SayRecord(client, &pb.HelloRequest{})
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
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

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		r.Name = strconv.Itoa(n)
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("client call resp: %v", resp)
	return nil
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
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
