package DependencyInjection

import (
	"github.com/defval/inject/v2"
)

func (sc ServiceCollection) Build() IServiceProvider {
	var providers []inject.Option
	for _, desc := range sc.serviceDescriptors {

		var providerOptions []inject.ProvideOption
		if desc.Implements != nil {
			providerOptions = append(providerOptions, inject.As(desc.Implements))
		}
		if desc.Name != "" {
			providerOptions = append(providerOptions, inject.WithName(desc.Name))
		}
		if desc.Lifetime == Singleton {
			providerOptions = append(providerOptions, inject.Prototype())
		}

		provider := inject.Provide(desc.Provider, providerOptions...)
		providers = append(providers, provider)

	}
	container := inject.New(providers...)

	return &DefaultServiceProvider{container}
}
