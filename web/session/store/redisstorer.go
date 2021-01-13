package store

import (
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
	"time"
)

var (
	keyPrefix = "session:"
)

type Redis struct {
	client       redis.IClient
	mMaxLifeTime int64
}

func NewRedis() ISessionStore {
	return &Redis{mMaxLifeTime: 3600}
}

func (r *Redis) NewID(id string) string {
	return ""
}

func (r *Redis) GC() {}

func (r *Redis) SetValue(sessionID string, key string, value interface{}) {
	panic("implement me")
}

func (r *Redis) GetValue(sessionID string, key string) (interface{}, bool) {
	panic("implement me")
}

func (r *Redis) GetAllSessionId() []string {
	panic("Not support method")
}

func (r *Redis) Clear() {
	panic("Not support method")
}

func (r *Redis) Remove(sessionId string) {
	r.client.Delete(keyPrefix + sessionId)
}

func (r *Redis) UpdateLastTimeAccessed(sessionId string) {
	_, _ = r.client.SetExpire(keyPrefix+sessionId, time.Duration(r.mMaxLifeTime))
}

func (r *Redis) SetMaxLifeTime(lifetime int64) {
	panic("implement me")
}
