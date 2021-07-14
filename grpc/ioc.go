package grpc

import "github.com/yoyofxteam/dependencyinjection"

func AddService(collection *dependencyinjection.ServiceCollection, serviceCtor interface{}) {
	collection.AddSingleton(serviceCtor)
}
