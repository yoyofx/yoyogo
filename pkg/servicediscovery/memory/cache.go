package memory

import "github.com/yoyofx/yoyogo/abstractions/servicediscovery"

type MemoryCache struct {
}

func (memoryCache *MemoryCache) GetService(serviceName string) (*servicediscovery.Service, error) {
	services := []string{"mnurtestapi.mengniu.com.cn", "mnurtestapi.mengniu.com.cn", "mnurtestapi.mengniu.com.cn"}
	sd := NewServerDiscovery("operations", services)
	s := &servicediscovery.Service{Name: "operations", Nodes: sd.GetAllInstances("operations")}
	return s, nil
}

func (memoryCache *MemoryCache) Stop() {

	panic("no ")
}
