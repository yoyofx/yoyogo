package servicediscovery

import (
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/web/context"
	"strconv"
)

func CreateServiceInstance(environment *context.HostEnvironment) servicediscovery.ServiceInstance {
	port, _ := strconv.ParseInt(environment.Port, 10, 64)

	return &servicediscovery.DefaultServiceInstance{
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
