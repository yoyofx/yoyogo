package servicediscovery

type Selector struct {
	discoveryCache Cache    //service discovery cache
	strategy       Strategy //load balancing strategy

}

// will set strategy and cache options
// Selector( strategy ,  cache ).Select(serviceName).(ServiceInstance)
func (s *Selector) Select(serviceName string) (ServiceInstance, error) {
	//service:= s.discoveryCache.GetService(serviceName)
	//return  s.strategy.Next(service)
	return nil, nil
}
