package scheduler

import (
	"context"
	xxl "github.com/xxl-job/xxl-job-executor-go"
)

type JobHandler interface {
	// Execute 任务的执行函数
	Execute(cxt context.Context, param *xxl.RunReq) (msg string)
	// GetJobName 自定义任务的名字
	GetJobName() string
}
