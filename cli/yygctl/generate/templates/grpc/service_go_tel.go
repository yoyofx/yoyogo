package grpc

const Greeter_Server_Tel = `
package services

import (
	"context"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	pb "{{.ModelName}}/proto/helloworld"
	"io"
	"strconv"
)

type GreeterServer struct {
	log  xlog.ILogger
	demo *IOCDemo
}

func NewGreeterServer(dd *IOCDemo) pb.GreeterServer {
	log := xlog.GetXLogger("GreeterServer")
	log.SetCustomLogFormat(nil)
	return &GreeterServer{log: log, demo: dd}
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.Debug("server.SayHello ," + s.demo.Print())
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

		s.log.Debug("server recv resp: %v", resp)
	}
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
		s.log.Debug("recv client resp at server: %v", resp)
	}
}

`

const Demo_Tel = `
package services

type IOCDemo struct {
}

func NewIOCDemo() *IOCDemo {
	return &IOCDemo{}
}

func (demo *IOCDemo) Print() string {
	return "IOC Demo test"
}
`

const ServiceConfig_Tel = `
yoyogo:
  application:
    name: yoyogo_grpc_dev
    metadata: "develop"
    server:
      type: "grpc"
      address: ":31127"
      tls:
        cert: ""
        key: ""
  cloud:
    apm:
      skyworking:
        address: localhost:11800
    discovery:
      cache:
        ttl: 30     # seconds
      type: "nacos"
      metadata:
        url: "120.53.133.30"
        port: 80
        namespace: "public"
        group: ""
        cluster: ""
        auth:
          enable: true
          username: "root"
          password: "1234"
#      type: "consul"
#      metadata:
#        address: "localhost:8500"
#        health_check: "/actuator/health"
#        tags: [""]
#      type: "eureka"
#      metadata:
#         datacenter: ""
#         address:
#           - "http://sd.yoyogo.run/eureka"
#         ttl: 30   # seconds
#      type: "etcd"
#      metadata:
#        address:
#          - "81.70.154.55:32379"
#        namespace: "public"
#        ttl: 60
#        auth:
#          enable: true
#          username: "root"
#          password: "5z1fLbYw8A"
  datasource:
      pool:
        init_cap: 2
        max_cap: 5
        idle_timeout: 5
      db:
        name: db1
        type: mysql
        url: tcp(cdb-amqub3mo.bj.tencentcdb.com:10042)/yoyoBlog?charset=utf8&parseTime=True
        username: root
        password: 1234abcd
        debug: true
      redis:
        name: reids1
        url: 62.234.6.120:31379
        password:
        db: 0

`
