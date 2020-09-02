package LB

import (
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"sync/atomic"
)

type roundRobin struct {
	s ServiceDiscovery.IServiceDiscovery
	c uint64
}

// NewRandom returns a load balancer that selects services randomly.
func NewRound(sd ServiceDiscovery.IServiceDiscovery) ServiceDiscovery.Balancer {
	return &roundRobin{
		s: sd,
		c: 0,
	}
}

func (r *roundRobin) Next(serviceName string) (ServiceDiscovery.ServiceInstance, error) {
	endpoints := r.s.GetAllInstances(serviceName)
	if len(endpoints) <= 0 {
		return nil, ServiceDiscovery.ErrNoEndpoints
	}
	old := atomic.AddUint64(&r.c, 1) - 1
	idx := old % uint64(len(endpoints))
	return endpoints[idx], nil

}
