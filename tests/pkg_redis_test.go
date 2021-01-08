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

func Test_RedisListOps(t *testing.T) {
	key := "go2list"
	client.SetSerializer(&redis.JsonSerializer{})
	listOps := client.GetListOps()
	_, _ = listOps.Clear(key)
	_ = listOps.AddElements(key, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	var list []int
	listOps.GetElements(key, 0, 2, &list)
	assert.Equal(t, len(list) > 0, true)
	var a1 int
	_ = listOps.GetElement(key, 0, &a1)
	assert.Equal(t, a1, 1)
	size, _ := listOps.Size(key)
	assert.Equal(t, size, int64(9))
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

func initGeoPosition() {
	list := make([]redis.GeoPosition, 0)
	list = append(list, redis.GeoPosition{Member: "北京东", Longitude: 116.49065, Latitude: 39.908294})
	list = append(list, redis.GeoPosition{Member: "慈云寺", Longitude: 116.495429, Latitude: 39.919307})
	list = append(list, redis.GeoPosition{Member: "四惠", Longitude: 116.49546146, Latitude: 39.90874867})
	list = append(list, redis.GeoPosition{Member: "八里庄", Longitude: 116.49889469, Latitude: 39.91773496})
	list = append(list, redis.GeoPosition{Member: "国贸", Longitude: 116.46190166, Latitude: 39.9091437})
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ops.GeoAddArr(geoKey, list)
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
	list = append(list, redis.GeoPosition{Member: "四惠", Longitude: 116.49546146, Latitude: 39.90874867})
	list = append(list, redis.GeoPosition{Member: "八里庄", Longitude: 116.49889469, Latitude: 39.91773496})
	list = append(list, redis.GeoPosition{Member: "国贸", Longitude: 116.46190166, Latitude: 39.9091437})
	res := ops.GeoAddArr(geoKey, list)
	fmt.Println(res)
}
func TestRedisGeoPos(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	initGeoPosition()
	ops := client.GetGeoOps()
	ERR, res := ops.GeoPos(geoKey, "四惠")
	assert.Equal(t, ERR, nil)
	assert.Equal(t, res.Member, "四惠")
}

func TestGeoDist(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	initGeoPosition()
	ops := client.GetGeoOps()
	ERR, res := ops.GeoDist(geoKey, "国贸", "四惠", redis.M)
	assert.Equal(t, ERR, nil)
	assert.Equal(t, res.Dist, 2863.586)
}

func TestGeoRadius(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	initGeoPosition()
	ops := client.GetGeoOps()
	ERR, res := ops.GeoRadius(geoKey, redis.GeoRadiusQuery{Longitude: 116.49546146, Latitude: 39.90874867, Radius: 10, Unit: redis.KM, WithDist: true, Count: 5, WithCoord: true})
	assert.Equal(t, ERR, nil)
	assert.Equal(t, len(res), 5)
}

func TestGeoRadiusByMember(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	initGeoPosition()
	ops := client.GetGeoOps()
	ERR, res := ops.GeoRadiusByMember(geoKey, redis.GeoRadiusByMemberQuery{Member: "四惠", Radius: 10, Unit: redis.KM, WithDist: true, Count: 3, WithCoord: true})
	assert.Equal(t, ERR, nil)
	assert.Equal(t, len(res), 3)
}

func TestGetLock(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	var key = "dlock"
	ops := client.GetLockOps()
	ops.DisposeLock(key)
	err, success := ops.GetDLock(key, 5)
	assert.Equal(t, err, nil)
	assert.Equal(t, success, true)
}

func TestHashSet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	client.SetSerializer(&redis.JsonSerializer{})
	key := "hashset:app"
	client.Delete(key)
	hashOps := client.GetHashOps()
	err := hashOps.Put(key, "a1", 1)
	assert.Equal(t, err, nil)
	b, _ := hashOps.PutIfAbsent(key, "a1", 2)
	assert.Equal(t, b, false)
	a1Value := hashOps.Increment(key, "a1", 1)
	assert.Equal(t, a1Value, int64(2))

	doctor := &Doctor{Name: "hash_doctor", Age: 0}
	hasDoc1, _ := hashOps.PutIfAbsent(key, "doc1", doctor)
	assert.Equal(t, hasDoc1, true)
	var a1 int64
	err = hashOps.Get(key, "a1", &a1)
	assert.Equal(t, err, nil)
	assert.Equal(t, a1, int64(2))
	var doctor1 Doctor
	_ = hashOps.Get(key, "doc1", &doctor1)
	assert.Equal(t, doctor1.Name, "hash_doctor")
	keys, _ := hashOps.GetHashKeys(key)
	assert.Equal(t, int64(len(keys)), hashOps.Size(key))

	assert.Equal(t, hashOps.Exists(key, "doc1"), true)

	_ = hashOps.Put(key, "doc2", "my dcoker")
	str, _ := hashOps.GetString(key, "doc2")
	assert.Equal(t, str, "my dcoker")
	m, _ := hashOps.GetEntries(key)
	assert.Equal(t, m != nil, true)
}
