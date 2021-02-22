package servicediscovery

import "github.com/yoyofx/yoyogo/abstractions/servicediscovery"

type Cache struct {
	discoveryClient servicediscovery.IServiceDiscoveryClient
}

func (c *Cache) GetService(serviceName string) *servicediscovery.Service {
	// cache as get and refresh service nodes by ttl
	return nil
}
