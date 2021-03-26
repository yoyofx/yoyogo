package servicediscovery

import (
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"strconv"
	"strings"
)

func CreateServiceInstance(environment *abstractions.HostEnvironment) servicediscovery.ServiceInstance {
	port, _ := strconv.ParseInt(environment.Port, 10, 64)

	return &servicediscovery.DefaultServiceInstance{
		Id:          uuid.New().String(),
		ServiceName: strings.ToUpper(environment.ApplicationName),
		Host:        environment.Host,
		Port:        uint64(port),
		Enable:      true,
		Healthy:     true,
		Metadata: map[string]string{
			"version": environment.Version,
		},
	}
}
