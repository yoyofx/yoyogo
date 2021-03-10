package servicediscovery

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	DefaultTTL = 5 * time.Minute
)

type Cache interface {
	GetService(serviceName string) (*Service, error)
	// stop the cache of watcher
	Stop()
}

type CacheOptions struct {
	// TTL is the cache TTL
	TTL time.Duration
}
type Option func(o *CacheOptions)

func NewCache(r IServiceDiscoveryClient) Cache {
	rand.Seed(time.Now().UnixNano())
	options := CacheOptions{
		TTL: DefaultTTL,
	}

	//for _, o := range opts {
	//	o(&options)
	//}

	return &cache{
		discoveryClient: r,
		opts:            options,
		watched:         make(map[string]bool),
		cache:           make(map[string][]*Service),
		ttls:            make(map[string]time.Time),
		exit:            make(chan bool),
		log:             xlog.GetXLogger("servicediscovery.cache"),
	}
}

type cache struct {
	discoveryClient IServiceDiscoveryClient
	opts            CacheOptions
	// registry cache
	sync.RWMutex
	cache   map[string][]*Service
	ttls    map[string]time.Time
	watched map[string]bool
	// used to stop the cache
	exit    chan bool
	running bool
	status  error
	log     xlog.ILogger
}

func backoff(attempts int) time.Duration {
	if attempts == 0 {
		return time.Duration(0)
	}
	return time.Duration(math.Pow(10, float64(attempts))) * time.Millisecond
}

func (c *cache) GetService(serviceName string) (*Service, error) {
	// get the service
	services, err := c.get(serviceName)
	if err != nil {
		return nil, err
	}

	// if there's nothing return err
	if len(services) == 0 {
		return nil, errors.New("ErrNotFound")
	}

	// return services
	return services[0], nil
}

func (c *cache) Stop() {
	c.log.Debug("cache stopped!")
	c.Lock()
	defer c.Unlock()

	select {
	case <-c.exit:
		return
	default:
		close(c.exit)
	}
}

func (c *cache) getStatus() error {
	c.RLock()
	defer c.RUnlock()
	return c.status
}

func (c *cache) setStatus(err error) {
	c.Lock()
	c.status = err
	c.Unlock()
}

func (c *cache) isValid(services []*Service, ttl time.Time) bool {
	// no services exist
	if len(services) == 0 {
		return false
	}

	// ttl is invalid
	if ttl.IsZero() {
		return false
	}

	// time since ttl is longer than timeout
	if time.Since(ttl) > 0 {
		return false
	}

	// ok
	return true
}

func (c *cache) quit() bool {
	select {
	case <-c.exit:
		return true
	default:
		return false
	}
}

func (c *cache) del(service string) {
	// don't blow away cache in error state
	if err := c.status; err != nil {
		return
	}
	// otherwise delete entries
	delete(c.cache, service)
	delete(c.ttls, service)
}

func (c *cache) set(service string, services []*Service) {
	c.cache[service] = services
	c.ttls[service] = time.Now().Add(c.opts.TTL)
}

func (c *cache) get(service string) ([]*Service, error) {
	// read lock
	c.RLock()
	// check the cache first
	services := c.cache[service]
	// get cache ttl
	ttl := c.ttls[service]
	cp := Copy(services)

	if c.isValid(cp, ttl) {
		c.log.Debug("get by service with cache: %s", service)
		c.RUnlock()
		// return services
		return cp, nil
	}
	c.log.Debug("get by service without cache: %s", service)
	get := func(service string, cached []*Service) ([]*Service, error) {
		val, err := c.discoveryClient.GetService(service)
		services := []*Service{val}
		if err != nil {
			// check the cache
			if len(cached) > 0 {
				// set the error status
				c.setStatus(err)
				// return the stale cache
				return cached, nil
			}
		}
		// reset the status
		if err := c.getStatus(); err != nil {
			c.setStatus(nil)
		}
		// cache results
		c.Lock()
		c.set(service, Copy(services))
		c.Unlock()

		return services, nil
	}

	// watch service if not watched
	_, ok := c.watched[service]

	// unlock the read lock
	c.RUnlock()

	// check if its being watched
	if !ok {
		c.Lock()

		// set to watched
		c.watched[service] = true

		// only kick it off if not running
		if !c.running {
			go c.run(service)
		}

		c.Unlock()
	}

	// get and return services
	return get(service, cp)
}

func (c *cache) run(serviceName string) {
	c.Lock()
	c.running = true
	c.Unlock()
	// reset watcher on exit
	defer func() {
		c.Lock()
		c.watched = make(map[string]bool)
		c.running = false
		c.Unlock()
	}()
	var a, b int

	for {
		// exit early if already dead
		if c.quit() {
			return
		}

		// jitter before starting
		j := rand.Int63n(100)
		time.Sleep(time.Duration(j) * time.Millisecond)

		// create new watcher
		w, err := c.discoveryClient.Watch(func(options *WatchOptions) {
			options.Service = serviceName
		})
		if err != nil {
			if c.quit() {
				return
			}

			d := backoff(a)
			c.setStatus(err)

			if a > 3 {
				logger.Debug("rcache: ", err, " backing off ", d)
				a = 0
			}

			time.Sleep(d)
			a++

			continue
		}
		// reset a
		a = 0

		// watch for events
		if err := c.watch(w); err != nil {
			if c.quit() {
				return
			}

			d := backoff(b)
			c.setStatus(err)

			if b > 3 {
				logger.Debug("rcache: ", err, " backing off ", d)
				b = 0
			}

			time.Sleep(d)
			b++

			continue
		}

		// reset b
		b = 0
	}
}

func (c *cache) watch(w Watcher) error {
	// used to stop the watch
	stop := make(chan bool)

	// manage this loop
	go func() {
		defer w.Stop()

		select {
		// wait for exit
		case <-c.exit:
			return
		// we've been stopped
		case <-stop:
			return
		}
	}()

	for {
		res, err := w.Next()
		if err != nil {
			close(stop)
			return err
		}

		// reset the error status since we succeeded
		if err := c.getStatus(); err != nil {
			// reset status
			c.setStatus(nil)
		}

		c.update(res)
	}
}

func (c *cache) update(res *Result) {
	if res == nil || res.Service == nil {
		return
	}

	c.Lock()
	defer c.Unlock()

	// only save watched services
	if _, ok := c.watched[res.Service.Name]; !ok {
		return
	}

	services, ok := c.cache[res.Service.Name]
	if !ok {
		// we're not going to cache anything
		// unless there was already a lookup
		return
	}

	if len(res.Service.Nodes) == 0 {
		switch res.Action {
		case "delete":
			c.del(res.Service.Name)
		}
		return
	}

	// existing service found
	var service *Service
	var index int
	for i, s := range services {
		if s.Version == res.Service.Version {
			service = s
			index = i
		}
	}

	switch res.Action {
	case "create", "update":
		if service == nil {
			c.set(res.Service.Name, append(services, res.Service))
			return
		}

		// append old nodes to new service
		for _, cur := range service.Nodes {
			var seen bool
			for _, node := range res.Service.Nodes {
				if cur.GetId() == node.GetId() {
					seen = true
					break
				}
			}
			if !seen {
				res.Service.Nodes = append(res.Service.Nodes, cur)
			}
		}

		services[index] = res.Service
		c.set(res.Service.Name, services)
	case "delete":
		if service == nil {
			return
		}

		var nodes []ServiceInstance

		// filter cur nodes to remove the dead one
		for _, cur := range service.Nodes {
			var seen bool
			for _, del := range res.Service.Nodes {
				if del.GetId() == cur.GetId() {
					seen = true
					break
				}
			}
			if !seen {
				nodes = append(nodes, cur)
			}
		}

		// still got nodes, save and return
		if len(nodes) > 0 {
			service.Nodes = nodes
			services[index] = service
			c.set(service.Name, services)
			return
		}

		// zero nodes left

		// only have one thing to delete
		// nuke the thing
		if len(services) == 1 {
			c.del(service.Name)
			return
		}

		// still have more than 1 service
		// check the version and keep what we know
		var srvs []*Service
		for _, s := range services {
			if s.Version != service.Version {
				srvs = append(srvs, s)
			}
		}

		// save
		c.set(service.Name, srvs)
	case "override":
		if service == nil {
			return
		}

		c.del(service.Name)
	}
}
