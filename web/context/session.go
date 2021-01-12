package context

// Session current session by http context
type Session struct {
	sessionId string
	manager   ISessionManager
}

func NewSession(sessionId string, mgr ISessionManager) *Session {
	return &Session{
		sessionId: sessionId,
		manager:   mgr,
	}
}

func (session Session) SetValue(key string, value interface{}) {
	session.manager.SetValue(session.sessionId, key, value)
}

func (session Session) GetValue(key string) (interface{}, bool) {
	return session.manager.GetValue(session.sessionId, key)
}

func (session Session) GetString(key string) string {
	val, has := session.manager.GetValue(session.sessionId, key)
	if has {
		return val.(string)
	} else {
		return ""
	}
}

func (session Session) GetInt(key string) int {
	val, has := session.manager.GetValue(session.sessionId, key)
	if has {
		return val.(int)
	} else {
		return 0
	}
}

func (session Session) GetInt64(key string) int64 {
	val, has := session.manager.GetValue(session.sessionId, key)
	if has {
		return val.(int64)
	} else {
		return 0
	}
}

func (session Session) GetFloat(key string) float32 {
	val, has := session.manager.GetValue(session.sessionId, key)
	if has {
		return val.(float32)
	} else {
		return 0
	}
}

func (session Session) GetFloat64(key string) float64 {
	val, has := session.manager.GetValue(session.sessionId, key)
	if has {
		return val.(float64)
	} else {
		return 0
	}
}
