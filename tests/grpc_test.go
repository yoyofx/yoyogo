package tests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/grpc/conn"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/memory"
	proto "github.com/yoyofx/yoyogo/tests/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"testing"
	"time"
)

func TestGrpcFactory(t *testing.T) {
	assert.Equal(t, 1, 1)
	lis, err := net.Listen("tcp", ":25003")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	cache := &memory.MemoryCache{Services: []string{"localhost"}, Port: 25003}

	gfc := conn.NewFactory(cache)
	clientConn, err := gfc.CreateClientConn("grpc://public/[yoyogo_grpc_dev]")

	c := proto.NewGreeterClient(clientConn)

	name := "world"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, err := c.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		panic(err)
	}
	fmt.Println("client response: " + res.Message)
	assert.Equal(t, "hello world", res.Message)
	wg.Done()

	wg.Wait()
	cancel()

}

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println("server request:" + in.GetName())

	return &proto.HelloReply{Message: "hello " + in.GetName()}, nil
}
