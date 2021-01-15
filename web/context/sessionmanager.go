package context

//ISessionManager session manager of interface
type ISessionManager interface {
	GetIDList() []string
	Clear(interface{})
	Load(interface{}) string
	NewSession(sessionId string) string
	Remove(sessionId string)
	SetValue(sessionID string, key string, value interface{})
	GetValue(sessionID string, key string) (interface{}, bool)
	GC()
}
