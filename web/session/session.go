package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strconv"
	"sync"
	"time"
)

type Manager struct {
	mLock        sync.RWMutex
	mMaxLifeTime int64
	store        ISessionStore
	mSessions    map[string]*sessionInfo //

}

// NewID new session id
func (mgr *Manager) NewID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano() //微秒
		return strconv.FormatInt(nano, 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// GetKeys get all session id to array
func (mgr *Manager) GetKeys() []string {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()

	sessionIDList := make([]string, 0)
	for k, _ := range mgr.mSessions {
		sessionIDList = append(sessionIDList, k)
	}
	return sessionIDList[0:len(sessionIDList)]
}

// Load init session information to the session store
func (mgr *Manager) Load() string {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	newSessionID := mgr.NewID()
	session := &sessionInfo{mSessionID: newSessionID,
		mLastTimeAccessed: time.Now(), mValues: make(map[interface{}]interface{})}
	mgr.mSessions[newSessionID] = session
	// set value to the session store
	mgr.store.Set("sessionId", newSessionID)
	return newSessionID
}

// Clear clear session
func (mgr *Manager) Clear() {
	panic("implement me")
}

func (mgr *Manager) Remove(key string) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	delete(mgr.mSessions, key)
}

func (mgr *Manager) Set(key string, value interface{}) {
	panic("implement me")
}

func (mgr *Manager) Get(key string) interface{} {
	panic("implement me")
}

// GC GC session at time
func (mgr *Manager) GC() {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	for sessionID, session := range mgr.mSessions {
		//删除超过时限的session
		if session.mLastTimeAccessed.Unix()+mgr.mMaxLifeTime < time.Now().Unix() {
			delete(mgr.mSessions, sessionID)
		}
	}
	//定时回收
	time.AfterFunc(time.Duration(mgr.mMaxLifeTime)*time.Second, func() { mgr.GC() })
}

// Session http user session
type sessionInfo struct {
	mSessionID        string                      //唯一id
	mLastTimeAccessed time.Time                   //最后访问时间
	mValues           map[interface{}]interface{} //其它对应值(保存用户所对应的一些值，比如用户权限之类)
}
