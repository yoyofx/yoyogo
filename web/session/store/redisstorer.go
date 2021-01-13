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

//func NewRedis(client *RedisDataSource) ISessionStore {
//	//return &Redis{ client: client , mMaxLifeTime: 3600 }
//	return nil
//}

func (r *Redis) NewID(id string) string {
	return id
}

func (r *Redis) GC() {}

func (r *Redis) SetValue(sessionID string, key string, value interface{}) {
	_ = r.client.GetHashOps().Put(keyPrefix+sessionID, key, value)
}

func (r *Redis) GetValue(sessionID string, key string) (interface{}, bool) {
	value, err := r.client.GetHashOps().GetString(keyPrefix+sessionID, key)
	return value, err == nil
}

func (r *Redis) GetAllSessionId() []string {
	return nil
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
	r.mMaxLifeTime = lifetime
}
