package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/Router"
)

type HostEnv struct {
	ApplicationName string
	DefaultAddress  string
	Version         string
	AppMode         string
	Args            []string
	Addr            string
	Port            string
	PID             int
}

type HostBuilder struct {
	server             IServer
	context            *HostBuildContext
	configures         []func(*ApplicationBuilder)
	routeconfigures    []func(Router.IRouterBuilder)
	servicesconfigures []func(*DependencyInjection.ServiceCollection)
}

func (self *HostBuilder) Configure(configure func(*ApplicationBuilder)) *HostBuilder {
	self.configures = append(self.configures, configure)
	return self
}

func (self *HostBuilder) UseRouter(configure func(Router.IRouterBuilder)) *HostBuilder {
	self.routeconfigures = append(self.routeconfigures, configure)
	return self
}

func (self *HostBuilder) ConfigureServices(configure func(*DependencyInjection.ServiceCollection)) *HostBuilder {
	self.servicesconfigures = append(self.servicesconfigures, configure)
	return self
}

func (self *HostBuilder) UseServer(server IServer) *HostBuilder {
	self.server = server
	return self
}

func (self *HostBuilder) UseFastHttp(addr string) *HostBuilder {
	self.server = NewFastHttp(addr)
	return self
}

func (self *HostBuilder) UseHttp(addr string) *HostBuilder {
	self.server = DefaultHttpServer(addr)
	return self
}

func (self *HostBuilder) Build() WebHost {
	self.context.hostingEnvironment.AppMode = "Dev"
	self.context.hostingEnvironment.DefaultAddress = ":8080"

	services := DependencyInjection.NewServiceCollection()

	builder := NewApplicationBuilder(self.context)

	for _, configure := range self.configures {
		configure(builder)
	}

	for _, configure := range self.routeconfigures {
		configure(builder)
	}

	for _, configure := range self.servicesconfigures {
		configure(services)
	}

	serviceProvider := services.Build()
	self.context.applicationServices = serviceProvider

	return NewWebHost(self.server, builder.Build(), self.context)

}

func NewWebHostBuilder() *HostBuilder {
	return &HostBuilder{context: &HostBuildContext{hostingEnvironment: &HostEnv{}}}
}
