package grpc

import "github.com/yoyofx/yoyogo/dependencyinjection"

func AddService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingleton(serviceCtor)
}
