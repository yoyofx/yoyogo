package dependencyinjection

type IServiceProviderFactory interface {
	CreateServiceProvider() IServiceProvider
}
