package session

import (
	"github.com/yoyofx/yoyogo/web/session/identity"
	"github.com/yoyofx/yoyogo/web/session/store"
)

//IManager session manager of interface
type IManager interface {
	GetIDList() []string
	Clear(identity.IProvider)
	Load(identity.IProvider) string
	Remove(sessionId string)
	SetValue(sessionID string, key string, value interface{})
	GetValue(sessionID string, key string) (interface{}, bool)
	GC()
}

//IManager session manager
type Manager struct {
	mMaxLifeTime int64
	identity     identity.IProvider
	store        store.ISessionStore
}

//NewSession ctor for session manager
func NewSession(mMaxLifeTime int64) IManager {
	mgr := &Manager{
		mMaxLifeTime: mMaxLifeTime,
		store:        store.NewMemory(mMaxLifeTime),
	}
	go mgr.GC()
	return mgr
}

func NewSessionWithStore(store store.ISessionStore, mMaxLifeTime int64) IManager {
	mgr := &Manager{
		mMaxLifeTime: mMaxLifeTime,
		store:        store,
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
func (mgr *Manager) Load(identity identity.IProvider) string {
	identity.SetMaxLifeTime(mgr.mMaxLifeTime)
	sessionId := identity.GetID()
	return mgr.store.NewID(sessionId)
}

// Clear clear session
func (mgr *Manager) Clear(identity identity.IProvider) {
	mgr.store.Clear()
	identity.Clear()
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
