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

}
