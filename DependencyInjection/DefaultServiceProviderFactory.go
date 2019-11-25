package DependencyInjection

import "github.com/maxzhang1985/inject"

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
			providerOptions = append(providerOptions, inject.Lifetime(Singleton))
		} else {
			providerOptions = append(providerOptions, inject.Lifetime(Transient))
		}

		provider := inject.Provide(desc.Provider, providerOptions...)
		providers = append(providers, provider)

	}
	container, err := inject.New(providers...)
	if err != nil {
		panic(err)
		return nil
	}

	return DefaultServiceProvider{container}
}
