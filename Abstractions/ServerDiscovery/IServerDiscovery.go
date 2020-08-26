package ServerDiscovery

type IServerDiscovery interface {
	GetName() string
	Register() error
	Update() error
	Unregister() error
	GetInstances(serviceName string) []ServiceInstance
	Destroy() error
}
