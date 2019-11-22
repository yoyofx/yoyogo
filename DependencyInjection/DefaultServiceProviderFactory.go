package DependencyInjection

type DefaultServiceProviderFactory struct {
}

func (factory DefaultServiceProviderFactory) CreateServiceProvider(collection ServiceCollection) IServiceProvider {
	return DefaultServiceProvider{collection}
}
