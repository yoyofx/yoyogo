package servicediscovery

type IServiceDiscoveryClient interface {
	GetAllServices() ([]*Service, error)
	GetAllServiceNames() ([]string, error)
	GetAllInstances(serviceName string) []ServiceInstance
	Watch(opts ...WatchOption) (Watcher, error)
	GetService(name string) (*Service, error)
}
