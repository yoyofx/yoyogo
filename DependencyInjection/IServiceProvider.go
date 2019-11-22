package DependencyInjection

type IServiceProvider interface {
	GetService(refObject interface{}, providerType interface{})

	GetServiceByName(refObject interface{}, providerType interface{}, name string)

	GetServices(refObjects interface{}, implementType interface{})
}
