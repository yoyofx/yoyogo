package strategy

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"sync/atomic"
)

type RoundRobin struct {
	//s servicediscovery.IServiceDiscovery
	c uint64
}

// NewRandom returns a load balancer that selects services randomly.
func NewRound() servicediscovery.Strategy {
	return &RoundRobin{
		//s: sd,
		c: 0,
	}
}

func (r *RoundRobin) Next(instanceList []servicediscovery.ServiceInstance) (servicediscovery.ServiceInstance, error) {
	//endpoints := r.s.GetAllInstances(serviceName)
	if len(instanceList) <= 0 {
		return nil, servicediscovery.ErrNoEndpoints
	}
	old := atomic.AddUint64(&r.c, 1) - 1
	idx := old % uint64(len(instanceList))
	return instanceList[idx], nil

}
