package redis

import (
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/pool"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
	"github.com/yoyofx/yoyogo/pkg/datasources"
	"sync"
	"time"
)

// DataSourceConfig 数据源配置
type redisConfig struct {
	Name     string                      `mapstructure:"name" config:"name"`
	Url      string                      `mapstructure:"url" config:"url"`
	Password string                      `mapstructure:"password" config:"password"`
	DB       int                         `mapstructure:"db" config:"db"`
	Pool     *datasources.DataSourcePool `mapstructure:"pool" config:"pool"`
}

// DataSourcePool 数据源连接池配置
type redisPool struct {
	InitCap     int `mapstructure:"init_cap" config:"init_cap"`
	MaxCap      int `mapstructure:"max_cap" config:"max_cap"`
	Idletimeout int `mapstructure:"idle_timeout" config:"idle_timeout"`
}

type RedisDataSource struct {
	name             string
	config           abstractions.IConfiguration
	connectionString string
	connPool         map[string]pool.Pool
	count            int
	lock             sync.Mutex
	log              xlog.ILogger
}

func NewRedisClient(source *RedisDataSource) redis.IClient {
	conn, _, _ := source.Open()
	client := conn.(redis.IClient)
	return client
}

// NewMysqlDataSource 初始化MySQL数据源
func NewRedis(configuration abstractions.IConfiguration) *RedisDataSource {
	redisConfigSection := configuration.GetSection("yoyogo.datasource.redis")
	poolConfig := configuration.GetSection("yoyogo.datasource.pool")
	var redisdatasourcesConfig redisConfig
	redisConfigSection.Unmarshal(&redisdatasourcesConfig)
	poolConfig.Unmarshal(&redisdatasourcesConfig.Pool)
	log := xlog.GetXLogger("RedisDataSource")

	p := createReidsPool(redisdatasourcesConfig, log)

	dataSource := &RedisDataSource{
		name:             redisdatasourcesConfig.Name,
		connectionString: redisdatasourcesConfig.Url,
		config:           configuration,
		connPool:         make(map[string]pool.Pool, 0),
		log:              log,
	}
	if p != nil {
		dataSource.insertPool(redisdatasourcesConfig.Name, p)
	}
	return dataSource
}

func (datasource *RedisDataSource) GetName() string {
	return datasource.name
}

func (datasource *RedisDataSource) Open() (conn interface{}, put func(), err error) {

	if _, ok := datasource.connPool[datasource.name]; !ok {
		return nil, put, errors.New("no redis connect")
	}

	conn, err = datasource.connPool[datasource.name].Get()
	if err != nil {
		return nil, put, errors.New(fmt.Sprintf("redis get connect err:%v", err))
	}

	put = func() {
		_ = datasource.connPool[datasource.name].Put(conn)
	}

	return conn, put, nil
}

func (datasource *RedisDataSource) Close() {
	//panic("implement me")
}

func (datasource *RedisDataSource) Ping() bool {
	conn, put, err := datasource.Open()
	if err != nil {
		return false
	}
	defer put()
	ret := datasource.connPool[datasource.name].Ping(conn) == nil
	return ret
}

func (datasource *RedisDataSource) GetConnectionString() string {
	return datasource.connectionString
}

// insertPool 将连接池插入map,支持多个不同mysql链接
func (datasource *RedisDataSource) insertPool(name string, p pool.Pool) {
	if datasource.connPool == nil {
		datasource.connPool = make(map[string]pool.Pool, 0)
	}
	datasource.lock.Lock()
	defer datasource.lock.Unlock()
	datasource.connPool[name] = p
}

func createReidsPool(redisdatasourcesConfig redisConfig, log xlog.ILogger) pool.Pool {
	if redisdatasourcesConfig.Pool != nil && (redisdatasourcesConfig.Pool.InitCap == 0 || redisdatasourcesConfig.Pool.MaxCap == 0 || redisdatasourcesConfig.Pool.Idletimeout == 0) {
		log.Error("redis config is error initCap,maxCap,idleTimeout should be gt 0")
		return nil
	}

	// connRedis 建立连接
	connRedis := func() (interface{}, error) {
		options := &redis.Options{}

		options.Addr = redisdatasourcesConfig.Url

		if redisdatasourcesConfig.Password != "" {
			options.Password = redisdatasourcesConfig.Password
		}
		if redisdatasourcesConfig.DB > 0 {
			options.DB = redisdatasourcesConfig.DB
		}

		return redis.NewClient(options), nil
	}

	// closeRedis 关闭连接
	closeRedis := func(v interface{}) error {
		return v.(redis.IClient).Close()
	}

	// pingRedis 检测连接连通性
	pingRedis := func(v interface{}) error {
		conn := v.(redis.IClient)

		val, err := conn.Ping()

		if err != nil {
			return err
		}
		if val != "PONG" {
			return errors.New("redis ping is error ping => " + val)
		}

		return nil
	}

	//创建一个连接池： 初始化5，最大连接30
	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap: redisdatasourcesConfig.Pool.InitCap,
		MaxCap:     redisdatasourcesConfig.Pool.MaxCap,
		Factory:    connRedis,
		Close:      closeRedis,
		Ping:       pingRedis,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: time.Duration(redisdatasourcesConfig.Pool.Idletimeout) * time.Second,
	})
	if err != nil {
		log.Error("register redis conn [%s] error:%v", redisdatasourcesConfig.Name, err)
		return nil
	}

	return p
}
