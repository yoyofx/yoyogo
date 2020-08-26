package Abstractions

import (
	"fmt"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/Abstractions/Configs"
	"github.com/yoyofx/yoyogo/Abstractions/Env"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net"
	"os"
	"strings"
)

// host builder
type HostBuilder struct {
	Server             IServer                                        // Server
	Context            *HostBuildContext                              // context of Host builder
	Decorator          IHostBuilderDecorator                          // host builder decorator or extension
	configures         []interface{}                                  // []func(IApplicationBuilder), configure function by application builder.
	servicesConfigures []func(*DependencyInjection.ServiceCollection) // configure function by ServiceCollection of DI.
	lifeConfigure      func(*ApplicationLife)                         // on application life event
}

// SetEnvironment set value(Dev,Test,Prod) by environment
func (host *HostBuilder) SetEnvironment(mode string) *HostBuilder {
	host.Context.HostingEnvironment.Profile = mode
	return host
}

// Configure function func(IApplicationBuilder)
func (host *HostBuilder) Configure(configure interface{}) *HostBuilder {
	host.configures = append(host.configures, configure)
	return host
}

func (host *HostBuilder) UseConfiguration(configuration IConfiguration) *HostBuilder {
	host.Context.Configuration = configuration
	host.Context.HostingEnvironment.Profile = configuration.GetProfile()
	section := host.Context.Configuration.GetSection("application")
	if section != nil {
		config := &Configs.HostConfig{}
		section.Unmarshal(config)
		host.Context.HostConfiguration = config
	}
	return host
}

// ConfigureServices configure function by ServiceCollection of DI.
func (host *HostBuilder) ConfigureServices(configure func(*DependencyInjection.ServiceCollection)) *HostBuilder {
	host.servicesConfigures = append(host.servicesConfigures, configure)
	return host
}

// OnApplicationLifeEvent on application life event
func (host *HostBuilder) OnApplicationLifeEvent(lifeConfigure func(*ApplicationLife)) *HostBuilder {
	host.lifeConfigure = lifeConfigure
	return host
}

// UseServer set IServer to host builder
func (host *HostBuilder) UseServer(server IServer) *HostBuilder {
	host.Server = server
	return host
}

// getLocalIP get localhost ip
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

// RunningHostEnvironmentSetting ,get running Env setting.
func RunningHostEnvironmentSetting(hostEnv *Context.HostEnvironment) {
	hostEnv.Host = getLocalIP()
	hostEnv.PID = os.Getpid()
}

//buildingHostEnvironmentSetting  build each configuration by init , such as file or Env or args ...
func buildingHostEnvironmentSetting(context *HostBuildContext) {
	hostEnv := context.HostingEnvironment
	hostEnv.Version = YoyoGo.Version
	hostEnv.Addr = DetectAddress("")
	config := context.HostConfiguration
	if config != nil {
		hostEnv.ApplicationName = config.Name
		if config.Server.Address != "" {
			hostEnv.Addr = config.Server.Address
		}
	}

	hostEnv.Port = strings.Replace(hostEnv.Addr, ":", "", -1)
	hostEnv.Args = os.Args

	if hostEnv.Profile == "" {
		hostEnv.Profile = Env.Dev
	}

}

// Build host
func (host *HostBuilder) Build() IServiceHost {
	services := DependencyInjection.NewServiceCollection()

	buildingHostEnvironmentSetting(host.Context)
	host.Context.ApplicationCycle = NewApplicationLife()

	innerConfigures(host.Context, services)
	for _, configure := range host.servicesConfigures {
		configure(services)
	}

	applicationBuilder := host.Decorator.OverrideNewApplicationBuilder(host.Context)

	for _, configure := range host.configures {
		//configure(applicationBuilder)
		host.Decorator.OverrideConfigure(configure, applicationBuilder)
	}

	host.Context.ApplicationServicesDef = services
	applicationBuilder.SetHostBuildContext(host.Context)
	host.Context.HostServices = services.Build()              //serviceProvider
	host.Context.RequestDelegate = applicationBuilder.Build() // ServeHTTP(w http.IResponseWriter, r *http.Request)
	host.Context.ApplicationServices = services.Build()       //serviceProvider

	if host.lifeConfigure != nil {
		go host.lifeConfigure(host.Context.ApplicationCycle)
	}

	return host.Decorator.OverrideNewHost(host.Server, host.Context)
}

// inner configures function for DI.
func innerConfigures(hostContext *HostBuildContext, serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddSingleton(func() *ApplicationLife { return hostContext.ApplicationCycle })
	serviceCollection.AddSingleton(func() *Context.HostEnvironment { return hostContext.HostingEnvironment })
}
