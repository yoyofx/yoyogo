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
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(option.Address, ";"),
		DialTimeout: 15 * time.Second,
	})
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

func (r *Registrar) Register() error {
	r.cacheLocalInstance = sd.CreateServiceInstance(r.config.ENV)
	if r.client != nil {
		ticker := time.NewTicker(time.Second * time.Duration(r.config.Ttl))
		go func() {
			for {
				addr := fmt.Sprintf("%s:%v", r.cacheLocalInstance.GetHost(), r.cacheLocalInstance.GetPort())
				r.servicePath = fmt.Sprintf("/%s/%s/%s", r.config.Namespace,
					r.cacheLocalInstance.GetServiceName(), r.cacheLocalInstance.GetId())
				getResp, err := r.client.Get(context.Background(), r.servicePath)
				if err != nil {
					r.logger.Error("Register:%s", err)
				} else if getResp.Count == 0 {
					err = withAlive(r.client, r.servicePath, addr, r.config.Ttl)
					if err != nil {
						r.logger.Error("keep alive:%s", err)
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
		return err
	}
	return nil
}

func (r *Registrar) GetHealthyInstances(serviceName string) []servicediscovery.ServiceInstance {
	return r.GetAllInstances(serviceName)
}

func (r *Registrar) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	var serviceList []servicediscovery.ServiceInstance
	serviceRoot := fmt.Sprintf("/%s/%s", r.config.Namespace, serviceName)
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
	panic("implement me")
	//getResp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	for i := range getResp.Kvs {
	//		addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(getResp.Kvs[i].Key), keyPrefix)})
	//	}
	//}
	//
	//r.cc.NewAddress(addrList)
	//
	//rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	//for n := range rch {
	//	for _, ev := range n.Events {
	//		addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
	//		switch ev.Type {
	//		case mvccpb.PUT:
	//			if !exist(addrList, addr) {
	//				addrList = append(addrList, resolver.Address{Addr: addr})
	//				r.cc.NewAddress(addrList)
	//			}
	//		case mvccpb.DELETE:
	//			if s, ok := remove(addrList, addr); ok {
	//				addrList = s
	//				r.cc.NewAddress(addrList)
	//			}
	//		}
	//		//log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	//	}
	//}
}

func (r *Registrar) GetAllServices() ([]*servicediscovery.Service, error) {
	services := make([]*servicediscovery.Service, 0)
	getResp, _ := r.client.Get(context.Background(), r.config.Namespace, clientv3.WithPrefix())

	for i := range getResp.Kvs {
		serviceName := string(getResp.Kvs[i].Key)
		services = append(services, &servicediscovery.Service{Name: serviceName})
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
