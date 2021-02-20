package servicediscovery

import "github.com/yoyofx/yoyogo/abstractions/servicediscovery"

type Client struct {
	discoveryClient servicediscovery.IServiceDiscovery
}

func NewClient(discovery servicediscovery.IServiceDiscovery) *Client {
	return &Client{discoveryClient: discovery}
}

func (c *Client) GetAllInstances(serviceName string) []servicediscovery.ServiceInstance {
	return c.discoveryClient.GetAllInstances(serviceName)
}

func (c *Client) GetAllServices() ([]*servicediscovery.Service, error) {
	serviceList, _ := c.discoveryClient.GetAllServices()
	for _, service := range serviceList {
		service.Nodes = c.discoveryClient.GetAllInstances(service.Name)
	}
	return serviceList, nil
}

func (c *Client) GetAllServiceNames() ([]*servicediscovery.Service, error) {
	return c.discoveryClient.GetAllServices()
}

func (c *Client) Watch(opts ...servicediscovery.WatchOption) (servicediscovery.Watcher, error) {
	return c.discoveryClient.Watch(opts...)
}
