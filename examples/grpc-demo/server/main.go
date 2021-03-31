package main

import (
	"context"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-demo/proto/helloworld"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			UnaryAccessLog,
		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(),
			StreamAccessLog,
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &GreeterServer{})
	reflection.Register(server)
	lis, _ := net.Listen("tcp", ":31127")
	_ = server.Serve(lis)
}

func UnaryAccessLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestLog := "access request log: method: %s, begin_time: %d, request: %v"
	beginTime := time.Now().Local().Unix()
	log.Printf(requestLog, info.FullMethod, beginTime, req)

	resp, err := handler(ctx, req)

	responseLog := "access response log: method: %s, begin_time: %d, end_time: %d, response: %v"
	endTime := time.Now().Local().Unix()
	log.Printf(responseLog, info.FullMethod, beginTime, endTime, resp)
	return resp, err
}

func StreamAccessLog(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	requestLog := "stream access request log: method: %s, begin_time: %d , client streaming: %v, server streaming: %v"
	beginTime := time.Now().Local().Unix()
	log.Printf(requestLog, info.FullMethod, beginTime, info.IsClientStream, info.IsServerStream)

	err := handler(srv, stream)

	responseLog := "stream access response log: method: %s, begin_time: %d, end_time: %d"
	endTime := time.Now().Local().Unix()
	log.Printf(responseLog, info.FullMethod, beginTime, endTime)
	return err
}

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("server.SayHello")
	return &pb.HelloReply{Message: "hello.world.at.server: " + r.Name}, nil
}

// Server-side streaming RPC
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "server-side.streaming.hello.list." + strconv.Itoa(n)})
	}
	return nil
}

// Client-side streaming
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "client-side.streaming.say.record.end.at.server"})
		}
		if err != nil {
			return err
		}

		log.Printf("server recv resp: %v", resp)
	}

	return nil
}

//     // Bidirectional streaming RPC 双向流式 RPC
func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "bidirectional.streaming.say.route.at.server." + strconv.Itoa(n)})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("recv client resp at server: %v", resp)
	}
}
