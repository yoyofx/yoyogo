package redis

import "time"

//IClient redis client interface
type IClient interface {
	//SetSerializer set value serializer
	SetSerializer(ISerializer)
	//GetKVOps Returns the operations performed on simple values (or Strings in Redis terminology).
	GetKVOps() KV
	//GetListOps Returns the operations performed on list values.
	GetListOps() List
	//GetHashOps Returns the operations performed on hash values.
	GetHashOps() Hash
	//GetSetOps Returns the operations performed on set values.
	GetSetOps() Set
	//GetZSetOps Returns the operations performed on zset values (also known as sorted sets).
	GetZSetOps() ZSet
	//GetGeoOps Geo Returns the operations performed on geo values (also known GIS system).
	GetGeoOps() Geo
	//GetLockOps Returns the operations performed on locker values.
	GetLockOps() Lock

	// SetExpire cmd by expire
	Ping() (string, error)
	// SetExpire cmd by expire
	SetExpire(key string, expiration time.Duration) (bool, error)
	// SetExpire cmd by TTL
	GetExpire(key string) (time.Duration, error)
	// Delete delete the key, cmd by del
	Delete(key string) bool
	// HasKey exists the key, cmd by exists
	HasKey(key string) bool
	// RandomKey return random Key for db
	RandomKey() (string, error)
}
