package DataSources

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/Pool"
	"sync"
)

//type MySqlDataSourceConfig struct {
//	db1  `mapstructure:"type"`
//}

type MySqlDataSource struct {
	config   Abstractions.IConfiguration
	connPool map[string]Pool.Pool
	count    int
	lock     sync.Mutex
}

// NewMysqlDataSource 初始化MySQL数据源
func NewMySQLDataSource(configuration Abstractions.IConfiguration) Abstractions.IDataSource {
	ss := configuration.GetSection("yoyogo.database")
	_ = ss
	return &MySqlDataSource{
		config:   configuration,
		connPool: make(map[string]Pool.Pool, 0),
	}
}

func (datasource *MySqlDataSource) Open() interface{} {
	panic("implement me")
}

func (datasource *MySqlDataSource) Close() {
	panic("implement me")
}

func (datasource *MySqlDataSource) Ping() bool {
	panic("implement me")
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
