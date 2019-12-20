package DependencyInjection

type ServiceDescriptor struct {
	Name       string
	Provider   interface{}
	Implements interface{}
	Lifetime   ServiceLifetime
}

func NewServiceDescriptor(name string, provider interface{}, implements interface{}, lifetime ServiceLifetime) *ServiceDescriptor {
	return &ServiceDescriptor{Name: name, Provider: provider, Implements: implements, Lifetime: lifetime}
}

func NewServiceDescriptorByProvider(provider interface{}, lifetime ServiceLifetime) *ServiceDescriptor {
	return &ServiceDescriptor{Provider: provider, Lifetime: lifetime}
}

func NewServiceDescriptorByName(name string, provider interface{}, lifetime ServiceLifetime) *ServiceDescriptor {
	return &ServiceDescriptor{Name: name, Provider: provider, Lifetime: lifetime}
}

func NewServiceDescriptorByImplements(provider interface{}, implements interface{}, lifetime ServiceLifetime) *ServiceDescriptor {
	return &ServiceDescriptor{Provider: provider, Implements: implements, Lifetime: lifetime}
}
