package servicediscovery

type IServiceDiscoveryClient interface {
	GetAllServices() ([]*Service, error)
	GetAllServiceNames() ([]*Service, error)
	GetAllInstances(serviceName string) []ServiceInstance
	Watch(opts ...WatchOption) (Watcher, error)
}
