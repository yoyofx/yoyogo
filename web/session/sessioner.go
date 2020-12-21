package session

type ISession interface {
	NewID() string
	GetKeys() []string
	Clear()
	Load() string
	Remove(key string)
	Set(key string, value interface{})
	Get(key string) interface{}
}
