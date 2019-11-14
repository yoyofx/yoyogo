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
