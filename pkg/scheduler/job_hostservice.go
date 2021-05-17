package scheduler

import (
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofx/yoyogo/dependencyinjection"
)

func UseXxlJob(collection *dependencyinjection.ServiceCollection) {
	hosting.AddHostService(collection, NewXxlJobService)
}

func AddJob(collection *dependencyinjection.ServiceCollection, jobCtor interface{}) {
	collection.AddSingletonByImplements(jobCtor, new(JobHandler))
}

func AddJobs(collection *dependencyinjection.ServiceCollection, jobListCtor ...interface{}) {
	for _, jobCtor := range jobListCtor {
		collection.AddSingletonByImplements(jobCtor, new(JobHandler))
	}
}

// XxlJobService as IHostService
type XxlJobService struct {
	Executor xxl.Executor
	jobList  []JobHandler
}

func NewXxlJobService(configuration abstractions.IConfiguration, environment *abstractions.HostEnvironment, jobList []JobHandler) *XxlJobService {
	xxlSection := configuration.GetSection("yoyogo.application.xxl")
	var ops *ExecutorOptions
	xxlSection.Unmarshal(&ops)
	ops.RegistryKey = environment.ApplicationName
	if ops.ExecutorIp == "" {
		ops.ExecutorIp = environment.Host
	}

	service := &XxlJobService{Executor: ops.BuildExecutor(), jobList: jobList}
	return service
}

func (service *XxlJobService) Run() error {
	service.Executor.Init()
	service.registerJob(service.jobList...)
	return service.Executor.Run()
}

func (service *XxlJobService) Stop() error {
	service.Executor.Stop()
	return nil
}

// RegisterJob 将不再返回application,确保注册任务是最后一步执行
func (service *XxlJobService) registerJob(jobList ...JobHandler) {
	if len(jobList) > 0 {
		for _, x := range jobList {
			service.Executor.RegTask(x.GetJobName(), x.Execute)
		}
	}
}
