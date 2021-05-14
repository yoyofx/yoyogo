package abstractions

import (
	"fmt"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/abstractions/hostenv"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"net"
	"os"
	"strings"
)

// host builder
type HostBuilder struct {
	Server             IServer                                        // Server
	Context            *HostBuilderContext                            // context of Host builder
	Decorator          IHostBuilderDecorator                          // host builder decorator or extension
	configures         []interface{}                                  // []func(IApplicationBuilder), configure function by application builder.
	servicesConfigures []func(*dependencyinjection.ServiceCollection) // configure function by ServiceCollection of DI.
	lifeConfigure      func(*ApplicationLife)                         // on application life event
}

// SetEnvironment set value(Dev,tests,Prod) by environment
func (host *HostBuilder) SetEnvironment(mode string) *HostBuilder {
	host.Context.HostingEnvironment.Profile = mode
	return host
}

func (host *HostBuilder) UseStartup(startupFunc func() IStartup) *HostBuilder {
	if startupFunc != nil {
		startup := startupFunc()
		host.ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			startup.ConfigureServices(collection)
		})
	}
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
	section := host.Context.Configuration.GetSection("yoyogo.application")
	if section != nil {
		config := &hostenv.HostConfig{}
		section.Unmarshal(config)
		portInterface := configuration.Get("port")
		if portInterface != nil && portInterface.(string) != "" {
			config.Server.Address = ":" + portInterface.(string)
		}
		appName := configuration.GetString("app")
		if appName != "" {
			config.Name = appName
		}
		host.Context.HostConfiguration = config
	}
	return host
}

// ConfigureServices configure function by ServiceCollection of DI.
func (host *HostBuilder) ConfigureServices(configure func(*dependencyinjection.ServiceCollection)) *HostBuilder {
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

// RunningHostEnvironmentSetting ,get running hostenv setting.
func RunningHostEnvironmentSetting(hostEnv *HostEnvironment) {
	hostEnv.Host = getLocalIP()
	hostEnv.PID = os.Getpid()
}

//buildingHostEnvironmentSetting  build each configuration by init , such as file or hostenv or args ...
func buildingHostEnvironmentSetting(serviceCollection *dependencyinjection.ServiceCollection, context *HostBuilderContext) {
	hostEnv := context.HostingEnvironment
	hostEnv.Version = yoyogo.Version
	hostEnv.Addr = DetectAddress("")
	if context.Configuration != nil {
		hostEnv.MetaData = make(map[string]string)
		hostEnv.MetaData["config.path"] = context.Configuration.GetConfDir()
		hostEnv.MetaData["server.path"] = context.Configuration.GetString("yoyogo.application.server.path")
		mvc_template := context.Configuration.GetString("yoyogo.application.server.mvc.template")
		if mvc_template != "" {
			hostEnv.MetaData["mvc.template"] = mvc_template
		}
	}
	config := context.HostConfiguration
	if config != nil {
		hostEnv.ApplicationName = config.Name
		hostEnv.Server = config.Server.ServerType
		if config.Server.Address != "" {
			hostEnv.Addr = config.Server.Address
		}
	}

	hostEnv.Port = strings.Replace(hostEnv.Addr, ":", "", -1)
	hostEnv.Args = os.Args

	if hostEnv.Profile == "" {
		hostEnv.Profile = hostenv.Dev
	}
	httpserverConfig := hostenv.HttpServerConfig{}
	if config != nil {
		httpserverConfig := config.Server.Tls
		if httpserverConfig.CertFile != "" && httpserverConfig.KeyFile != "" {
			httpserverConfig.IsTLS = true
		}
	}
	httpserverConfig.Addr = hostEnv.Addr
	serviceCollection.AddSingleton(func() hostenv.HttpServerConfig { return httpserverConfig })
}

// Build host
func (host *HostBuilder) Build() IServiceHost {
	services := dependencyinjection.NewServiceCollection()

	buildingHostEnvironmentSetting(services, host.Context)
	host.Context.ApplicationCycle = NewApplicationLife()

	innerConfigures(host.Context, services)
	host.Decorator.OverrideIOCInnerConfigures(services)

	for _, configProcessorRegFunc := range configurationProcessors {
		configProcessorRegFunc(host.Context.Configuration, services)
	}

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
func innerConfigures(hostContext *HostBuilderContext, serviceCollection *dependencyinjection.ServiceCollection) {
	serviceCollection.AddSingleton(func() IConfiguration { return hostContext.Configuration })
	serviceCollection.AddSingleton(func() *ApplicationLife { return hostContext.ApplicationCycle })
	serviceCollection.AddSingleton(func() *HostEnvironment { return hostContext.HostingEnvironment })
}
