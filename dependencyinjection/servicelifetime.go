package dependencyinjection

type ServiceLifetime int32

const (
	Singleton ServiceLifetime = 0
	Scoped    ServiceLifetime = 1
	Transient ServiceLifetime = 2
)
