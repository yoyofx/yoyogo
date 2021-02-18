package servicediscovery

type Service struct {
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	Metadata map[string]string `json:"metadata"`
	Nodes    []ServiceInstance `json:"nodes"`
}
