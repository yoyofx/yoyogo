package etcd

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	sd "github.com/yoyofx/yoyogo/pkg/servicediscovery"
	"github.com/yoyofx/yoyogo/utils"
	"strconv"
	"strings"
	"time"
)

type Registrar struct {
	cacheLocalInstance servicediscovery.ServiceInstance
	logger             xlog.ILogger
	config             Config
	client             *clientv3.Client
	servicePath        string
}

func NewServerDiscoveryWithDI(configuration abstractions.IConfiguration, env *abstractions.HostEnvironment) servicediscovery.IServiceDiscovery {
	sdType, ok := configuration.Get("yoyogo.cloud.discovery.type").(string)
	if !ok || sdType != "etcd" {
		panic(errors.New("yoyogo.cloud.discovery.type is not config node"))
	}
	section := configuration.GetSection("yoyogo.cloud.discovery.metadata")
	if section == nil {
		panic(errors.New("yoyogo.cloud.discovery.metadata is not config node"))
	}
	option := Config{}
	section.Unmarshal(&option)
	option.ENV = env
	return NewServerDiscovery(option)
}

func NewServerDiscovery(option Config) servicediscovery.IServiceDiscovery {
	logger := xlog.GetXLogger("Service Discovery ETCD")

	v3config := clientv3.Config{
		Endpoints:   option.Address,
		DialTimeout: 15 * time.Second,
	}
	if option.Auth != nil && option.Auth.Enable {
		v3config.Username = option.Auth.User
		v3config.Password = option.Auth.Password
	}

	cli, err := clientv3.New(v3config)
	if err != nil {
		logger.Error("connect to etcd err:%s", err)
		return nil
	}
	ETCDRegister := &Registrar{
		logger: logger,
		config: option,
		client: cli,
	}
	return ETCDRegister
}

func (r *Registrar) GetName() string {
	return "ETCD"
}

func nodePath(namespace string, serviceName string, nodeId string) string {
	return fmt.Sprintf("/%s/%s#%s", namespace,
		serviceName, nodeId)
}

func (r *Registrar) Register() error {
	if r.config.Ttl < 5 {
		panic("etcd ttl value must be bigger than 5")
	}
	r.cacheLocalInstance = sd.CreateServiceInstance(r.config.ENV)
	if r.client != nil {
		ticker := time.NewTicker(time.Second * time.Duration(r.config.Ttl-1))
		go func() {
			for {
				addr := fmt.Sprintf("%s:%v", r.cacheLocalInstance.GetHost(), r.cacheLocalInstance.GetPort())
				r.servicePath = nodePath(r.config.Namespace, r.cacheLocalInstance.GetServiceName(), r.cacheLocalInstance.GetId())
				getResp, err := r.client.Get(context.Background(), r.servicePath)
				if err != nil {
					r.logger.Error("Register:%s", err)
				} else if getResp.Count == 0 {
					err = withAlive(r.client, r.servicePath, addr, r.config.Ttl)
					if err != nil {
						r.logger.Error("keep alive:%s", err)
					} else {
						r.logger.Debug("ETCD Register: %s", r.servicePath)
					}
				}
				<-ticker.C
			}
		}()
	}

	return nil

}

func withAlive(cli *clientv3.Client, servicePath string, addr string, ttl int64) error {
	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}
	_, err = cli.Put(context.Background(), servicePath, addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		fmt.Printf("put etcd error:%s", err)
		return err
	}

	_, err = cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		fmt.Printf("keep alive error:%s", err)
		return err
	}
	return nil
}

func (r *Registrar) Update() error {
	panic("implement me")
}

func (r *Registrar) Unregister() error {
	if r.client != nil {
		_, err := r.client.Delete(context.Background(), r.servicePath)

		if err == nil {
			r.logger.Debug("ETCD UnRegister Succeeded: %s", r.servicePath)
		}

		return err
	}
	return nil
}

func (r *Registrar) GetHealthyInstances(serviceName string) []servicediscovery.ServiceInstance {
	return r.GetAllInstances(serviceName)
}

func (r *Registrar) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	serviceRoot := serviceName
	if !strings.Contains(serviceName, r.config.Namespace) {
		serviceRoot = fmt.Sprintf("/%s/%s", r.config.Namespace, serviceName)
	}
	var serviceList []servicediscovery.ServiceInstance
	getResp, err := r.client.Get(context.Background(), serviceRoot, clientv3.WithPrefix())
	if err != nil {
		r.logger.Error("etcd get path error:%s ", err)
	} else {
		for i := range getResp.Kvs {
			serviceID := string(getResp.Kvs[i].Key)
			serviceAddr := string(getResp.Kvs[i].Value)
			address := strings.Split(serviceAddr, ":")
			ip := address[0]
			port, _ := strconv.Atoi(address[1])
			instance := &servicediscovery.DefaultServiceInstance{
				Id:          serviceID,
				ServiceName: serviceName,
				Host:        ip,
				Port:        uint64(port),
				Enable:      true,
				Weight:      0,
				Healthy:     true,
				ClusterName: r.config.Namespace,
			}
			serviceList = append(serviceList, instance)
		}
	}
	return serviceList
}

func (r *Registrar) Destroy() error {
	return r.Unregister()
}

func (r *Registrar) Watch(opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	return newWatcher(r.client, r.config.Namespace, opts...)
}

func (r *Registrar) GetAllServices() ([]*servicediscovery.Service, error) {
	services := make([]*servicediscovery.Service, 0)
	serviceRoot := fmt.Sprintf("/%s", r.config.Namespace)
	getResp, _ := r.client.Get(context.Background(), serviceRoot, clientv3.WithPrefix(), clientv3.WithSerializable())
	servicesMap := make(map[string]string)
	nslen := len(serviceRoot) + 1
	for i := range getResp.Kvs {
		serviceName := string(getResp.Kvs[i].Key)
		slastIdx := strings.Index(serviceName, "#")
		serviceName = utils.Substr(serviceName, 0, slastIdx)
		serviceName = utils.Substr(serviceName, nslen, len(serviceName))
		_, ok := servicesMap[serviceName]
		if !ok {
			servicesMap[serviceName] = serviceName
			services = append(services, &servicediscovery.Service{Name: serviceName})
		}

	}
	return services, nil
}

func exist(l []string, addr string) bool {
	for i := range l {
		if l[i] == addr {
			return true
		}
	}
	return false
}

func remove(s []string, addr string) ([]string, bool) {
	for i := range s {
		if s[i] == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
