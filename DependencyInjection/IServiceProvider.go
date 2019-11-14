package DependencyInjection

type IServiceProvider interface {
	GetService(providerType interface{}) interface{}

	GetServiceByName(providerType interface{}, name string) interface{}
}
