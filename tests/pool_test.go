package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/abstractions/pool"
	"net"
	"testing"
	"time"
)

func Test_Pool(t *testing.T) {
	addr := ":18087"
	//创建一个连接池： 初始化5，最大连接30
	poolConfig := &pool.Config{
		InitialCap: 5,
		MaxCap:     30,
		Factory:    func() (interface{}, error) { return net.Dial("udp", addr) },
		Close:      func(v interface{}) error { return v.(net.Conn).Close() },
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	p, err := pool.NewChannelPool(poolConfig)
	assert.Equal(t, err, nil)
	//从连接池中取得一个连接
	v, err := p.Get()
	//do something
	//conn=v.(net.Conn)
	//将连接放回连接池中
	current := p.Len()
	fmt.Println("len=", current)
	_ = p.Put(v)
	//释放连接池中的所有连接
	//p.Release()
	assert.Equal(t, current, 4)
	//查看当前连接中的数量
	current = p.Len()
	fmt.Println("len=", current)
	assert.Equal(t, current, 5)
}
