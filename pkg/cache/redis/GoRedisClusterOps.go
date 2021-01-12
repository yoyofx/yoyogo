package redis

import (
	"github.com/go-redis/redis/v8"
)

type GoRedisClusterOps struct {
	GoRedisStandaloneOps
}

func NewClusterOps(options *Options) *GoRedisClusterOps {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    options.Addrs,
		Password: options.Password,
	})
	return &GoRedisClusterOps{GoRedisStandaloneOps{client: client}}
}
