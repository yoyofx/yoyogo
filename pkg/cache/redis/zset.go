package redis

import "strconv"

type ZSet struct {
	ops Ops
}

type ZStoreEnum string

const (
	MAX ZStoreEnum = "MAX"
	MIN ZStoreEnum = "MIN"
	SUM ZStoreEnum = "SUM"
)

type ZMember struct {
	Score  float64
	Member string
}

type ZStore struct {
	Key    string
	Weight float64
}

func (zSet *ZSet) ZAdd(key string, member ZMember) int64 {
	return zSet.ops.ZAdd(key, member)
}

func (zSet *ZSet) ZCard(key string) int64 {
	return zSet.ops.ZCard(key)
}

func (zSet *ZSet) ZCount(key string, min, max float64) int64 {
	return zSet.ops.ZCount(key, strconv.FormatFloat(min, 'E', -1, 32), strconv.FormatFloat(max, 'E', -1, 32))
}

func (zSet *ZSet) ZIncrby(key string, incr float64, member string) float64 {
	return zSet.ops.ZIncrby(key, incr, member)
}

func (zSet *ZSet) ZInterStore(destination string, store []ZStore, arg ZStoreEnum) int64 {
	return zSet.ops.ZInterStore(destination, store, arg)
}

func (zSet *ZSet) ZLexCount(key, min, max string) int64 {
	return zSet.ops.ZLexCount(key, min, max)
}

func (zSet *ZSet) ZRange(key string, start, stop int64) []string {
	return zSet.ops.ZRange(key, start, stop)
}

func (zSet *ZSet) ZRangeByLex(key, min, max string, offset int64, count int64) []string {
	return zSet.ops.ZRangeByLex(key, min, max, offset, count)
}

func (zSet *ZSet) ZRangeByScore(key, min, max string, offset int64, count int64) []string {
	return zSet.ops.ZRangeByScore(key, min, max, offset, count)
}

func (zSet *ZSet) ZRank(key, member string) int64 {
	return zSet.ops.ZRank(key, member)
}

func (zSet *ZSet) ZRem(key string, member ...string) int64 {
	return zSet.ops.ZRem(key, member...)
}

func (zSet *ZSet) ZRemRangeByLex(key, min, max string) int64 {
	return zSet.ops.ZRemRangeByLex(key, min, max)
}

func (zSet *ZSet) ZRemRangeByRank(key string, start, stop int64) int64 {
	return zSet.ops.ZRemRangeByRank(key, start, stop)
}

func (zSet *ZSet) ZRevRange(key string, start, stop int64) []string {
	return zSet.ops.ZRevRange(key, start, stop)
}

func (zSet *ZSet) ZRevRank(key, member string) int64 {
	return zSet.ops.ZRevRank(key, member)
}

func (zSet *ZSet) ZScore(key, member string) float64 {
	return zSet.ops.ZScore(key, member)
}
