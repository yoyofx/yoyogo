package redis

type IClient interface {
	SetConnection(*Ops)
	SetSerializer(ISerializer)
	GetKVOps() KV
	GetListOps() List
	GetHashOps() Hash
	GetSetOps() Set
	GetZSetOps() ZSet
	GetGeoOps() Geo
}
