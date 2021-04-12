package dependencyinjection

type IServiceProvider interface {
	GetService(refObject interface{}) error
	GetServiceByName(refObject interface{}, name string) error
	GetGraph() string
	InvokeService(fn interface{}) error
}
