package redis

import (
	"time"
)

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
	GeoAddArr(key string, geoLocation ...GeoPosition) int64
	GeoPos(key string, members ...string) (error, []GeoPosition)
	GeoDist(key string, member1, member2 string, unit GeoUnit) (error,GeoDistInfo)
	GeoRadius(key string, query GeoRadiusQuery) (error, []GeoPosition)
	GeoRadiusByMember(key string, member string, query GeoRadiusByMemberQuery)(error, []GeoPosition)
	SAdd(key string, members ...interface{}) (int64, error)
	SDiff(keys ...string) ([]string, error)
	SCard(key string) (int64, error)
	SInter(keys ...string) ([]string, error)
	SInterStore(destination string, keys ...string) (int64, error)
	SIsMember(key string, member interface{}) (bool, error)
	SMembers(key string) ([]string, error)
	SMove(source string, destination string, member interface{}) (bool, error)
	SPop(key string) (string, error)
	SRandMembers(key string, count int64) ([]string, error)
	SRem(key string, members ...interface{}) (int64, error)
	SUnion(keys ...string) ([]string, error)
	SUnionStore(destination string, keys ...string) (int64, error)
}
