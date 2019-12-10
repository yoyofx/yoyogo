package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/Router"
	"os"
)

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

func (self *HostBuilder) UseFastHttpByAddr(addr string) *HostBuilder {
	self.server = NewFastHttp(addr)
	return self
}

func (self *HostBuilder) UseFastHttp() *HostBuilder {
	self.server = NewFastHttp("")
	return self
}

func (self *HostBuilder) UseHttpByAddr(addr string) *HostBuilder {
	self.server = DefaultHttpServer(addr)
	return self
}

func (self *HostBuilder) UseHttp() *HostBuilder {
	self.server = DefaultHttpServer("")
	return self
}

func runningHostEnvironmentSetting(hostEnv *Context.HostEnvironment) {
	hostEnv.Port = detectAddress(hostEnv.Addr)
	hostEnv.PID = os.Getpid()
}

func buildingHostEnvironmentSetting(hostEnv *Context.HostEnvironment) {
	// build each configuration by init , such as file or env or args ...
	hostEnv.Args = os.Args
	hostEnv.ApplicationName = "app"
	hostEnv.Version = "v1.0.0"
	if hostEnv.Profile == "" {
		hostEnv.Profile = Context.Dev
	}
	hostEnv.Addr = ":8080"
}

func (self *HostBuilder) Build() WebHost {
	services := DependencyInjection.NewServiceCollection()

	buildingHostEnvironmentSetting(self.context.hostingEnvironment)
	self.context.ApplicationCycle = NewApplicationLife()

	configures(self.context, services)
	for _, configure := range self.servicesconfigures {
		configure(services)
	}

	self.context.applicationServices = services.Build() //serviceProvider

	builder := NewApplicationBuilder(self.context)

	for _, configure := range self.configures {
		configure(builder)
	}

	for _, configure := range self.routeconfigures {
		configure(builder.routerBuilder)
	}

	self.context.RequestDelegate = builder.Build() // ServeHTTP(w http.ResponseWriter, r *http.Request)

	go self.lifeConfigure(self.context.ApplicationCycle)
	return NewWebHost(self.server, self.context)
}

func configures(hostContext *HostBuildContext, serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddSingleton(func() *ApplicationLife { return hostContext.ApplicationCycle })
	serviceCollection.AddSingleton(func() *Context.HostEnvironment { return hostContext.hostingEnvironment })
}

func NewWebHostBuilder() *HostBuilder {
	return &HostBuilder{context: &HostBuildContext{hostingEnvironment: &Context.HostEnvironment{}}}
}
