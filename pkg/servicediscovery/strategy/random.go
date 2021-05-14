package strategy

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"math/rand"
	"time"
)

type random struct {
	//s servicediscovery.IServiceDiscovery
	r *rand.Rand
}

// NewRandom returns a load balancer that selects services randomly.
func NewRandom() servicediscovery.Strategy {
	return &random{
		//s: sd,
		r: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (rand *random) Next(instanceList []servicediscovery.ServiceInstance) (servicediscovery.ServiceInstance, error) {
	//endpoints := rand.s.GetAllInstances(serviceName)

	if len(instanceList) <= 0 {
		return nil, servicediscovery.ErrNoEndpoints
	}
	return instanceList[rand.r.Intn(len(instanceList))], nil
}
