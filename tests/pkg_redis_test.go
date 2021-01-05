package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
	"strconv"
	"testing"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "62.234.6.120:31379",
	Password: "",
	DB:       0,
})

func Test_RedisValueOps(t *testing.T) {
	client.SetSerializer(&redis.JsonSerializer{})
	pong, _ := client.Ping()
	assert.Equal(t, pong, "PONG")

	dckey := "dcctor1"
	doctor := &Doctor{Name: "钟南山", Age: 85}
	hask := client.HasKey(dckey)
	b, _ := client.GetKVOps().SetIfAbsent(dckey, doctor)
	assert.Equal(t, hask, !b)
	var doc Doctor
	_ = client.GetKVOps().Get(dckey, &doc)
	assert.Equal(t, doc.Age, 85)

	_ = client.GetKVOps().SetString("say1", "hello", 3*time.Hour)
	client.GetKVOps().Append("say1", " world")
	s, _ := client.GetKVOps().GetString("say1")
	assert.Equal(t, s, "hello world")
	incr := 0
	add := 2 // -1
	strincr, _ := client.GetKVOps().GetString("test_incr")
	incr, _ = strconv.Atoi(strincr)

	tincr, _ := client.GetKVOps().Increment("test_incr", int64(add))
	assert.Equal(t, int64(incr+add), tincr)

	sete, _ := client.SetExpire("test_incr", 3*time.Hour)
	assert.Equal(t, sete, true)
	t1, _ := client.GetExpire("test_incr")
	assert.Equal(t, int64(t1) > 0, true)
	s1, _ := client.RandomKey()
	fmt.Println(s1)
}

var geoKey = "Geo"

func TestRedisGeo(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	res := ops.GeoAdd(geoKey, 116.488566, 39.914741, "GUOMAO")
	fmt.Println(res)
}

func TestRedisGeoAdd(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	list := make([]redis.GeoPosition, 0)
	list = append(list, redis.GeoPosition{Member: "北京东", Longitude: 116.49065, Latitude: 39.908294})
	list = append(list, redis.GeoPosition{Member: "慈云寺", Longitude: 116.495429, Latitude: 39.919307})
	res := ops.GeoAddArr(geoKey, list)
	fmt.Println(res)
}
func TestRedisGeoPos(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR, res := ops.GeoPos(geoKey, "GUOMAO")
	assert.Equal(t, ERR, nil)
	assert.Equal(t, res.Member, "GUOMAO")
}

func TestGeoDist(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR, res := ops.GeoDist(geoKey, "GUOMAO", "SIHUI", redis.M)
	assert.Equal(t, ERR, nil)
	assert.Equal(t, res.Dist, float64(1128.2414))
}

func TestGeoRadius(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR, res := ops.GeoRadius(geoKey, redis.GeoRadiusQuery{Longitude: 116.514724, Latitude: 39.922378, Radius: 10, Unit: redis.KM, WithDist: true, Count: 5, WithCoord: true})
	assert.Equal(t, ERR, nil)
	assert.Equal(t, len(res), 5)
}

func TestGeoRadiusByMember(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR, res := ops.GeoRadiusByMember(geoKey, redis.GeoRadiusByMemberQuery{Member: "SIHUI", Radius: 10, Unit: redis.KM, WithDist: true, Count: 3, WithCoord: true})
	assert.Equal(t, ERR, nil)
	assert.Equal(t, len(res), 3)
}
