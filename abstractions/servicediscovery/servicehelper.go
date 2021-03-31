package servicediscovery

// CopyService make a copy of service
func CopyService(service *Service) *Service {
	// copy service
	s := new(Service)
	*s = *service

	// copy nodes
	nodes := make([]ServiceInstance, len(service.Nodes))
	for j, node := range service.Nodes {
		n := new(DefaultServiceInstance)
		nn := node.(*DefaultServiceInstance)
		*n = *nn
		nodes[j] = n
	}
	s.Nodes = nodes

	return s
}

// Copy makes a copy of services
func Copy(current []*Service) []*Service {
	services := make([]*Service, len(current))
	for i, service := range current {
		services[i] = CopyService(service)
	}
	return services
}
