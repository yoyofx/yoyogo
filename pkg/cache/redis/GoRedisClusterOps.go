package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type GoRedisClusterOps struct {
	client *redis.ClusterClient
}

func NewClusterOps(options *Options) *GoRedisClusterOps {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    options.Addrs,
		Password: options.Password,
	})
	return &GoRedisClusterOps{client: client}
}

func (ops *GoRedisClusterOps) Close() error {
	return ops.client.Close()
}

func (ops *GoRedisClusterOps) Ping() (string, error) {
	return ops.client.Ping(ctx).Result()
}

// value ops

func (ops *GoRedisClusterOps) GetRange(key string, start int64, end int64) (string, error) {
	return ops.client.GetRange(ctx, key, start, end).Result()
}

func (ops *GoRedisClusterOps) StrLen(key string) (int64, error) {
	return ops.client.StrLen(ctx, key).Result()
}

func (ops *GoRedisClusterOps) Append(key string, value string) (int64, error) {
	return ops.client.Append(ctx, key, value).Result()
}

func (ops *GoRedisClusterOps) DeleteKey(keys ...string) (int64, error) {
	return ops.client.Del(ctx, keys...).Result()
}

func (ops *GoRedisClusterOps) Exists(key string) (bool, error) {
	n, e := ops.client.Exists(ctx, key).Result()
	return n > 0, e
}

func (ops *GoRedisClusterOps) SetExpire(key string, expiration time.Duration) (bool, error) {
	return ops.client.Expire(ctx, key, expiration).Result()
}

func (ops *GoRedisClusterOps) TTL(key string) (time.Duration, error) {
	return ops.client.TTL(ctx, key).Result()
}

// MultiSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (ops *GoRedisClusterOps) MultiSet(values ...interface{}) error {
	return ops.client.MSet(ctx, values).Err()
}

func (ops *GoRedisClusterOps) SetValue(key string, value interface{}, expiration time.Duration) error {
	return ops.client.Set(ctx, key, value, expiration).Err()
}

func (ops *GoRedisClusterOps) Set(key string, value string, expiration time.Duration) error {
	return ops.client.Set(ctx, key, value, expiration).Err()
}

func (ops *GoRedisClusterOps) SetNX(key string, value interface{}) (bool, error) {
	return ops.client.SetNX(ctx, key, value, 0).Result()
}

func (ops *GoRedisClusterOps) GetValue(key string) ([]byte, error) {
	return ops.client.Get(ctx, key).Bytes()
}

func (ops *GoRedisClusterOps) Get(key string) (string, error) {
	return ops.client.Get(ctx, key).Result()
}

func (ops *GoRedisClusterOps) MultiGet(key ...string) ([]interface{}, error) {
	return ops.client.MGet(ctx, key...).Result()
}

func (ops *GoRedisClusterOps) IncrBy(key string, step int64) (int64, error) {
	return ops.client.IncrBy(ctx, key, step).Result()
}

func (ops *GoRedisClusterOps) RandomKey() (string, error) {
	return ops.client.RandomKey(ctx).Result()
}

//---------------------------------------------------------------------------------------------------
// list ops
func (ops *GoRedisClusterOps) LPop(key string) (string, error) {
	return ops.client.LPop(ctx, key).Result()
}

func (ops *GoRedisClusterOps) LIndex(key string, index int64) (string, error) {
	return ops.client.LIndex(ctx, key, index).Result()
}

func (ops *GoRedisClusterOps) LPush(key string, values ...interface{}) (int64, error) {
	return ops.client.LPush(ctx, key, values).Result()
}

//--------------------------------------------------------------------------------------------------
//geo ops

func (ops *GoRedisClusterOps) GeoAddArr(key string, geoLocation ...GeoPosition) int64 {
	var geoList = make([]*redis.GeoLocation, 0)
	for _, x := range geoLocation {
		geoEle := redis.GeoLocation{
			Longitude: x.Longitude,
			Latitude:  x.Latitude,
			Name:      x.Member,
		}
		geoList = append(geoList, &geoEle)
	}
	return ops.client.GeoAdd(ctx, key, geoList...).Val()
}

func (ops *GoRedisClusterOps) GeoPos(key string, members ...string) (error, []GeoPosition) {
	resList := ops.client.GeoPos(ctx, key, members...)
	if len(resList.Val()) == 0 {
		return errors.New("not find any geo info"), make([]GeoPosition, 0)
	}
	resGeoList := make([]GeoPosition, 0)
	resListVal := resList.Val()
	for i, x := range members {
		resValEle := resListVal[i]
		if resValEle != nil {
			resGeoList = append(resGeoList, GeoPosition{Longitude: resValEle.Longitude, Latitude: resValEle.Latitude, Member: x})
		}
	}
	return nil, resGeoList
}

func (ops *GoRedisClusterOps) GeoDist(key string, member1, member2 string, unit GeoUnit) (error, GeoDistInfo) {
	unitStr := getUnit(unit)
	if unitStr == "" {
		return errors.New("error unit"), GeoDistInfo{}
	}
	res := ops.client.GeoDist(ctx, key, member1, member2, unitStr).Val()
	return nil, GeoDistInfo{Unit: unit, Dist: res}
}

func (ops *GoRedisClusterOps) GeoRadius(key string, query GeoRadiusQuery) (error, []GeoPosition) {
	unitStr := getUnit(query.Unit)
	if unitStr == "" {
		return errors.New("error unit"), make([]GeoPosition, 0)
	}
	res := ops.client.GeoRadius(ctx, key, query.Longitude, query.Latitude, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        unitStr,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        GetSort(query.Sort),
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	})
	geoList := make([]GeoPosition, 0)
	for _, x := range res.Val() {
		geoList = append(geoList, GeoPosition{
			Member:    x.Name,
			Longitude: x.Longitude,
			Latitude:  x.Latitude,
			Dist:      x.Dist,
			GeoHash:   x.GeoHash,
			Unit:      query.Unit,
		})
	}
	return nil, geoList
}
func (ops *GoRedisClusterOps) GeoRadiusByMember(key string, member string, query GeoRadiusByMemberQuery) (error, []GeoPosition) {

	unitStr := getUnit(query.Unit)
	if unitStr == "" {
		return errors.New("error unit"), make([]GeoPosition, 0)
	}
	res := ops.client.GeoRadiusByMember(ctx, key, member, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        unitStr,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        GetSort(query.Sort),
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	})
	geoList := make([]GeoPosition, 0)
	for _, x := range res.Val() {
		geoList = append(geoList, GeoPosition{
			Member:    x.Name,
			Longitude: x.Longitude,
			Latitude:  x.Latitude,
			Dist:      x.Dist,
			GeoHash:   x.GeoHash,
		})
	}
	return nil, geoList
}

func (ops *GoRedisClusterOps) LRange(key string, start int64, end int64) ([]string, error) {
	return ops.client.LRange(ctx, key, start, end).Result()
}

func (ops *GoRedisClusterOps) LTrim(key string, start int64, end int64) error {
	return ops.client.LTrim(ctx, key, start, end).Err()
}

func (ops *GoRedisClusterOps) RPop(key string) (string, error) {
	return ops.client.RPop(ctx, key).Result()
}

func (ops *GoRedisClusterOps) RPush(key string, values ...interface{}) (int64, error) {
	return ops.client.RPush(ctx, key, values...).Result()
}

func (ops *GoRedisClusterOps) LSet(key string, index int64, value interface{}) error {
	return ops.client.LSet(ctx, key, index, value).Err()
}

func (ops *GoRedisClusterOps) LSize(key string) (int64, error) {
	return ops.client.LLen(ctx, key).Result()
}

func (ops *GoRedisClusterOps) LRemove(key string, count int64, value interface{}) (int64, error) {
	return ops.client.LRem(ctx, key, count, value).Result()
}

func (ops *GoRedisClusterOps) SAdd(key string, members ...interface{}) (int64, error) {
	return ops.client.SAdd(ctx, key, members...).Result()
}

func (ops *GoRedisClusterOps) SDiff(keys ...string) ([]string, error) {
	return ops.client.SDiff(ctx, keys...).Result()
}

func (ops *GoRedisClusterOps) SCard(key string) (int64, error) {
	return ops.client.SCard(ctx, key).Result()
}

func (ops *GoRedisClusterOps) SInter(keys ...string) ([]string, error) {
	return ops.client.SInter(ctx, keys...).Result()
}

func (ops *GoRedisClusterOps) SInterStore(destination string, keys ...string) (int64, error) {
	return ops.client.SInterStore(ctx, destination, keys...).Result()
}

func (ops *GoRedisClusterOps) SIsMember(key string, member interface{}) (bool, error) {
	return ops.client.SIsMember(ctx, key, member).Result()
}

func (ops *GoRedisClusterOps) SMembers(key string) ([]string, error) {
	return ops.client.SMembers(ctx, key).Result()
}

func (ops *GoRedisClusterOps) SMove(source string, destination string, member interface{}) (bool, error) {
	return ops.client.SMove(ctx, source, destination, member).Result()
}

func (ops *GoRedisClusterOps) SPop(key string) (string, error) {
	return ops.client.SPop(ctx, key).Result()
}

func (ops *GoRedisClusterOps) SRandMembers(key string, count int64) ([]string, error) {
	return ops.client.SRandMemberN(ctx, key, count).Result()
}

func (ops *GoRedisClusterOps) SRem(key string, members ...interface{}) (int64, error) {
	return ops.client.SRem(ctx, key, members...).Result()
}

func (ops *GoRedisClusterOps) SUnion(keys ...string) ([]string, error) {
	return ops.client.SUnion(ctx, keys...).Result()
}

func (ops *GoRedisClusterOps) SUnionStore(destination string, keys ...string) (int64, error) {
	return ops.client.SUnionStore(ctx, destination, keys...).Result()
}

func (ops *GoRedisClusterOps) HDel(key string, fields ...string) (int64, error) {
	return ops.client.HDel(ctx, key, fields...).Result()
}

func (ops *GoRedisClusterOps) HExists(key string, field string) (bool, error) {
	return ops.client.HExists(ctx, key, field).Result()
}

func (ops *GoRedisClusterOps) HGet(key string, field string) (string, error) {
	return ops.client.HGet(ctx, key, field).Result()
}

func (ops *GoRedisClusterOps) HGetInt64(key string, field string) (int64, error) {
	return ops.client.HGet(ctx, key, field).Int64()
}

func (ops *GoRedisClusterOps) HGetFloat64(key string, field string) (float64, error) {
	return ops.client.HGet(ctx, key, field).Float64()
}

func (ops *GoRedisClusterOps) HGetAll(key string) (map[string]string, error) {
	return ops.client.HGetAll(ctx, key).Result()
}

func (ops *GoRedisClusterOps) HIncrBy(key string, field string, increment int64) (int64, error) {
	return ops.client.HIncrBy(ctx, key, field, increment).Result()
}

func (ops *GoRedisClusterOps) HKeys(key string) ([]string, error) {
	return ops.client.HKeys(ctx, key).Result()
}

func (ops *GoRedisClusterOps) HLen(key string) (int64, error) {
	return ops.client.HLen(ctx, key).Result()
}

func (ops *GoRedisClusterOps) HMGet(key string, fields ...string) ([]interface{}, error) {
	return ops.client.HMGet(ctx, key, fields...).Result()
}

func (ops *GoRedisClusterOps) HSet(key string, field string, value interface{}) (int64, error) {
	return ops.client.HSet(ctx, key, field, value).Result()
}

func (ops *GoRedisClusterOps) HSetNX(key string, field string, value interface{}) (bool, error) {
	return ops.client.HSetNX(ctx, key, field, value).Result()
}

func (ops *GoRedisClusterOps) HVals(key string) ([]string, error) {
	return ops.client.HVals(ctx, key).Result()
}
