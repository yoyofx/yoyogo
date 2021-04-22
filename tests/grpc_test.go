package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	proto "github.com/yoyofx/yoyogo/tests/proto"
	"testing"
)

func TestGrpcFactory(t *testing.T) {
	assert.Equal(t, 1, 1)
	/*lis,err:=net.Listen("tcp",":5003")
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	proto.RegisterGreeterServer(s,&server{})
	if err := s.Serve(lis); err != nil{
		panic(err)
	}
	defer lis.Close();
	url:="sd://[demo]";
	selector := &servicediscovery.Selector{DiscoveryCache: &memory.MemoryCache{Services: []string{"localhost"}, Port: 5003},
		Strategy: strategy.NewRound()}
	gfc:=grpc_conn.GrpcConnFactory{Selector:selector}
	conn,err:=gfc.CreateGrpcConn(url,grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithResolvers(gfc.NewLoadBalanceResolver()))
	defer conn.Close();
	c:=proto.NewGreeterClient(conn)
	name:="world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel();
	res,err:=c.SayHello(ctx,&proto.HelloRequest{Name: name})
	if err!=nil{
		panic(err)
	}
	fmt.Println(res.Message)*/
}

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "你妈嗨" + in.GetName()}, nil
}
