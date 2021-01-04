package store

import "time"

type session struct {
	mSessionID        string                 //唯一id
	mLastTimeAccessed time.Time              //最后访问时间
	mValues           map[string]interface{} //其它对应值保存用户所对应的一些值
}

func NewSessionInfo(sessionId string) *session {
	return &session{
		mSessionID:        sessionId,
		mLastTimeAccessed: time.Now(),
		mValues:           make(map[string]interface{}),
	}
}

func (session *session) SetValue(key string, value interface{}) {
	session.mValues[key] = value
}

func (session *session) GetValue(key string) (interface{}, bool) {
	val, ok := session.mValues[key]
	return val, ok
}

func (session *session) GetLastTimeAccessed() time.Time {
	return session.mLastTimeAccessed
}

func (session *session) UpdateLastTimeAccessed() {
	session.mLastTimeAccessed = time.Now()
}
