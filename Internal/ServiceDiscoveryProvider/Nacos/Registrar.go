package Nacos

import (
	"errors"
	"github.com/google/uuid"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/Internal/ServiceDiscoveryProvider"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type Registrar struct {
	cacheLocalInstance ServiceDiscovery.ServiceInstance
	logger             XLog.ILogger
	config             Config
	client             naming_client.INamingClient
}

func NewServerDiscoveryWithDI(configuration Abstractions.IConfiguration, env *Context.HostEnvironment) ServiceDiscovery.IServiceDiscovery {
	sdType, ok := configuration.Get("yoyogo.cloud.discovery.type").(string)
	if !ok || sdType != "nacos" {
		panic(errors.New("yoyogo.cloud.discovery.type is not config node"))
	}
	section := configuration.GetSection("yoyogo.cloud.discovery.metadata")
	if section == nil {
		panic(errors.New("yoyogo.cloud.discovery.metadata is not config node"))
	}
	option := Config{}
	section.Unmarshal(&option)
	if option.GroupName == "" {
		option.GroupName = GroupName
	}
	if option.ClusterName == "" {
		option.ClusterName = Cluster
	}
	option.ENV = env

	return NewServerDiscovery(option)
}

func NewServerDiscovery(option Config) ServiceDiscovery.IServiceDiscovery {
	logger := XLog.GetXLogger("Server Discovery Nacos")
	nacosRegister := &Registrar{}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      option.Url,
			ContextPath: "/nacos",
			Port:        option.Port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         option.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "error",
	}

	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	nacosRegister.client = namingClient
	nacosRegister.config = option
	nacosRegister.logger = logger

	logger.Debug("url:%s, namespace:%s", option.Url, option.NamespaceId)
	return nacosRegister
}

func (register Registrar) GetName() string {
	return "Nacos"
}

func (register *Registrar) Register() error {
	register.cacheLocalInstance = ServiceDiscoveryProvider.CreateServiceInstance(register.config.ENV)
	success, err := register.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          register.cacheLocalInstance.GetHost(),
		Port:        register.cacheLocalInstance.GetPort(),
		ServiceName: register.cacheLocalInstance.GetServiceName(),
		Weight:      10,
		ClusterName: register.config.ClusterName,
		GroupName:   register.config.GroupName,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata: map[string]string{
			"VERSION": register.config.ENV.Version,
		},
	})
	if err != nil {
		register.logger.Error(err.Error())
	}
	register.logger.Debug("Registrar IP: %s , Success: %v", register.config.ENV.Host, success)
	return err
}

func (register Registrar) Update() error {

	return nil
}

func (register Registrar) Unregister() error {
	if register.cacheLocalInstance == nil {
		return nil
	}
	_, err := register.client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          register.cacheLocalInstance.GetHost(),
		Port:        register.cacheLocalInstance.GetPort(),
		Cluster:     register.cacheLocalInstance.GetClusterName(),
		ServiceName: register.cacheLocalInstance.GetServiceName(),
		GroupName:   register.cacheLocalInstance.GetGroupName(),
		Ephemeral:   true,
	})
	if err != nil {
		register.logger.Error(err.Error())
	}
	return err
}

func (register Registrar) GetHealthyInstances(serviceName string) []ServiceDiscovery.ServiceInstance {
	// SelectInstances only return the instances of healthy=${HealthyOnly},enable=true and weight>0
	instances, err := register.client.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   register.config.GroupName,             // default value is DEFAULT_GROUP
		Clusters:    []string{register.config.ClusterName}, // default value is DEFAULT
		HealthyOnly: true,
	})
	if err != nil {
		return nil
	}
	return convInstance(register.config.GroupName, instances)
}

func (register Registrar) GetAllInstances(serviceName string) []ServiceDiscovery.ServiceInstance {
	instances, err := register.client.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: serviceName,
		GroupName:   register.config.GroupName,             // default value is DEFAULT_GROUP
		Clusters:    []string{register.config.ClusterName}, // default value is DEFAULT
	})

	if err != nil {
		return nil
	}
	return convInstance(register.config.GroupName, instances)
}

func convInstance(groupName string, sourceInstances []model.Instance) []ServiceDiscovery.ServiceInstance {
	var serviceList []ServiceDiscovery.ServiceInstance
	for _, s := range sourceInstances {
		instance := &ServiceDiscovery.DefaultServiceInstance{
			Id:          uuid.New().String(),
			ServiceName: s.ServiceName,
			Host:        s.Ip,
			Port:        s.Port,
			ClusterName: s.ClusterName,
			GroupName:   groupName,
			Enable:      true,
			Weight:      s.Weight,
			Healthy:     s.Healthy,
			Metadata:    s.Metadata,
		}
		serviceList = append(serviceList, instance)
	}
	return serviceList
}

func (register Registrar) Destroy() error {
	register.logger.Debug("Destroy")
	err := register.Unregister()
	return err
}
