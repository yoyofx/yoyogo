package servicediscovery

import "errors"

// Balancer yields endpoints according to some heuristic.
type Strategy interface {
	Next(serviceName string) (ServiceInstance, error)
}

type Next func() (ServiceInstance, error)

// ErrNoEndpoints is returned when no qualifying endpoints are available.
var ErrNoEndpoints = errors.New("no endpoints available")
