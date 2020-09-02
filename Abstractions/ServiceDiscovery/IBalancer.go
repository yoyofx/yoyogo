package ServiceDiscovery

import "errors"

// Balancer yields endpoints according to some heuristic.
type Balancer interface {
	Next(serviceName string) (ServiceInstance, error)
}

// ErrNoEndpoints is returned when no qualifying endpoints are available.
var ErrNoEndpoints = errors.New("no endpoints available")
