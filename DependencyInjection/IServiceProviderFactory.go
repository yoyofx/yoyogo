package DependencyInjection

type IServiceProviderFactory interface {
	CreateServiceProvider() IServiceProvider
}
