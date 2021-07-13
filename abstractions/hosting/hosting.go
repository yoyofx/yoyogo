package hosting

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

func AddHostService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingletonByImplements(serviceCtor, new(abstractions.IHostService))
}
