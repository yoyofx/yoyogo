package datasources

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/pool"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"sync"
	"time"
)

// DataSourceConfig 数据源配置
type dataSourceConfig struct {
	Name     string          `mapstructure:"name"`
	Url      string          `mapstructure:"url"`
	UserName string          `mapstructure:"username"`
	Password string          `mapstructure:"password"`
	Debug    bool            `mapstructure:"debug"`
	Pool     *dataSourcePool `mapstructure:"pool"`
}

// DataSourcePool 数据源连接池配置
type dataSourcePool struct {
	InitCap     int `mapstructure:"init_cap"`
	MaxCap      int `mapstructure:"max_cap"`
	Idletimeout int `mapstructure:"idle_timeout"`
}

type MySqlDataSource struct {
	name             string
	config           abstractions.IConfiguration
	connectionString string
	connPool         map[string]pool.Pool
	count            int
	lock             sync.Mutex
	log              xlog.ILogger
}

// NewMysqlDataSource 初始化MySQL数据源
func NewMysqlDataSource(configuration abstractions.IConfiguration) *MySqlDataSource {
	databaseConfig := configuration.GetSection("yoyogo.datasource.mysql")
	var datasourcesConfig dataSourceConfig
	databaseConfig.Unmarshal(&datasourcesConfig)
	log := xlog.GetXLogger("MysqlDataSource")

	if datasourcesConfig.Pool != nil && (datasourcesConfig.Pool.InitCap == 0 || datasourcesConfig.Pool.MaxCap == 0 || datasourcesConfig.Pool.Idletimeout == 0) {
		log.Error("database config is error initCap,maxCap,idleTimeout should be gt 0")
		return nil
	}

	dsnPath := fmt.Sprintf("%s:%s@%s", datasourcesConfig.UserName, datasourcesConfig.Password, datasourcesConfig.Url)

	// connMysql 建立连接
	connMysql := func() (interface{}, error) {
		conn, err := sql.Open("mysql", dsnPath)
		return conn, err
	}

	// closeMysql 关闭连接
	closeMysql := func(v interface{}) error {
		return v.(*sql.DB).Close()
	}

	// pingMysql 检测连接连通性
	pingMysql := func(v interface{}) error {
		conn := v.(*sql.DB)
		return conn.Ping()
	}

	//创建一个连接池： 初始化5，最大连接30
	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap: datasourcesConfig.Pool.InitCap,
		MaxCap:     datasourcesConfig.Pool.MaxCap,
		Factory:    connMysql,
		Close:      closeMysql,
		Ping:       pingMysql,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: time.Duration(datasourcesConfig.Pool.Idletimeout) * time.Second,
	})
	if err != nil {
		log.Error("register mysql conn [%s] error:%v", datasourcesConfig.Name, err)
		return nil
	}

	dataSource := &MySqlDataSource{
		name:             datasourcesConfig.Name,
		connectionString: dsnPath,
		config:           configuration,
		connPool:         make(map[string]pool.Pool, 0),
		log:              log,
	}

	dataSource.insertPool(datasourcesConfig.Name, p)

	return dataSource
}

func (datasource *MySqlDataSource) GetName() string {
	return datasource.name
}

func (datasource *MySqlDataSource) Open() (conn interface{}, put func(), err error) {

	if _, ok := datasource.connPool[datasource.name]; !ok {
		return nil, put, errors.New("no mysql connect")
	}

	conn, err = datasource.connPool[datasource.name].Get()
	if err != nil {
		return nil, put, errors.New(fmt.Sprintf("mysql get connect err:%v", err))
	}

	put = func() {
		_ = datasource.connPool[datasource.name].Put(conn)
	}

	return conn, put, nil
}

func (datasource *MySqlDataSource) Close() {
	//panic("implement me")
}

func (datasource *MySqlDataSource) Ping() bool {
	conn, put, err := datasource.Open()
	if err != nil {
		return false
	}
	defer put()
	ret := datasource.connPool[datasource.name].Ping(conn) == nil
	return ret
}

func (datasource *MySqlDataSource) GetConnectionString() string {
	return datasource.connectionString
}

func (datasource *MySqlDataSource) createPool() {

}

// insertPool 将连接池插入map,支持多个不同mysql链接
func (datasource *MySqlDataSource) insertPool(name string, p pool.Pool) {
	if datasource.connPool == nil {
		datasource.connPool = make(map[string]pool.Pool, 0)
	}
	datasource.lock.Lock()
	defer datasource.lock.Unlock()
	datasource.connPool[name] = p
}
