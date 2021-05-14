package dependencyinjection

import (
	"github.com/yoyofxteam/reflectx"
	"strings"
)

type ServiceCollection struct {
	serviceDescriptors    []*ServiceDescriptor
	serviceDescriptorMaps map[string]int
}

func NewServiceCollection() *ServiceCollection {
	return &ServiceCollection{serviceDescriptorMaps: make(map[string]int)}
}

//Singleton
//Scoped
//Transient
func (sc *ServiceCollection) AddServiceDescriptor(sd *ServiceDescriptor) {
	typeName := sd.Name
	if typeName == "" {
		typeName, _ = reflectx.GetCtorFuncOutTypeName(sd.Provider)
		typeName = strings.ToLower(typeName)
	}

	index := len(sc.serviceDescriptors)
	defIndex, exist := sc.serviceDescriptorMaps[typeName]
	if exist {
		sc.serviceDescriptors[defIndex] = sd
	} else {
		sc.serviceDescriptorMaps[typeName] = index
		sc.serviceDescriptors = append(sc.serviceDescriptors, sd)
	}
}

func (sc *ServiceCollection) AddSingleton(provider interface{}) {
	sd := NewServiceDescriptorByProvider(provider, Singleton)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddSingletonAndName(name string, provider interface{}) {
	sc.AddSingletonByName(name, provider)
	sc.AddSingleton(provider)
}

func (sc *ServiceCollection) AddSingletonByName(name string, provider interface{}) {
	sd := NewServiceDescriptorByName(name, provider, Singleton)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddSingletonByImplementsAndName(name string, provider interface{}, implements interface{}) {
	sc.AddSingletonByName(name, provider)
	sc.AddSingletonByImplements(provider, implements)
}

func (sc *ServiceCollection) AddSingletonByImplements(provider interface{}, implements interface{}) {
	sd := NewServiceDescriptorByImplements(provider, implements, Singleton)
	sc.AddServiceDescriptor(sd)
}

func (sc *ServiceCollection) AddSingletonByNameAndImplements(name string, provider interface{}, implements interface{}) {
	sd := NewServiceDescriptorByImplements(provider, implements, Singleton)
	sd.Name = name
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

func (sc *ServiceCollection) AddTransientByImplements(provider interface{}, implements interface{}) {
	sd := NewServiceDescriptorByImplements(provider, implements, Transient)
	sc.AddServiceDescriptor(sd)
}
