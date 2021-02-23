package servicediscovery

import (
	"errors"
	"sync"
	"time"
)

type ServiceCache struct {
	sync.RWMutex
	cache map[string][]ServiceInstance
	ttls  map[string]time.Time
}

/**
初始化缓存
*/
func NewServiceCache(discovery IServiceDiscovery, cacheSecond int64) *ServiceCache {
	serviceInstance := make(map[string][]ServiceInstance)
	ttls := make(map[string]time.Time)
	allServices, _ := discovery.GetAllServices()
	for _, s := range allServices {
		serviceInstance[s.Name] = discovery.GetAllInstances(s.Name)
		//默认缓存30秒
		if cacheSecond == 0 {
			cacheSecond = 30000
		}
		ttls[s.Name] = time.Now().Add(time.Duration(cacheSecond))
	}
	return &ServiceCache{
		cache: serviceInstance,
		ttls:  ttls,
	}
}

/**
更新全部
*/
func (cache *ServiceCache) UpdateAll(discovery IServiceDiscovery, cacheSecond int64) {
	cache.Lock()
	defer cache.Unlock()
	allServices, _ := discovery.GetAllServices()
	for _, s := range allServices {
		cache.cache[s.Name] = discovery.GetAllInstances(s.Name)
		cache.ttls[s.Name] = time.Now().Add(time.Duration(cacheSecond))
	}
}

/**
更新单个
*/
func (cache *ServiceCache) UpdateService(serviceName string, instances []ServiceInstance, cacheSecond int64) (bool, error) {
	cache.Lock()
	defer cache.Unlock()
	if serviceName != "" {
		return false, errors.New("serviceName can not be empty")
	}
	cache.cache[serviceName] = instances
	cache.ttls[serviceName] = time.Now().Add(time.Duration(cacheSecond))
	return true, nil
}

/**
获取缓存的服务信息
 */
func (cache *ServiceCache) GetService(serviceName string) ([]ServiceInstance, error) {
	if serviceName == "" {
		return nil, errors.New("serviceName can not be empty")
	}
	ttl := cache.ttls[serviceName]
	if ttl.Before(time.Now()) {
		return nil, errors.New("service is timeout")
	}
	services := cache.cache[serviceName]
	return services, nil
}

func ()  {
	
}

