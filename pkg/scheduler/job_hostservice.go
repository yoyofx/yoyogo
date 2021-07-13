package scheduler

import (
	"bufio"
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/hosting"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"github.com/yoyofx/yoyogo/utils"
	"github.com/yoyofxteam/dependencyinjection"
	"strings"
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
	fmt.Println(consolecolors.Green("xxl-job executor is running..."))
	service.Executor.Init()
	//设置日志查看handler
	service.Executor.LogHandler(service.viewLogs)
	service.registerJob(service.jobList...)
	return service.Executor.Run()
}

func (service *XxlJobService) Stop() error {
	service.Executor.Stop()
	return nil
}

func (service *XxlJobService) viewLogs(req *xxl.LogReq) *xxl.LogRes {
	jobDone := IsDone(req.LogID)
	fi, err := getLogFile(req.LogID, false)

	if err != nil {
		return nil
	}
	defer fi.Close()

	var lines []string
	rowCount := 0
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		rowCount++
		line := scanner.Text()
		if !jobDone {
			if strings.Contains(line, "xxljob-done") {
				jobDone = true
			}
		}

		if rowCount >= req.FromLineNum {
			lines = append(lines, scanner.Text())
		}
	}

	return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   rowCount,
		LogContent:  strings.Join(lines, utils.NewLine()),
		IsEnd:       jobDone,
	}}
}

// RegisterJob 将不再返回application,确保注册任务是最后一步执行
func (service *XxlJobService) registerJob(jobList ...JobHandler) {
	if len(jobList) > 0 {
		for _, x := range jobList {
			service.Executor.RegTask(x.GetJobName(), func(cxt context.Context, param *xxl.RunReq) string {
				jobContext := GetContext(cxt, param)
				return x.Execute(jobContext)
			})
		}
	}
}
