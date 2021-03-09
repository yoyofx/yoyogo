package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	sd "github.com/yoyofx/yoyogo/pkg/servicediscovery"
	"strings"
)

type Registrar struct {
	cacheLocalInstance servicediscovery.ServiceInstance
	logger             xlog.ILogger
	config             Config
	client             naming_client.INamingClient
}

func NewServerDiscoveryWithDI(configuration abstractions.IConfiguration, env *abstractions.HostEnvironment) servicediscovery.IServiceDiscovery {
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

func NewServerDiscovery(option Config) servicediscovery.IServiceDiscovery {
	logger := xlog.GetXLogger("Server Discovery nacos")
	nacosRegister := &Registrar{}
	var serverConfigs []constant.ServerConfig
	urls := strings.Split(option.Url, ";")
	for _, url := range urls {
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			ContextPath: "/nacos",
			IpAddr:      url,
			Port:        option.Port,
		})
	}
	//serverConfigs = []constant.ServerConfig{
	//	{
	//		IpAddr:      option.Url,
	//		ContextPath: "/nacos",
	//		Port:        option.Port,
	//	},
	//}

	clientConfig := constant.ClientConfig{
		NamespaceId:         option.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "info",
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

func (registrar Registrar) GetName() string {
	return "nacos"
}

func (registrar *Registrar) Register() error {
	registrar.cacheLocalInstance = sd.CreateServiceInstance(registrar.config.ENV)
	success, err := registrar.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          registrar.cacheLocalInstance.GetHost(),
		Port:        registrar.cacheLocalInstance.GetPort(),
		ServiceName: registrar.cacheLocalInstance.GetServiceName(),
		Weight:      10,
		ClusterName: registrar.config.ClusterName,
		GroupName:   registrar.config.GroupName,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata: map[string]string{
			"VERSION": registrar.config.ENV.Version,
		},
	})
	if err != nil {
		registrar.logger.Error(err.Error())
	}
	registrar.logger.Debug("Registrar IP: %s , Success: %v", registrar.config.ENV.Host, success)
	return err
}

func (registrar *Registrar) Update() error {

	return nil
}

func (registrar *Registrar) Unregister() error {
	if registrar.cacheLocalInstance == nil {
		return nil
	}
	_, err := registrar.client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          registrar.cacheLocalInstance.GetHost(),
		Port:        registrar.cacheLocalInstance.GetPort(),
		Cluster:     registrar.cacheLocalInstance.GetClusterName(),
		ServiceName: registrar.cacheLocalInstance.GetServiceName(),
		GroupName:   registrar.cacheLocalInstance.GetGroupName(),
		Ephemeral:   true,
	})
	if err != nil {
		registrar.logger.Error(err.Error())
	}
	return err
}

func (registrar *Registrar) GetHealthyInstances(serviceName string) []servicediscovery.ServiceInstance {
	// SelectInstances only return the instances of healthy=${HealthyOnly},enable=true and weight>0
	instances, err := registrar.client.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   registrar.config.GroupName,             // default value is DEFAULT_GROUP
		Clusters:    []string{registrar.config.ClusterName}, // default value is DEFAULT
		HealthyOnly: true,
	})
	if err != nil {
		return nil
	}
	return convInstance(registrar.config.GroupName, instances)
}

func (registrar *Registrar) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	instances, err := registrar.client.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: serviceName,
		GroupName:   registrar.config.GroupName,             // default value is DEFAULT_GROUP
		Clusters:    []string{registrar.config.ClusterName}, // default value is DEFAULT
	})

	if err != nil {
		return nil
	}
	return convInstance(registrar.config.GroupName, instances)
}

func convInstance(groupName string, sourceInstances []model.Instance) []servicediscovery.ServiceInstance {
	var serviceList []servicediscovery.ServiceInstance
	for _, s := range sourceInstances {
		instance := &servicediscovery.DefaultServiceInstance{
			Id:          s.InstanceId,
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

func (registrar *Registrar) Destroy() error {
	registrar.logger.Debug("Destroy")
	err := registrar.Unregister()
	return err
}

func (registrar *Registrar) Watch(opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	return newWatcher(registrar.client, registrar.logger, opts...)
}

func (registrar *Registrar) GetAllServices() ([]*servicediscovery.Service, error) {
	serviceList, _ := registrar.client.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: registrar.config.NamespaceId,
		GroupName: registrar.config.GroupName,
		PageNo:    1,
		PageSize:  1000,
	})

	services := make([]*servicediscovery.Service, 0)
	for _, serviceName := range serviceList.Doms {
		services = append(services, &servicediscovery.Service{Name: serviceName})
	}
	return services, nil
}
