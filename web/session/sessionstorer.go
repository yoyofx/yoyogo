package session

type ISessionStore interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Remove(key string)
}
