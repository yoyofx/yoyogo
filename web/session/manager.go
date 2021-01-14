package session

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/session/identity"
	"github.com/yoyofx/yoyogo/web/session/store"
)

//ISessionManager session manager
type Manager struct {
	mMaxLifeTime int64
	identity     identity.IProvider
	store        store.ISessionStore
}

//NewSession ctor for session manager
//func NewSession(mMaxLifeTime int64) context.ISessionManager {
//	mgr := &Manager{
//		mMaxLifeTime: mMaxLifeTime,
//		store:        store.NewMemory(mMaxLifeTime),
//	}
//	go mgr.GC()
//	return mgr
//}

//NewSessionWithStore ctor for session manager , must be used to session.UseSession ,that add dependents to IOC.
func NewSessionWithStore(store store.ISessionStore) context.ISessionManager {
	mgr := &Manager{
		store: store,
	}
	go mgr.GC()
	return mgr
}

//GC clear the session list
func (mgr *Manager) GC() {
	mgr.store.GC()
}

// GetIDList get all session id to array
func (mgr *Manager) GetIDList() []string {
	return mgr.store.GetAllSessionId()
}

// Load init and restore session information to the session store
func (mgr *Manager) Load(provider interface{}) string {
	id, _ := provider.(identity.IProvider)
	id.SetMaxLifeTime(mgr.mMaxLifeTime)
	sessionId := id.GetID()
	if sessionId == "" {
		return ""
	}
	return sessionId
	//return mgr.store.NewID(sessionId)
}

func (mgr *Manager) NewSession(sessionId string) string {
	return mgr.store.NewID(sessionId)
}

// Clear clear session
func (mgr *Manager) Clear(provider interface{}) {
	id, _ := provider.(identity.IProvider)
	mgr.store.Clear()
	id.Clear()
}

//Remove remove session store by id
func (mgr *Manager) Remove(sessionId string) {
	mgr.store.Remove(sessionId)
}

//SetValue set session value for the key/value
func (mgr *Manager) SetValue(sessionID string, key string, value interface{}) {
	mgr.store.SetValue(sessionID, key, value)
}

// GetValue get session value for the key
func (mgr *Manager) GetValue(sessionID string, key string) (interface{}, bool) {
	return mgr.store.GetValue(sessionID, key)
}
