package ServerDiscovery

// ServiceInstance is the model class of an instance of a service, which is used for service registration and discovery.
type ServiceInstance interface {

	// GetId will return this instance's id. It should be unique.
	GetId() string

	// GetServiceName will return the serviceName
	GetServiceName() string

	// GetHost will return the hostname
	GetHost() string

	// GetPort will return the port.
	GetPort() int

	// IsEnable will return the enable status of this instance
	IsEnable() bool

	// IsHealthy will return the value represent the instance whether healthy or not
	IsHealthy() bool

	// GetMetadata will return the metadata
	GetMetadata() map[string]string
}

// DefaultServiceInstance the default implementation of ServiceInstance
// or change the ServiceInstance to be struct???
type DefaultServiceInstance struct {
	Id          string
	ServiceName string
	Host        string
	Port        int
	ClusterName string
	GroupName   string
	Enable      bool
	Healthy     bool
	Metadata    map[string]string
}

// GetId will return this instance's id. It should be unique.
func (d *DefaultServiceInstance) GetId() string {
	return d.Id
}

// GetServiceName will return the serviceName
func (d *DefaultServiceInstance) GetServiceName() string {
	return d.ServiceName
}

// GetHost will return the hostname
func (d *DefaultServiceInstance) GetHost() string {
	return d.Host
}

// GetPort will return the port.
func (d *DefaultServiceInstance) GetPort() int {
	return d.Port
}

// IsEnable will return the enable status of this instance
func (d *DefaultServiceInstance) IsEnable() bool {
	return d.Enable
}

// IsHealthy will return the value represent the instance whether healthy or not
func (d *DefaultServiceInstance) IsHealthy() bool {
	return d.Healthy
}

// GetMetadata will return the metadata, it will never return nil
func (d *DefaultServiceInstance) GetMetadata() map[string]string {
	if d.Metadata == nil {
		d.Metadata = make(map[string]string, 0)
	}
	return d.Metadata
}
