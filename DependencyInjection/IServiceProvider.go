package DependencyInjection

type IServiceProvider interface {
	GetService(refObject interface{}) error
	GetServiceByName(refObject interface{}, name string) error
	GetGraph() string
}
