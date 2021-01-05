package tests

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	cache "github.com/yoyofx/yoyogo/pkg/cache/redis"

	"testing"
	"time"
)

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	return client
}

var ctx = context.Background()

func TestRedisConn(t *testing.T) {
	client := newClient()
	//defer client.Close()

	//ping
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return
	}
	assert.Equal(t, pong, "PONG")
}

type Doctor struct {
	Name string
	Age  int
}

func TestRedisStringValue(t *testing.T) {
	client := newClient()
	defer client.Close()

	// string
	client.Set(ctx, "yoyogo:version", "v1.6.1", 15*time.Minute)
	version, _ := client.Get(ctx, "yoyogo:version").Result()
	assert.Equal(t, version, "v1.6.1")

	// json 序列化
	doctor := Doctor{Name: "钟南山", Age: 83}

	serializer := cache.JsonSerializer{}

	doctorJson, _ := serializer.Serialization(doctor)
	client.Set(ctx, "doctor2", doctorJson, time.Hour)
	var doctor2 Doctor
	doctorResult, _ := client.Get(ctx, "doctor2").Bytes()
	_ = serializer.Deserialization(doctorResult, &doctor2)
	assert.Equal(t, doctor2.Name, "钟南山")
	assert.Equal(t, doctor2.Age, 83)

	//client.SetNX()
}

func TestRedisList(t *testing.T) {
	client := newClient()
	defer client.Close()
	listKey := "go2list"
	_, _ = client.Del(ctx, listKey).Result()

	client.RPush(ctx, listKey, 1, 2, 3)

	first, _ := client.LPop(ctx, listKey).Int()
	assert.Equal(t, first, 1)

	i1, _ := client.LIndex(ctx, listKey, 1).Int64()
	assert.Equal(t, i1, int64(3))

}

var geoKey="Geo"

func TestRedisGeo(t *testing.T)  {
	client:=cache.NewClient(&cache.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	res:= ops.GeoAdd(geoKey,116.488566,39.914741,"GUOMAO")
	fmt.Println(res)
}

func TestRedisGeoAdd(t *testing.T)  {
	client:=cache.NewClient(&cache.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	list:=make([]cache.GeoPosition,0)
	list=append(list,cache.GeoPosition{Member:"北京东",Longitude: 116.49065,Latitude: 39.908294})
	list=append(list,cache.GeoPosition{Member:"慈云寺",Longitude: 116.495429,Latitude: 39.919307})
	res:= ops.GeoAddArr(geoKey,list)
	fmt.Println(res)
}
func TestRedisGeoPos(t *testing.T)  {
	client:=cache.NewClient(&cache.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR,res:= ops.GeoPos(geoKey,"GUOMAO")
	fmt.Println(ERR)
	fmt.Println(res)
}

func TestGeoDist(t *testing.T)  {
	client:=cache.NewClient(&cache.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR,res:= ops.GeoDist(geoKey,"GUOMAO","SIHUI",cache.M)
	fmt.Println(ERR)
	fmt.Println(res)
}

func TestGeoRadius(t *testing.T)  {
	client:=cache.NewClient(&cache.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR,res:= ops.GeoRadius(geoKey,cache.GeoRadiusQuery{Longitude:116.514724,Latitude: 39.922378, Radius: 10,Unit: cache.KM,WithDist: true,Count: 5,WithCoord: true})
	fmt.Println(ERR)
	fmt.Println(res)
}

func TestGeoRadiusByMember(t *testing.T)  {
	client:=cache.NewClient(&cache.Options{
		Addr:     "62.234.6.120:31379",
		Password: "",
		DB:       0,
	})
	ops := client.GetGeoOps()
	ERR,res:= ops.GeoRadiusByMember(geoKey,cache.GeoRadiusByMemberQuery{Member: "SIHUI", Radius: 10,Unit: cache.KM,WithDist: true,Count: 3,WithCoord: true})
	fmt.Println(ERR)
	fmt.Println(res)
}