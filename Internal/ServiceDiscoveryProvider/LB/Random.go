package LB

import (
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"math/rand"
)

type random struct {
	s ServiceDiscovery.IServiceDiscovery
	r *rand.Rand
}

// NewRandom returns a load balancer that selects services randomly.
func NewRandom(sd ServiceDiscovery.IServiceDiscovery, seed int64) ServiceDiscovery.Balancer {
	return &random{
		s: sd,
		r: rand.New(rand.NewSource(seed)),
	}
}

func (rand *random) Next(serviceName string) (ServiceDiscovery.ServiceInstance, error) {
	endpoints := rand.s.GetAllInstances(serviceName)

	if len(endpoints) <= 0 {
		return nil, ServiceDiscovery.ErrNoEndpoints
	}
	return endpoints[rand.r.Intn(len(endpoints))], nil
}
