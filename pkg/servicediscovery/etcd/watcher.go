package etcd

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"strconv"
	"strings"
	"time"
)

type Watcher struct {
	serviceName string
	namespace   string
	stop        chan bool
	watcher     clientv3.WatchChan
	client      *clientv3.Client
	timeout     time.Duration
}

func newWatcher(client *clientv3.Client, namespace string, opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	var wo servicediscovery.WatchOptions
	for _, o := range opts {
		o(&wo)
	}

	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan bool, 1)

	go func() {
		<-stop
		cancel()
	}()

	if wo.Service == "" {
		panic("service is not nil")
	}
	serviceName := wo.Service
	if !strings.Contains(wo.Service, namespace) {
		wo.Service = fmt.Sprintf("/%s/%s#", namespace, wo.Service)
	}
	watchPath := wo.Service
	watcher := client.Watch(ctx, watchPath, clientv3.WithPrefix(), clientv3.WithPrevKV())
	return &Watcher{
		serviceName: serviceName,
		namespace:   namespace,
		stop:        stop,
		watcher:     watcher,
		client:      client,
	}, nil
}

func (w *Watcher) Next() (*servicediscovery.Result, error) {
	for eps := range w.watcher {
		if eps.Err() != nil {
			return nil, eps.Err()
		}
		if eps.Canceled {
			return nil, errors.New("could not get next")
		}
		for _, ev := range eps.Events {
			var action string

			service := appToService(w.namespace, w.serviceName, string(ev.Kv.Key), string(ev.Kv.Value))
			switch ev.Type {
			case clientv3.EventTypePut:
				if ev.IsCreate() {
					action = "create"
				} else if ev.IsModify() {
					action = "update"
				}
			case clientv3.EventTypeDelete:
				action = "delete"
				// get service from prevKv
				service = appToService(w.namespace, w.serviceName, string(ev.PrevKv.Key), string(ev.PrevKv.Value))
			}

			if service == nil {
				continue
			}
			return &servicediscovery.Result{
				Action:  action,
				Service: service,
			}, nil
		}

	}
	return nil, errors.New("could not get next")
}

func (w *Watcher) Stop() {
	select {
	case <-w.stop:
		return
	default:
		close(w.stop)
	}
}

func appToService(namespace string, serviceName string, key string, value string) *servicediscovery.Service {
	//serviceNameEndIndex := strings.LastIndex(key, "#")
	//serviceName := utils.Substr(key, 0, serviceNameEndIndex)

	service := &servicediscovery.Service{
		Name: serviceName,
	}
	var ip string
	var port int
	//names := strings.Split(serviceName, "/")
	if value != "" {
		serviceAddr := value
		address := strings.Split(serviceAddr, ":")
		ip = address[0]
		port, _ = strconv.Atoi(address[1])
	}
	s := &servicediscovery.DefaultServiceInstance{
		Id:          key,
		ServiceName: service.Name,
		Host:        ip,
		Port:        uint64(port),
		ClusterName: namespace,
		Enable:      true,
		Weight:      10,
		Healthy:     true,
		Metadata:    nil,
	}
	service.Nodes = append(service.Nodes, s)

	return service
}
