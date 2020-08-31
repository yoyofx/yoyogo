package Eureka

import (
	"errors"
	"fmt"
	"github.com/hudl/fargo"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/Internal/ServiceDiscoveryProvider"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type Option struct {
	ENV     *Context.HostEnvironment
	Address string `mapstructure:"address"`
}

type Registrar struct {
	cacheLocalInstance ServiceDiscovery.ServiceInstance
	logger             XLog.ILogger
	config             Option
	client             *Client
	eurekaConnection   *fargo.EurekaConnection
}

func NewServerDiscoveryWithDI(configuration Abstractions.IConfiguration, env *Context.HostEnvironment) ServiceDiscovery.IServiceDiscovery {
	sdType, ok := configuration.Get("yoyogo.cloud.discovery.type").(string)
	if !ok || sdType != "eureka" {
		panic(errors.New("yoyogo.cloud.discovery.type is not config node"))
	}
	section := configuration.GetSection("yoyogo.cloud.discovery.metadata")
	if section == nil {
		panic(errors.New("yoyogo.cloud.discovery.metadata is not config node"))
	}
	option := Option{}
	section.Unmarshal(&option)
	option.ENV = env
	return NewServerDiscovery(option)
}
func NewServerDiscovery(option Option) ServiceDiscovery.IServiceDiscovery {
	logger := XLog.GetXLogger("Server Discovery Eureka")
	eurekaRegister := &Registrar{}
	var fargoConfig fargo.Config
	fargoConfig.Eureka.ServiceUrls = []string{option.Address}
	// 订阅服务器应轮询更新的频率。
	fargoConfig.Eureka.PollIntervalSeconds = 1
	fargoConnection := fargo.NewConnFromConfig(fargoConfig)
	eurekaRegister.logger = logger
	eurekaRegister.eurekaConnection = &fargoConnection
	eurekaRegister.config = option
	return eurekaRegister
}

func (registrar *Registrar) GetName() string {
	return "eureka"
}

func (registrar *Registrar) Register() error {
	registrar.cacheLocalInstance = ServiceDiscoveryProvider.CreateServiceInstance(registrar.config.ENV)
	if registrar.client == nil {
		instance := &fargo.Instance{
			InstanceId:     registrar.cacheLocalInstance.GetId(),
			HostName:       registrar.cacheLocalInstance.GetHost(),
			Port:           int(registrar.cacheLocalInstance.GetPort()),
			App:            registrar.cacheLocalInstance.GetServiceName(),
			IPAddr:         registrar.cacheLocalInstance.GetHost(),
			HealthCheckUrl: fmt.Sprintf("http://%s:%d%s", registrar.cacheLocalInstance.GetHost(), registrar.cacheLocalInstance.GetPort(), "/actuator/health"),
			Status:         fargo.UP,
			DataCenterInfo: fargo.DataCenterInfo{Name: fargo.MyOwn},
			LeaseInfo:      fargo.LeaseInfo{RenewalIntervalInSecs: 1},
		}
		registrar.client = NewClient(registrar.eurekaConnection, instance)
	}
	registrar.client.Register()
	return nil
}

func (registrar *Registrar) Update() error {
	panic("implement me")
}

func (registrar *Registrar) Unregister() error {
	registrar.client.Deregister()
	return nil
}

func (registrar *Registrar) GetHealthyInstances(serviceName string) []ServiceDiscovery.ServiceInstance {
	panic("implement me")
}

func (registrar *Registrar) GetAllInstances(serviceName string) []ServiceDiscovery.ServiceInstance {
	app, err := registrar.eurekaConnection.GetApp(serviceName)
	if err != nil {
		return nil
	}

	var serviceList []ServiceDiscovery.ServiceInstance
	for _, service := range app.Instances {
		instance := &ServiceDiscovery.DefaultServiceInstance{
			Id:          service.InstanceId,
			ServiceName: service.App,
			Host:        service.IPAddr,
			Port:        uint64(service.Port),
			Enable:      true,
			Weight:      0,
			Healthy:     true,
		}
		serviceList = append(serviceList, instance)
	}
	return serviceList
}

func (registrar *Registrar) Destroy() error {
	return registrar.Unregister()
}
