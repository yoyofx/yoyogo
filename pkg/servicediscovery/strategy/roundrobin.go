package strategy

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"sync/atomic"
)

type roundRobin struct {
	s servicediscovery.IServiceDiscovery
	c uint64
}

// NewRandom returns a load balancer that selects services randomly.
func NewRound(sd servicediscovery.IServiceDiscovery) servicediscovery.Strategy {
	return &roundRobin{
		s: sd,
		c: 0,
	}
}

func (r *roundRobin) Next(serviceName string) (servicediscovery.ServiceInstance, error) {
	endpoints := r.s.GetAllInstances(serviceName)
	if len(endpoints) <= 0 {
		return nil, servicediscovery.ErrNoEndpoints
	}
	old := atomic.AddUint64(&r.c, 1) - 1
	idx := old % uint64(len(endpoints))
	return endpoints[idx], nil

}
