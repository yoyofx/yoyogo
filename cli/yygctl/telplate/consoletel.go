package telplate

const ConsoleMainTel = `
package {{.ModelName}}
import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/console"
)

func main() {
	// -f ./conf/test_conf.yml 指定配置文件 , 默认读取 config_{profile}.yml , -profile [dev,test,prod]
	config := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	console.NewHostBuilder().
		UseConfiguration(config).
		UseStartup(Startup).
		Build().
		Run()
}
`

const ConsoleStartUpTel = `
package {{.ModelName}}

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

type AppStartup struct {
}

func Startup() abstractions.IStartup {
	return &AppStartup{}
}

func (s *AppStartup) ConfigureServices(collection *dependencyinjection.ServiceCollection) {
	hosting.AddHostService(collection, NewService)
}
`

const ConsoleGoModTel = `
module {{.ModelName}}
go 1.16
require github.com/yoyofx/yoyogo v1.7.2
`

const ConsoleConfigTel = `
yoyogo:
  application:
    name: {{.ModelName}}
    metadata: "dev"
    server:
      type: "console"`

const ConsoleServiceTel = `
package {{.ModelName}}

import "fmt"

type Service1 struct {
}

func NewService() *Service1 {
	return &Service1{}
}

func (s *Service1) Run() error {
	fmt.Println("host service Running")
	return nil
}

func (s *Service1) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}

`
