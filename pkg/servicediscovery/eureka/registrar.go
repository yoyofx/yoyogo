package eureka

import (
	"errors"
	"fmt"
	"github.com/hudl/fargo"
	"github.com/op/go-logging"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	sd "github.com/yoyofx/yoyogo/pkg/servicediscovery"
	"github.com/yoyofx/yoyogo/utils"
	"strings"
)

type Option struct {
	ENV        *abstractions.HostEnvironment
	Address    string `mapstructure:"address"`
	Ttl        int    `mapstructure:"ttl"`
	DataCenter string `mapstructure:"datacenter"`
}

type Registrar struct {
	cacheLocalInstance servicediscovery.ServiceInstance
	logger             xlog.ILogger
	config             *Option
	client             *Client
	eurekaConnection   *fargo.EurekaConnection
}

func NewServerDiscoveryWithDI(configuration abstractions.IConfiguration, env *abstractions.HostEnvironment) servicediscovery.IServiceDiscovery {
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
	return NewServerDiscovery(&option)
}
func NewServerDiscovery(option *Option) servicediscovery.IServiceDiscovery {
	logger := xlog.GetXLogger("Server Discovery eureka")
	if option.DataCenter == "" {
		option.DataCenter = fargo.MyOwn
	}
	eurekaRegister := &Registrar{}
	var fargoConfig fargo.Config
	fargoConfig.Eureka.ServiceUrls = strings.Split(option.Address, ";")
	// 订阅服务器应轮询更新的频率。
	fargoConfig.Eureka.PollIntervalSeconds = option.Ttl
	fargoConnection := fargo.NewConnFromConfig(fargoConfig)
	logging.SetLevel(logging.WARNING, "fargo")
	eurekaRegister.logger = logger
	eurekaRegister.eurekaConnection = &fargoConnection
	eurekaRegister.config = option
	return eurekaRegister
}

func (registrar *Registrar) getId() string {
	return utils.Md5String(fmt.Sprintf("%s#%s#%s:%v", registrar.config.DataCenter, registrar.cacheLocalInstance.GetServiceName(), registrar.cacheLocalInstance.GetHost(), registrar.cacheLocalInstance.GetPort()))
}

func (registrar *Registrar) GetName() string {
	return "eureka"
}

func (registrar *Registrar) Register() error {
	registrar.cacheLocalInstance = sd.CreateServiceInstance(registrar.config.ENV)
	if registrar.client == nil {
		instance := &fargo.Instance{
			InstanceId:     registrar.getId(),
			HostName:       registrar.cacheLocalInstance.GetHost(),
			Port:           int(registrar.cacheLocalInstance.GetPort()),
			App:            registrar.cacheLocalInstance.GetServiceName(),
			IPAddr:         registrar.cacheLocalInstance.GetHost(),
			HealthCheckUrl: fmt.Sprintf("http://%s:%d%s", registrar.cacheLocalInstance.GetHost(), registrar.cacheLocalInstance.GetPort(), "/actuator/health"),
			Status:         fargo.UP,
			DataCenterInfo: fargo.DataCenterInfo{Name: registrar.config.DataCenter},
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

func (registrar *Registrar) GetHealthyInstances(serviceName string) []servicediscovery.ServiceInstance {
	return registrar.GetAllInstances(serviceName)
}

func (registrar *Registrar) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	app, err := registrar.eurekaConnection.GetApp(serviceName)
	//registrar.eurekaConnection.UpdateApp()
	if err != nil {
		return nil
	}

	var serviceList []servicediscovery.ServiceInstance
	for _, service := range app.Instances {
		instance := &servicediscovery.DefaultServiceInstance{
			Id:          service.InstanceId,
			ServiceName: service.App,
			Host:        service.IPAddr,
			Port:        uint64(service.Port),
			ClusterName: registrar.config.DataCenter,
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

func (registrar *Registrar) Watch(opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	return newWatcher(registrar.eurekaConnection, registrar.logger, opts...)
}

func (registrar *Registrar) GetAllServices() ([]*servicediscovery.Service, error) {
	apps, err := registrar.eurekaConnection.GetApps()
	services := make([]*servicediscovery.Service, 0)
	for _, app := range apps {
		services = append(services, &servicediscovery.Service{Name: app.Name})
	}
	return services, err
}
