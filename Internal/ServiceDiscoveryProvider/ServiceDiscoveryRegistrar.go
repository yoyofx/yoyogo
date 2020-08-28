package ServiceDiscoveryProvider

import (
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"strconv"
)

func CreateServiceInstance(environment *Context.HostEnvironment) ServiceDiscovery.ServiceInstance {
	port, _ := strconv.ParseInt(environment.Port, 10, 64)

	return &ServiceDiscovery.DefaultServiceInstance{
		Id:          uuid.New().String(),
		ServiceName: environment.ApplicationName,
		Host:        environment.Host,
		Port:        uint64(port),
		Enable:      true,
		Healthy:     true,
		Metadata: map[string]string{
			"VERSION": environment.Version,
		},
	}
}
