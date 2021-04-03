package strategy

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"math/rand"
)

type random struct {
	//s servicediscovery.IServiceDiscovery
	r *rand.Rand
}

// NewRandom returns a load balancer that selects services randomly.
func NewRandom(sd servicediscovery.IServiceDiscovery, seed int64) servicediscovery.Strategy {
	return &random{
		//s: sd,
		r: rand.New(rand.NewSource(seed)),
	}
}

func (rand *random) Next(instanceList []servicediscovery.ServiceInstance) (servicediscovery.ServiceInstance, error) {
	//endpoints := rand.s.GetAllInstances(serviceName)

	if len(instanceList) <= 0 {
		return nil, servicediscovery.ErrNoEndpoints
	}
	return instanceList[rand.r.Intn(len(instanceList))], nil
}
