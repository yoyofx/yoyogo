package abstractions

import "github.com/yoyofx/yoyogo/dependencyinjection"

type IHostService interface {
	Run() error
	Stop() error
}

func AddHostService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingletonByImplements(serviceCtor, new(IHostService))
}
