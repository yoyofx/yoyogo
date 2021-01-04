package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx = context.Background()
)

type GoRedisStandaloneOps struct {
	client *redis.Client
}

func NewStandaloneOps(options *Options) *GoRedisStandaloneOps {
	client := redis.NewClient(&redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	})
	return &GoRedisStandaloneOps{client: client}
}

func (ops *GoRedisStandaloneOps) Ping() (string, error) {
	return ops.client.Ping(ctx).Result()
}

// value ops

func (ops *GoRedisStandaloneOps) GetRange(key string, start int64, end int64) (string, error) {
	return ops.client.GetRange(ctx, key, start, end).Result()
}

func (ops *GoRedisStandaloneOps) StrLen(key string) (int64, error) {
	return ops.client.StrLen(ctx, key).Result()
}

func (ops *GoRedisStandaloneOps) Append(key string, value string) (int64, error) {
	return ops.client.Append(ctx, key, value).Result()
}

func (ops *GoRedisStandaloneOps) DeleteKey(keys ...string) (int64, error) {
	return ops.client.Del(ctx, keys...).Result()
}

func (ops *GoRedisStandaloneOps) Exists(key string) (bool, error) {
	n, e := ops.client.Exists(ctx, key).Result()
	return n > 0, e
}

func (ops *GoRedisStandaloneOps) SetExpire(key string, expiration time.Duration) (bool, error) {
	return ops.client.Expire(ctx, key, expiration).Result()
}

func (ops *GoRedisStandaloneOps) TTL(key string) (time.Duration, error) {
	return ops.client.TTL(ctx, key).Result()
}

// MultiSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (ops *GoRedisStandaloneOps) MultiSet(values ...interface{}) error {
	return ops.client.MSet(ctx, values).Err()
}

func (ops *GoRedisStandaloneOps) SetValue(key string, value interface{}, expiration time.Duration) error {
	return ops.client.Set(ctx, key, value, expiration).Err()
}

func (ops *GoRedisStandaloneOps) Set(key string, value string, expiration time.Duration) error {
	return ops.client.Set(ctx, key, value, expiration).Err()
}

func (ops *GoRedisStandaloneOps) SetNX(key string, value interface{}) (bool, error) {
	return ops.client.SetNX(ctx, key, value, 0).Result()
}

func (ops *GoRedisStandaloneOps) GetValue(key string) ([]byte, error) {
	return ops.client.Get(ctx, key).Bytes()
}

func (ops *GoRedisStandaloneOps) Get(key string) (string, error) {
	return ops.client.Get(ctx, key).Result()
}

func (ops *GoRedisStandaloneOps) MultiGet(key ...string) ([]interface{}, error) {
	return ops.client.MGet(ctx, key...).Result()
}

func (ops *GoRedisStandaloneOps) IncrBy(key string, step int64) (int64, error) {
	return ops.client.IncrBy(ctx, key, step).Result()
}

func (ops *GoRedisStandaloneOps) RandomKey() (string, error) {
	return ops.client.RandomKey(ctx).Result()
}

//---------------------------------------------------------------------------------------------------
// list ops

func (ops *GoRedisStandaloneOps) LPush(key string, values ...interface{}) (int64, error) {
	return ops.client.LPush(ctx, key, values).Result()
}
