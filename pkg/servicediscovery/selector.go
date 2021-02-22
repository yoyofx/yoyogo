package servicediscovery

import "github.com/yoyofx/yoyogo/abstractions/servicediscovery"

type Selector struct {
	discoveryCache Cache                     //service discovery cache
	strategy       servicediscovery.Strategy //load balancing strategy

}

// will set strategy and cache options
// Selector( strategy ,  cache ).Select(serviceName).(ServiceInstance)
func (s *Selector) Select(serviceName string) (servicediscovery.ServiceInstance, error) {
	//service:= s.discoveryCache.GetService(serviceName)
	//return  s.strategy.Next(service)
	return nil, nil
}
