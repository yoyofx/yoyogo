package redis

import "time"

type Ops interface {
	Ping() (string, error)
	GetRange(key string, start int64, end int64) (string, error)
	StrLen(key string) (int64, error)
	Append(key string, value string) (int64, error)
	DeleteKey(keys ...string) (int64, error)
	Exists(key string) (bool, error)
	SetExpire(key string, expiration time.Duration) (bool, error)
	TTL(key string) (time.Duration, error)
	MultiSet(values ...interface{}) error
	SetValue(key string, value interface{}, expiration time.Duration) error
	Set(key string, value string, expiration time.Duration) error
	SetNX(key string, value interface{}) (bool, error)
	GetValue(key string) ([]byte, error)
	Get(key string) (string, error)
	MultiGet(key ...string) ([]interface{}, error)
	IncrBy(key string, step int64) (int64, error)
	RandomKey() (string, error)
	LIndex(key string, index int64) (string, error)
	LPop(key string) (string, error)
	LPush(key string, values ...interface{}) (int64, error)
	LRange(key string, start int64, end int64) ([]string, error)
	LTrim(key string, start int64, end int64) error
	RPop(key string) (string, error)
	RPush(key string, values ...interface{}) (int64, error)
	LSet(key string, index int64, value interface{}) error
	LSize(key string) (int64, error)
	LRemove(key string, count int64, value interface{}) (int64, error)
}
