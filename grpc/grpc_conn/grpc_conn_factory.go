package grpc_conn

import (
	"errors"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"strings"
)

type GrpcConnFactory struct {
	Selector *servicediscovery.Selector
}

type LoadBalanceResolver struct {
	selector   *servicediscovery.Selector
	serverName string
}

func (gcf *GrpcConnFactory) CreateGrpcConn(serverUrl string, grpcOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	serverName := strings.Split(strings.Split(serverUrl, "[")[1], "]")[0]
	if serverName == "" {
		return nil, errors.New("url don't contains serveName")
	}
	conn, err := grpc.Dial(serverUrl,
		grpcOpts...,
	)
	return conn, err
}

func (gcf *GrpcConnFactory) NewLoadBalanceResolver() *LoadBalanceResolver {
	return &LoadBalanceResolver{
		selector: gcf.Selector,
	}
}

// ResolveNow 实现了 resolver.Resolver.ResolveNow 方法
func (*LoadBalanceResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// Close 实现了 resolver.Resolver.Close 方法
func (r *LoadBalanceResolver) Close() {
}

func (lr *LoadBalanceResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	service, err := lr.selector.DiscoveryCache.GetService(lr.serverName)
	if err != nil {
		return nil, err
	}
	if len(service.Nodes) == 0 {
		return nil, errors.New("this service don't have any instance")
	}
	addressList := make([]resolver.Address, len(service.Nodes))
	for _, item := range service.Nodes {
		addressList = append(addressList, resolver.Address{
			Addr: item.GetHost() + ":" + string(item.GetPort()),
		})
	}
	cc.UpdateState(resolver.State{Addresses: addressList})
	return &LoadBalanceResolver{}, nil
}

// Scheme 实现了 resolver.Builder.Scheme 方法
// Scheme 方法定义了 sd resolver 的协议名
func (*LoadBalanceResolver) Scheme() string {
	return "sd"
}
