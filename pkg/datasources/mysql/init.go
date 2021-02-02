package mysql

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/health"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func init() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplementsAndName("mysql-master", NewMysqlDataSource, new(abstractions.IDataSource))
			serviceCollection.AddTransient(NewGormDb)
			serviceCollection.AddTransientByImplements(NewMysqlHealthIndicator, new(health.Indicator))
		})

}
