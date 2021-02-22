package consul

import (
	"errors"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	sd "github.com/yoyofx/yoyogo/pkg/servicediscovery"
)

type Registrar struct {
	cacheLocalInstance servicediscovery.ServiceInstance
	logger             xlog.ILogger
	client             *Client
	config             Option
}

func NewServerDiscoveryWithDI(configuration abstractions.IConfiguration, env *abstractions.HostEnvironment) servicediscovery.IServiceDiscovery {
	sdType, ok := configuration.Get("yoyogo.cloud.discovery.type").(string)
	if !ok || sdType != "consul" {
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

func NewServerDiscovery(option Option) servicediscovery.IServiceDiscovery {
	logger := xlog.GetXLogger("Server Discovery consul")
	client := NewClient(option)
	if client == nil {
		logger.Error("consul client is nil !!")
	}
	logger.Debug("url:%s", option.Address)
	return &Registrar{
		logger: logger,
		client: client,
		config: option,
	}
}

func (registrar *Registrar) Register() error {
	registrar.cacheLocalInstance = sd.CreateServiceInstance(registrar.config.ENV)

	registration := new(consul.AgentServiceRegistration)
	registration.ID = registrar.cacheLocalInstance.GetId()
	registration.Name = registrar.cacheLocalInstance.GetServiceName()
	registration.Port = int(registrar.cacheLocalInstance.GetPort())
	registration.Tags = registrar.config.Tags
	registration.Address = registrar.cacheLocalInstance.GetHost()
	registration.Tags = registrar.config.Tags

	registration.Check = &consul.AgentServiceCheck{ // 健康检查
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/actuator/health"),
		Timeout:                        "3s",
		Interval:                       "5s",  // 健康检查间隔
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务，注销时间，相当于过期时间
		// GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	err := registrar.client.Register(registration)
	registrar.logger.Debug("Registrar IP: %s , Success: %v", registrar.config.ENV.Host, err == nil)
	return err
}

func (registrar *Registrar) Update() error {
	panic("implement me")
}

func (registrar *Registrar) Unregister() error {
	if registrar.cacheLocalInstance == nil {
		return nil
	}
	registration := new(consul.AgentServiceRegistration)
	registration.ID = registrar.cacheLocalInstance.GetId()
	registrar.logger.Debug("unregister id: %s , success", registration.ID)
	return registrar.client.Deregister(registration)
}

func (registrar *Registrar) GetHealthyInstances(serviceName string) []servicediscovery.ServiceInstance {
	return registrar.GetAllInstances(serviceName)
}

func (registrar *Registrar) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	tag := ""
	if registrar.config.Tags != nil && len(registrar.config.Tags) > 0 {
		tag = registrar.config.Tags[0]
	}
	services, _, err := registrar.client.GetService(serviceName, tag, true, &consul.QueryOptions{})

	if err != nil {
		registrar.logger.Error("error retrieving instances from consul: %s", err.Error())
	}
	var serviceList []servicediscovery.ServiceInstance
	for _, service := range services {
		instance := &servicediscovery.DefaultServiceInstance{
			Id:          service.Service.ID,
			ServiceName: service.Service.Service,
			Host:        service.Service.Address,
			Port:        uint64(service.Service.Port),
			Tags:        service.Service.Tags,
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

func (registrar *Registrar) GetName() string {
	return "consul"
}

func (registrar *Registrar) Watch(opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	return newWatcher(registrar.client.consul, opts...)
}

func (registrar *Registrar) GetAllServices() ([]*servicediscovery.Service, error) {
	serviceNames := registrar.client.GetServices(&consul.QueryOptions{})
	services := make([]*servicediscovery.Service, 0)
	for _, serviceName := range serviceNames {
		services = append(services, &servicediscovery.Service{Name: serviceName})
	}
	return services, nil
}
