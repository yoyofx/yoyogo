package abstractions

import "github.com/yoyofx/yoyogo/dependencyinjection"

type IStartup interface {
	ConfigureServices(collection *dependencyinjection.ServiceCollection)
}
