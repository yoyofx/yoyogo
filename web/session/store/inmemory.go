package store

import (
	"sync"
	"time"
)

// Memory In-memory session store
type Memory struct {
	mLock        sync.RWMutex
	mMaxLifeTime int64
	mSessions    map[string]*session
}

func NewMemory() *Memory {
	return &Memory{
		mMaxLifeTime: 3600,
		mSessions:    make(map[string]*session),
	}
}

func (mem *Memory) UpdateLastTimeAccessed(sessionId string) {
	mem.mSessions[sessionId].UpdateLastTimeAccessed()
}

func (mem *Memory) SetMaxLifeTime(lifetime int64) {
	mem.mMaxLifeTime = lifetime
}

//NewID new session id
func (mem *Memory) NewID(sessionId string) string {
	mem.mLock.Lock()
	defer mem.mLock.Unlock()
	_, hasSession := mem.mSessions[sessionId]
	if !hasSession {
		newSession := NewSessionInfo(sessionId)
		mem.mSessions[sessionId] = newSession
	} else {
		mem.UpdateLastTimeAccessed(sessionId)
	}
	return sessionId
}

// GetAllSessionId get all session id list
func (mem *Memory) GetAllSessionId() []string {
	mem.mLock.RLock()
	defer mem.mLock.RUnlock()

	sessionIDList := make([]string, 0)
	for k, _ := range mem.mSessions {
		sessionIDList = append(sessionIDList, k)
	}
	return sessionIDList[0:len(sessionIDList)]
}

// SetValue set session value with client
func (mem *Memory) SetValue(sessionID string, key string, value interface{}) {
	mem.mLock.Lock()
	defer mem.mLock.Unlock()
	if s, ok := mem.mSessions[sessionID]; ok {
		s.SetValue(key, value)
	}
}

// GetValue get session value with client
func (mem *Memory) GetValue(sessionID string, key string) (interface{}, bool) {
	mem.mLock.RLock()
	defer mem.mLock.RUnlock()
	return mem.mSessions[sessionID].GetValue(key)
}

// Clear clear session
func (mem *Memory) Clear() {
	mem.mLock.Lock()
	defer mem.mLock.Unlock()
	mem.mSessions = make(map[string]*session)
}

//GC session at time
func (mem *Memory) GC() {
	mem.mLock.Lock()
	defer mem.mLock.Unlock()
	for sessionID, s := range mem.mSessions {
		//删除超过时限的session
		if s.GetLastTimeAccessed().Unix()+mem.mMaxLifeTime < time.Now().Unix() {
			delete(mem.mSessions, sessionID)
		}
	}
	time.AfterFunc(time.Duration(mem.mMaxLifeTime)*time.Second, func() { mem.GC() })
}

// Remove remove session by id
func (mem *Memory) Remove(sessionId string) {
	mem.mLock.Lock()
	defer mem.mLock.Unlock()

	delete(mem.mSessions, sessionId)
}
