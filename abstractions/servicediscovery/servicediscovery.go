package servicediscovery

type IServiceDiscovery interface {
	GetName() string
	Register() error
	Update() error
	Unregister() error
	GetHealthyInstances(serviceName string) []ServiceInstance
	GetAllInstances(serviceName string) []ServiceInstance
	Destroy() error
	Watch(opts ...WatchOption) (Watcher, error)
	GetAllServices() ([]*Service, error)
}
