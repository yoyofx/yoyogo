package xxl_job

const Main_Tel = `
package {{.CurrentModelName}}

import (
	"github.com/yoyofx/yoyogo/abstractions/configuration"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/pkg/scheduler"
)

func main() {
	// -f ./conf/test_conf.yml 指定配置文件 , 默认读取 config_{profile}.yml , -profile [dev,test,prod]
	config := configuration.YAML("config")
	
	scheduler.NewXxlJobBuilder(config).
		ConfigureServices(func(collection *dependencyinjection.ServiceCollection) {
			scheduler.AddJobs(collection, NewDemoJob)
		}).
		Build().Run()
}

`

const Mod_Tel = `
module "{{.ModelName}}"

go 1.16

require (
    github.com/yoyofx/yoyogo {{.Version}}
	github.com/xxl-job/xxl-job-executor-go v0.6.1
)

`

const Demo_Job_Tel = `
package {{.CurrentModelName}}

import (
	"github.com/yoyofx/yoyogo/pkg/scheduler"
	"time"
)

type DemoJob struct {
}

func NewDemoJob() *DemoJob {
	return &DemoJob{}
}

func (*DemoJob) Execute(cxt *scheduler.JobContext) (msg string) {
	cxt.Report("Job %d is beginning...", cxt.LogID)

	for i := 1; i <= 100; i++ {
		cxt.Report("Job Progress: %d Percent.", i)
		time.Sleep(time.Second)
	}

	return cxt.Done("666")
}

//GetJobName 自定义任务的名字
func (*DemoJob) GetJobName() string {
	return "job1"
}
`

const Config_Tel = `
yoyogo:
  application:
    name: console-xxl-job
    metadata: "dev"
    server:
      type: "console"
    xxl:
      serverAddr: http://127.0.0.1:8080/xxl-job-admin/
      #ip: ""
      port: 9999
      #accessToken: ""
      #timeout: 0

`
