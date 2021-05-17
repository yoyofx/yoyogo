package scheduler

import (
	"context"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/yoyofx/yoyogo/web"
	"time"
)

type JobHandler interface {
	//任务的执行函数
	Execute(cxt context.Context, param *xxl.RunReq) (msg string)
	//自定义任务的名字
	GetJobName() string
}

type ExecutorOptionsBuilder func() ExecutorOptions

type JobRegister struct {
	Application *web.ApplicationBuilder
	Executor    xxl.Executor
}

type ExecutorOptions struct {
	ServerAddr   string
	AccessToken  string
	Timeout      time.Duration
	ExecutorIp   string
	ExecutorPort string
	RegistryKey  string
	//LogDir       string
	logger xxl.Logger //日志处理
}

/**
构造执行器
*/
func (ops *ExecutorOptions) BuildExecutor() xxl.Executor {
	opsMap := make([]xxl.Option, 0)
	if ops.ServerAddr != "" {
		opsMap = append(opsMap, xxl.ServerAddr(ops.ServerAddr))
	}
	if ops.AccessToken != "" {
		opsMap = append(opsMap, xxl.AccessToken(ops.AccessToken))
	}
	if ops.ExecutorIp != "" {
		opsMap = append(opsMap, xxl.ExecutorIp(ops.ExecutorIp))
	}
	if ops.ExecutorPort != "" {
		opsMap = append(opsMap, xxl.ExecutorPort(ops.ExecutorPort))
	}
	if ops.RegistryKey != "" {
		opsMap = append(opsMap, xxl.RegistryKey(ops.RegistryKey))
	}
	/*if ops.LogDir!="" {
		opsMap=append(opsMap,xxl.ops.LogDir))
	}*/
	if ops.logger != nil {
		opsMap = append(opsMap, xxl.SetLogger(ops.logger))
	}
	return xxl.NewExecutor(opsMap...)
}

//RegisterJob将不再返回application,确保注册任务是最后一步执行
func (reg *JobRegister) RegisterJob(jobList []JobHandler) {
	reg.Executor.Init()
	if len(jobList) > 0 {
		for _, x := range jobList {
			reg.Executor.RegTask(x.GetJobName(), x.Execute)
		}
	}
	reg.Executor.Run()
}
