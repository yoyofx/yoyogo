package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/console"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/pkg/scheduler"
)

func main() {
	// -f ./conf/test_conf.yml 指定配置文件 , 默认读取 config_{profile}.yml , -profile [dev,test,prod]
	config := abstractions.NewConfigurationBuilder().AddYamlFile("config").Build()

	console.NewHostBuilder().
		UseConfiguration(config).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {

			scheduler.UseXxlJob(collection)

			scheduler.AddJobs(collection, NewDemoJob)
		}).
		Build().Run()
}
