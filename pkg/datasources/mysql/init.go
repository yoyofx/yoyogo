package mysql

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("db1", NewMysqlDataSource, new(abstractions.IDataSource))
		})
}
