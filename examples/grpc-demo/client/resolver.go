package main

import "google.golang.org/grpc/resolver"

type sdResolver struct{}

// ResolveNow 实现了 resolver.Resolver.ResolveNow 方法
func (*sdResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// Close 实现了 resolver.Resolver.Close 方法
func (r *sdResolver) Close() {}

func NewResolverBuilder() *SDResolverBuilder {
	return &SDResolverBuilder{}
}

type SDResolverBuilder struct{}

func (*SDResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	var newAddrs []resolver.Address
	list := []string{
		"172.12.0.22:31127",
		"172.12.0.23:31127",
		"172.12.0.24:31127",
		"127.0.0.1:31127",
		"localhost:31127",
	}

	for _, item := range list {
		newAddrs = append(newAddrs, resolver.Address{
			Addr: item,
		})
	}
	cc.UpdateState(resolver.State{Addresses: newAddrs})

	return &sdResolver{}, nil
}

// Scheme 实现了 resolver.Builder.Scheme 方法
// Scheme 方法定义了 sd resolver 的协议名
func (*SDResolverBuilder) Scheme() string {
	return "sd"
}
