package redis

import (
	"time"
)

type Ops interface {
	Close() error
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
	GeoDist(key string, member1, member2 string, unit GeoUnit) (error, GeoDistInfo)
	GeoRadius(key string, query GeoRadiusQuery) (error, []GeoPosition)
	GeoRadiusByMember(key string, member string, query GeoRadiusByMemberQuery) (error, []GeoPosition)
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
	HDel(key string, fields ...string) (int64, error)
	HExists(key string, field string) (bool, error)
	HGet(key string, field string) (string, error)
	HGetAll(key string) (map[string]string, error)
	HIncrBy(key string, field string, increment int64) (int64, error)
	HKeys(key string) ([]string, error)
	HLen(key string) (int64, error)
	HMGet(key string, fields ...string) ([]interface{}, error)
	HSet(key string, field string, value interface{}) (int64, error)
	HSetNX(key string, field string, value interface{}) (bool, error)
	HVals(key string) ([]string, error)
	ZAdd(key string, member ZMember) int64
	ZCard(key string) int64
	ZCount(key, min, max string) int64
	ZIncrby(key string, incr float64, member string) float64
	ZInterStore(destination string, store []ZStore, arg ZStoreEnum) int64
	ZLexCount(key, min, max string) int64
	ZRange(key string, start, stop int64) []string
	ZRangeByLex(key, min, max string, offset int64, count int64) []string
	ZRangeByScore(key, min, max string, offset int64, count int64) []string
	ZRank(key, member string) int64
	ZRem(key string, member ...string) int64
	ZRemRangeByLex(key, min, max string) int64
	ZRemRangeByRank(key string, start, stop int64) int64
	ZRevRange(key string, start, stop int64) []string
	ZRevRangeWithScores(key string, start, stop int64) ([]ZMember, error)
	ZRevRank(key, member string) int64
	ZScore(key, member string) float64
}
