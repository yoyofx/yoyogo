package abstractions

import "github.com/yoyofxteam/dependencyinjection"

type IStartup interface {
	ConfigureServices(collection *dependencyinjection.ServiceCollection)
}
