package conn

import (
	"context"
	"errors"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"strconv"
	"strings"
	"time"
)

type Factory struct {
	discoveryCache servicediscovery.Cache
}

func NewFactory(cache servicediscovery.Cache) *Factory {
	return &Factory{discoveryCache: cache}
}

type LoadBalanceResolver struct {
	discoveryCache servicediscovery.Cache
	serverName     string
}

func (gcf *Factory) CreateClientConn(serverUrl string, grpcOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if grpcOpts == nil {
		grpcOpts = append(grpcOpts, grpc.WithInsecure(), grpc.WithBlock(),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
			grpc.WithResolvers(gcf.NewLoadBalanceResolver()),
		)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)
	conn, err := grpc.DialContext(ctx, serverUrl, grpcOpts...)
	return conn, err
}

func (gcf *Factory) NewLoadBalanceResolver() *LoadBalanceResolver {
	return &LoadBalanceResolver{
		discoveryCache: gcf.discoveryCache,
	}
}

// ResolveNow 实现了 resolver.Resolver.ResolveNow 方法
func (*LoadBalanceResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// Close 实现了 resolver.Resolver.Close 方法
func (lr *LoadBalanceResolver) Close() {
}

func (lr *LoadBalanceResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	serverName := strings.Split(strings.Split(target.Endpoint(), "[")[1], "]")[0]
	service, err := lr.discoveryCache.GetService(serverName)
	if err != nil {
		return nil, err
	}
	if len(service.Nodes) == 0 {
		return nil, errors.New("this service don't have any instance")
	}
	addressList := make([]resolver.Address, len(service.Nodes))
	for _, item := range service.Nodes {
		addressList = append(addressList, resolver.Address{
			Addr: item.GetHost() + ":" + strconv.FormatUint(item.GetPort(), 10),
		})
	}
	cc.UpdateState(resolver.State{Addresses: addressList})
	return &LoadBalanceResolver{}, nil
}

// Scheme 实现了 resolver.Builder.Scheme 方法
// Scheme 方法定义了 sd resolver 的协议名
func (*LoadBalanceResolver) Scheme() string {
	return "grpc"
}
