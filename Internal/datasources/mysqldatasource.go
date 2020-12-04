package datasources

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/Pool"
	"sync"
)

// DataSourceConfig 数据源配置
type dataSourceConfig struct {
	Name     string         `mapstructure:"name"`
	Url      string         `mapstructure:"url"`
	UserName string         `mapstructure:"username"`
	Password string         `mapstructure:"password"`
	Debug    bool           `mapstructure:"debug"`
	Pool     dataSourcePool `mapstructure:"pool"`
}

// DataSourcePool 数据源连接池配置
type dataSourcePool struct {
	InitCap     uint16 `mapstructure:"init_cap"`
	MaxCap      uint16 `mapstructure:"max_cap"`
	Idletimeout uint16 `mapstructure:"idle_timeout"`
}

type MySqlDataSource struct {
	name     string
	config   Abstractions.IConfiguration
	connPool map[string]Pool.Pool
	count    int
	lock     sync.Mutex
}

// NewMysqlDataSource 初始化MySQL数据源
func NewMysqlDataSource(configuration Abstractions.IConfiguration) *MySqlDataSource {
	databaseConfig := configuration.GetSection("yoyogo.datasource.mysql")
	var datasources dataSourceConfig
	databaseConfig.Unmarshal(&datasources)

	return &MySqlDataSource{
		name:     datasources.Name,
		config:   configuration,
		connPool: make(map[string]Pool.Pool, 0),
	}
}

func (datasource *MySqlDataSource) GetName() string {
	return datasource.name
}

func (datasource *MySqlDataSource) Open() interface{} {
	panic("implement me")
}

func (datasource *MySqlDataSource) Close() {
	panic("implement me")
}

func (datasource *MySqlDataSource) Ping() bool {
	return true
}

// insertPool 将连接池插入map,支持多个不同mysql链接
func (datasource *MySqlDataSource) insertPool(name string, p Pool.Pool) {
	if datasource.connPool == nil {
		datasource.connPool = make(map[string]Pool.Pool, 0)
	}
	datasource.lock.Lock()
	defer datasource.lock.Unlock()
	datasource.connPool[name] = p
}
