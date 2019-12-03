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
	lifeConfigure      func(*ApplicationLife)
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

func (self *HostBuilder) OnApplicationLifeEvent(lifeConfigure func(*ApplicationLife)) *HostBuilder {
	self.lifeConfigure = lifeConfigure
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
	services := DependencyInjection.NewServiceCollection()

	self.context.hostingEnvironment.AppMode = "Dev"
	self.context.hostingEnvironment.DefaultAddress = ":8080"
	self.context.ApplicationCycle = NewApplicationLife()

	builder := NewApplicationBuilder(self.context)

	configures(self.context, services)

	for _, configure := range self.servicesconfigures {
		configure(services)
	}

	for _, configure := range self.configures {
		configure(builder)
	}

	for _, configure := range self.routeconfigures {
		configure(builder)
	}

	self.context.applicationServices = services.Build() //serviceProvider
	self.context.RequestDelegate = builder.Build()      // ServeHTTP(w http.ResponseWriter, r *http.Request)

	go self.lifeConfigure(self.context.ApplicationCycle)
	return NewWebHost(self.server, self.context)
}

func configures(hostContext *HostBuildContext, serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddSingleton(hostContext.ApplicationCycle)
}

func NewWebHostBuilder() *HostBuilder {
	return &HostBuilder{context: &HostBuildContext{hostingEnvironment: &HostEnv{}}}
}
