package store

type ISessionStore interface {
	NewID(id string) string
	GC()
	SetValue(sessionID string, key string, value interface{})
	GetValue(sessionID string, key string) (interface{}, bool)
	GetAllSessionId() []string
	Clear()
	Remove(sessionId string)
	UpdateLastTimeAccessed(sessionId string)
	SetMaxLifeTime(lifetime int64)
}
