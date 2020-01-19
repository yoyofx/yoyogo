package YoyoGo

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"net"
	"os"
)

type HostBuilder struct {
	server             IServer
	context            *HostBuildContext
	configures         []func(*ApplicationBuilder)
	servicesConfigures []func(*DependencyInjection.ServiceCollection)
	lifeConfigure      func(*ApplicationLife)
}

func (self *HostBuilder) Configure(configure func(*ApplicationBuilder)) *HostBuilder {
	self.configures = append(self.configures, configure)
	return self
}

func (self *HostBuilder) ConfigureServices(configure func(*DependencyInjection.ServiceCollection)) *HostBuilder {
	self.servicesConfigures = append(self.servicesConfigures, configure)
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

func getLocalIP() string {
	var localIp string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIp = ipnet.IP.String()
				break
			}
		}
	}
	return localIp
}

func runningHostEnvironmentSetting(hostEnv *Context.HostEnvironment) {
	hostEnv.Host = getLocalIP()
	hostEnv.Port = detectAddress(hostEnv.Addr)
	hostEnv.PID = os.Getpid()
}

func buildingHostEnvironmentSetting(hostEnv *Context.HostEnvironment) {
	// build each configuration by init , such as file or env or args ...
	hostEnv.ApplicationName = "app"
	hostEnv.Version = Version
	hostEnv.Addr = ":8080"

	hostEnv.Args = os.Args

	if hostEnv.Profile == "" {
		hostEnv.Profile = Context.Dev
	}

}

func (self *HostBuilder) Build() WebHost {
	services := DependencyInjection.NewServiceCollection()

	buildingHostEnvironmentSetting(self.context.hostingEnvironment)
	self.context.ApplicationCycle = NewApplicationLife()

	configures(self.context, services)
	for _, configure := range self.servicesConfigures {
		configure(services)
	}

	applicationBuilder := NewApplicationBuilder()

	for _, configure := range self.configures {
		configure(applicationBuilder)
	}

	self.context.applicationServicesDef = services
	applicationBuilder.SetHostBuildContext(self.context)
	self.context.RequestDelegate = applicationBuilder.Build() // ServeHTTP(w http.ResponseWriter, r *http.Request)
	self.context.applicationServices = services.Build()       //serviceProvider

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
