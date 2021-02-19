package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"reflect"
	"strconv"
	"sync"
)

type Watcher struct {
	serviceName    string
	namingClient   naming_client.INamingClient
	done           chan bool
	results        chan *servicediscovery.Result
	cacheLock      sync.Mutex
	instanceMap    map[string]model.Instance //Instance map of host for key
	subscribeParam *vo.SubscribeParam
	logger         xlog.ILogger
}

func newWatcher(client naming_client.INamingClient, log xlog.ILogger, opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	var wo servicediscovery.WatchOptions
	for _, o := range opts {
		o(&wo)
	}

	w := &Watcher{
		serviceName:  wo.Service,
		namingClient: client,
		logger:       log,
		done:         make(chan bool),
		results:      make(chan *servicediscovery.Result),
		instanceMap:  map[string]model.Instance{},
	}

	err := w.start()
	return w, err
}

func (w *Watcher) Next() (*servicediscovery.Result, error) {
	for {
		select {
		case <-w.done:
			w.logger.Warning("nacos listener is close!")
			return nil, errors.New("listener stopped")

		case e := <-w.results:
			w.logger.Debug("got nacos event %s", e)
			return e, nil
		}
	}
}

func (w *Watcher) Stop() {
	_ = w.namingClient.Unsubscribe(w.subscribeParam)
	close(w.done)
}

func (w *Watcher) start() error {
	if w.namingClient == nil {
		return errors.New("nacos naming client is nil")
	}
	w.subscribeParam = &vo.SubscribeParam{ServiceName: w.serviceName, SubscribeCallback: w.callback}
	go func() {
		_ = w.namingClient.Subscribe(w.subscribeParam)
	}()
	return nil
}

// Callback will be invoked when got subscribed events.
func (w *Watcher) callback(services []model.SubscribeService, err error) {
	if err != nil {
		w.logger.Error("nacos subscribe callback error:%s , subscribe:%+v ", err.Error(), w.subscribeParam)
		return
	}
	w.cacheLock.Lock()
	defer w.cacheLock.Unlock()

	addInstances := make([]model.Instance, 0, len(services))
	delInstances := make([]model.Instance, 0, len(services))
	updateInstances := make([]model.Instance, 0, len(services))
	newInstanceMap := make(map[string]model.Instance, len(services))

	//add and update cache
	for i := range services {
		if !services[i].Enable || !services[i].Valid {
			// instance is not available,so ignore it
			continue
		}
		host := services[i].Ip + ":" + strconv.Itoa(int(services[i].Port))
		instance := generateInstance(services[i])
		newInstanceMap[host] = instance
		if old, ok := w.instanceMap[host]; !ok {
			// instance does not exist in cache, add it to cache
			addInstances = append(addInstances, instance)
		} else {
			// instance is not different from cache, update it to cache
			if !reflect.DeepEqual(old, instance) {
				updateInstances = append(updateInstances, instance)
			}
		}
	}
	//del cache
	for host, inst := range w.instanceMap {
		if _, ok := newInstanceMap[host]; !ok {
			// cache instance does not exist in new instance list, remove it from cache
			delInstances = append(delInstances, inst)
		}
	}

	w.instanceMap = newInstanceMap

	for i := range addInstances {
		w.results <- &servicediscovery.Result{
			Action: "create", Service: &servicediscovery.Service{
				Name:  w.serviceName,
				Nodes: []servicediscovery.ServiceInstance{toServiceInstance(addInstances[i])},
			},
		}
	}
	for i := range delInstances {
		w.results <- &servicediscovery.Result{
			Action: "delete", Service: &servicediscovery.Service{
				Name:  w.serviceName,
				Nodes: []servicediscovery.ServiceInstance{toServiceInstance(delInstances[i])},
			},
		}
	}
	for i := range updateInstances {
		w.results <- &servicediscovery.Result{
			Action: "update", Service: &servicediscovery.Service{
				Name:  w.serviceName,
				Nodes: []servicediscovery.ServiceInstance{toServiceInstance(updateInstances[i])},
			},
		}
	}

}

func generateInstance(ss model.SubscribeService) model.Instance {
	return model.Instance{
		InstanceId:  ss.InstanceId,
		Ip:          ss.Ip,
		Port:        ss.Port,
		ServiceName: ss.ServiceName,
		Valid:       ss.Valid,
		Enable:      ss.Enable,
		Weight:      ss.Weight,
		Metadata:    ss.Metadata,
		ClusterName: ss.ClusterName,
	}
}

func toServiceInstance(s model.Instance) *servicediscovery.DefaultServiceInstance {
	instance := &servicediscovery.DefaultServiceInstance{
		Id:          s.InstanceId,
		ServiceName: s.ServiceName,
		Host:        s.Ip,
		Port:        s.Port,
		ClusterName: s.ClusterName,
		Enable:      true,
		Weight:      s.Weight,
		Healthy:     s.Healthy,
		Metadata:    s.Metadata,
	}
	return instance
}
