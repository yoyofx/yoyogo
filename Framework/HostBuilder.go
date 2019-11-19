package YoyoGo

import "github.com/maxzhang1985/yoyogo/Router"

type HostBuilder struct {
	server          IServer
	configures      []func(*ApplicationBuilder)
	routeconfigures []func(Router.IRouterBuilder)
}

func (self HostBuilder) Configure(configure func(*ApplicationBuilder)) HostBuilder {
	self.configures = append(self.configures, configure)
	return self
}

func (self HostBuilder) UseRouter(configure func(Router.IRouterBuilder)) HostBuilder {
	self.routeconfigures = append(self.routeconfigures, configure)
	return self
}

func (self HostBuilder) UseServer(server IServer) HostBuilder {
	self.server = server
	return self
}

func (self HostBuilder) Build() WebHost {
	builder := NewApplicationBuilder()

	for _, configure := range self.configures {
		configure(builder)
	}

	for _, configure := range self.routeconfigures {
		configure(builder)
	}

	return NewWebHost(self.server, builder.Build())

}

func NewWebHostBuilder() HostBuilder {
	return HostBuilder{}
}
