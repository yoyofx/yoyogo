package DependencyInjection

type ServiceCollection struct {
	serviceDescriptors []ServiceDescriptor
}

func NewServiceCollection() *ServiceCollection {
	return &ServiceCollection{}
}

//Singleton
//Scoped
//Transient
func (sc *ServiceCollection) AddServiceDescriptor(sd ServiceDescriptor) {
	sc.serviceDescriptors = append(sc.serviceDescriptors, sd)
}

func (sc *ServiceCollection) AddSingleton(provider interface{}) {
	sd := NewServiceDescriptorByProvider(provider, Singleton)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddSingletonByName(name string, provider interface{}) {
	sd := NewServiceDescriptorByName(name, provider, Singleton)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddSingletonByImplements(provider interface{}, implements ...interface{}) {
	sd := NewServiceDescriptorByImplements(provider, implements, Singleton)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddTransient(provider interface{}) {
	sd := NewServiceDescriptorByProvider(provider, Transient)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddTransientByName(name string, provider interface{}) {
	sd := NewServiceDescriptorByName(name, provider, Transient)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddTransientByImplements(provider interface{}, implements ...interface{}) {
	sd := NewServiceDescriptorByImplements(provider, implements, Transient)
	sc.AddServiceDescriptor(sd)
}
