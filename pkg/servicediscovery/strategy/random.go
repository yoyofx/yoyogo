package strategy

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"math/rand"
)

type random struct {
	s servicediscovery.IServiceDiscovery
	r *rand.Rand
}

// NewRandom returns a load balancer that selects services randomly.
func NewRandom(sd servicediscovery.IServiceDiscovery, seed int64) servicediscovery.Strategy {
	return &random{
		s: sd,
		r: rand.New(rand.NewSource(seed)),
	}
}

func (rand *random) Next(serviceName string) (servicediscovery.ServiceInstance, error) {
	endpoints := rand.s.GetAllInstances(serviceName)

	if len(endpoints) <= 0 {
		return nil, servicediscovery.ErrNoEndpoints
	}
	return endpoints[rand.r.Intn(len(endpoints))], nil
}
