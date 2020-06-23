package YoyoGo

import (
	"fmt"
	Yoyo "github.com/maxzhang1985/yoyogo"
	"github.com/maxzhang1985/yoyogo/Abstract"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
	"net"
	"os"
)

type WebHostBuilder struct {
	server             Abstract.IServer
	context            *Abstract.HostBuildContext
	configures         []func(*WebApplicationBuilder)
	servicesConfigures []func(*DependencyInjection.ServiceCollection)
	lifeConfigure      func(*Abstract.ApplicationLife)
}

func (self *WebHostBuilder) Configure(configure func(*WebApplicationBuilder)) *WebHostBuilder {
	self.configures = append(self.configures, configure)
	return self
}

func (self *WebHostBuilder) ConfigureServices(configure func(*DependencyInjection.ServiceCollection)) *WebHostBuilder {
	self.servicesConfigures = append(self.servicesConfigures, configure)
	return self
}

func (self *WebHostBuilder) OnApplicationLifeEvent(lifeConfigure func(*Abstract.ApplicationLife)) *WebHostBuilder {
	self.lifeConfigure = lifeConfigure
	return self
}

func (self *WebHostBuilder) UseServer(server Abstract.IServer) *WebHostBuilder {
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
	hostEnv.Port = Abstract.DetectAddress(hostEnv.Addr)
	hostEnv.PID = os.Getpid()
}

func buildingHostEnvironmentSetting(hostEnv *Context.HostEnvironment) {
	// build each configuration by init , such as file or env or args ...
	hostEnv.ApplicationName = "app"
	hostEnv.Version = Yoyo.Version
	hostEnv.Addr = ":8080"

	hostEnv.Args = os.Args

	if hostEnv.Profile == "" {
		hostEnv.Profile = Context.Dev
	}

}

func (self *WebHostBuilder) Build() WebHost {
	services := DependencyInjection.NewServiceCollection()

	buildingHostEnvironmentSetting(self.context.HostingEnvironment)
	self.context.ApplicationCycle = Abstract.NewApplicationLife()

	configures(self.context, services)
	for _, configure := range self.servicesConfigures {
		configure(services)
	}

	applicationBuilder := NewWebApplicationBuilder()

	for _, configure := range self.configures {
		configure(applicationBuilder)
	}

	self.context.ApplicationServicesDef = services
	applicationBuilder.SetHostBuildContext(self.context)
	self.context.RequestDelegate = applicationBuilder.Build() // ServeHTTP(w http.ResponseWriter, r *http.Request)
	self.context.ApplicationServices = services.Build()       //serviceProvider

	go self.lifeConfigure(self.context.ApplicationCycle)
	return NewWebHost(self.server, self.context)
}

func configures(hostContext *Abstract.HostBuildContext, serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddSingleton(func() *Abstract.ApplicationLife { return hostContext.ApplicationCycle })
	serviceCollection.AddSingleton(func() *Context.HostEnvironment { return hostContext.HostingEnvironment })
}

func NewWebHostBuilder() *WebHostBuilder {
	return &WebHostBuilder{context: &Abstract.HostBuildContext{HostingEnvironment: &Context.HostEnvironment{}}}
}
