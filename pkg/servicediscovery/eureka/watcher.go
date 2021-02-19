package eureka

import (
	"errors"
	"github.com/hudl/fargo"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"strings"
	"time"
)

type Watcher struct {
	conn    fargoConnection
	exit    chan bool
	results chan *servicediscovery.Result
}

func newWatcher(conn fargoConnection, log xlog.ILogger, opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	var wo servicediscovery.WatchOptions
	for _, o := range opts {
		o(&wo)
	}

	w := &Watcher{
		conn:    conn,
		exit:    make(chan bool),
		results: make(chan *servicediscovery.Result),
	}

	if len(wo.Service) > 0 {
		done := make(chan struct{})
		ch := conn.ScheduleAppUpdates(wo.Service, false, done)
		go w.watch(ch, done)
		go func() {
			<-w.exit
			close(done)
		}()
		return w, nil
	}

	// watch all services
	go w.poll()
	return w, nil
}

func (w *Watcher) Next() (*servicediscovery.Result, error) {
	select {
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	case r := <-w.results:
		return r, nil
	}
}

func (w *Watcher) Stop() {
	close(w.exit)
}

func (e *Watcher) poll() {
	// list service ticker
	t := time.NewTicker(time.Second * 10)

	done := make(chan struct{})
	services := make(map[string]<-chan fargo.AppUpdate)

	for {
		select {
		case <-e.exit:
			close(done)
			return
		case <-t.C:
			apps, err := e.conn.GetApps()
			if err != nil {
				continue
			}
			for _, app := range apps {
				if _, ok := services[app.Name]; ok {
					continue
				}
				ch := e.conn.ScheduleAppUpdates(app.Name, false, done)
				services[app.Name] = ch
				go e.watch(ch, done)
			}
		}
	}
}

func (e *Watcher) watch(ch <-chan fargo.AppUpdate, done chan struct{}) {
	for {
		select {
		// exit on exit
		case <-e.exit:
			return
		// exit on done
		case <-done:
			return
		// process updates
		case u := <-ch:
			if u.Err != nil {
				continue
			}

			// process instances independently
			for _, instance := range u.App.Instances {
				var action string

				switch instance.Status {
				// update
				case fargo.UP:
					action = "update"
				// delete
				case fargo.OUTOFSERVICE, fargo.UNKNOWN, fargo.DOWN:
					action = "delete"
				// skip
				default:
					continue
				}

				// construct the service with a single node
				service := appToService(&fargo.Application{
					Name:      u.App.Name,
					Instances: []*fargo.Instance{instance},
				})

				if len(service) == 0 {
					continue
				}

				// in case we get bounced during processing
				// check exit channels
				select {
				// send the update
				case e.results <- &servicediscovery.Result{Action: action, Service: service[0]}:
				case <-done:
					return
				case <-e.exit:
					return
				}
			}
		}
	}
}

func appToService(app *fargo.Application) []*servicediscovery.Service {
	services := make([]*servicediscovery.Service, 0)
	service := &servicediscovery.Service{
		Name: strings.ToLower(app.Name),
	}
	for _, instance := range app.Instances {
		s := &servicediscovery.DefaultServiceInstance{
			Id:          instance.Id(),
			ServiceName: service.Name,
			Host:        instance.IPAddr,
			Port:        uint64(instance.Port),
			ClusterName: "",
			Enable:      true,
			Weight:      10,
			Healthy:     true,
			Metadata:    nil,
		}
		service.Nodes = append(service.Nodes, s)
	}

	services = append(services, service)

	return services
}
