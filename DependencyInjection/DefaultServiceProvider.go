package DependencyInjection

type DefaultServiceProvider struct {
	servicecollection ServiceCollection
}

func (d DefaultServiceProvider) GetService(refObject interface{}, providerType interface{}) {
	panic("implement me")
}

func (d DefaultServiceProvider) GetServiceByName(refObject interface{}, providerType interface{}, name string) {
	panic("implement me")
}

func (d DefaultServiceProvider) GetServices(refObjects interface{}, implementType interface{}) {
	panic("implement me")
}
