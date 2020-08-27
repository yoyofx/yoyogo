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
	"github.com/yoyofx/yoyogo/Abstractions/ServerDiscovery"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"strconv"
)

type Register struct {
	cacheLocalInstance ServerDiscovery.DefaultServiceInstance
	logger             XLog.ILogger
	config             Config
	client             naming_client.INamingClient
}

func NewServerDiscoveryWithDI(configuration Abstractions.IConfiguration, env *Context.HostEnvironment) ServerDiscovery.IServerDiscovery {
	section := configuration.GetSection("server_discovery.metadata")
	if section == nil {
		panic(errors.New("server_discovery.metadata is not config node"))
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

func NewServerDiscovery(option Config) ServerDiscovery.IServerDiscovery {
	logger := XLog.GetXLogger("Server Discovery Nacos")
	nacosRegister := &Register{}

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
		LogLevel:            "debug",
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

func (register Register) GetName() string {
	return "Nacos"
}

func (register *Register) Register() error {
	port, _ := strconv.ParseInt(register.config.ENV.Port, 10, 64)

	register.cacheLocalInstance = ServerDiscovery.DefaultServiceInstance{
		Id:          uuid.New().String(),
		ServiceName: register.config.ENV.ApplicationName,
		Host:        register.config.ENV.Host,
		Port:        uint64(port),
		ClusterName: register.config.ClusterName,
		GroupName:   register.config.GroupName,
		Enable:      true,
		Healthy:     true,
		Metadata: map[string]string{
			"VERSION": register.config.ENV.Version,
		},
	}

	success, err := register.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          register.cacheLocalInstance.Host,
		Port:        uint64(register.cacheLocalInstance.Port),
		ServiceName: register.cacheLocalInstance.ServiceName,
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
	register.logger.Debug("Register IP: %s , Success: %v", register.config.ENV.Host, success)
	return err
}

func (register Register) Update() error {

	return nil
}

func (register Register) Unregister() error {
	_, err := register.client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          register.cacheLocalInstance.Host,
		Port:        register.cacheLocalInstance.Port,
		Cluster:     register.cacheLocalInstance.ClusterName,
		ServiceName: register.cacheLocalInstance.ServiceName,
		GroupName:   register.cacheLocalInstance.GroupName,
		Ephemeral:   true,
	})
	if err != nil {
		register.logger.Error(err.Error())
	}
	return err
}

func (register Register) GetHealthyInstances(serviceName string) []ServerDiscovery.ServiceInstance {
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

func (register Register) GetAllInstances(serviceName string) []ServerDiscovery.ServiceInstance {
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

func convInstance(groupName string, sourceInstances []model.Instance) []ServerDiscovery.ServiceInstance {
	var serviceList []ServerDiscovery.ServiceInstance
	for _, s := range sourceInstances {
		instance := &ServerDiscovery.DefaultServiceInstance{
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

func (register Register) Destroy() error {
	register.logger.Debug("Destroy")
	err := register.Unregister()
	return err
}
