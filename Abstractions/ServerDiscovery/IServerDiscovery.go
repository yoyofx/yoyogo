package ServerDiscovery

type IServerDiscovery interface {
	GetName() string
	Register() error
	Update() error
	Unregister() error
	GetHealthyInstances(serviceName string) []ServiceInstance
	GetAllInstances(serviceName string) []ServiceInstance
	Destroy() error
}
