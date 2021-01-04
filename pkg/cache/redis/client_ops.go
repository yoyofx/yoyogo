package redis

import "time"

type Ops interface {
	Ping() (string, error)
	DeleteKey(keys ...string) (int64, error)
	Exists(key string) (bool, error)
	SetExpire(key string, expiration time.Duration) (bool, error)
	TTL(key string) (time.Duration, error)
	MultiSet(values ...interface{}) error
	SetValue(key string, value interface{}, expiration time.Duration) error
	Set(key string, value string, expiration time.Duration) error
	SetNX(key string, value interface{}) (bool, error)
	Get(key string) (string, error)
	GetValue(key string) ([]byte, error)
	Append(key string, value string) (int64, error)
	StrLen(key string) (int64, error)
	GetRange(key string, start int64, end int64) (string, error)
	RandomKey() (string, error)
	MultiGet(key ...string) ([]interface{}, error)
	IncrBy(key string, step int64) (int64, error)
	LPush(key string, values ...interface{}) (int64, error)
}
