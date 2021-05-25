package scheduler

import (
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
	"sync"
)

type JobContext struct {
	*xxl.RunReq
	Logger  Logger
	Context context.Context
}

var (
	jobDoneList = sync.Map{}
)

func GetContext(ctx context.Context, request *xxl.RunReq) *JobContext {
	logger, err := NewXxlJobLogger(request.LogID)
	if err != nil {
		panic(err)
	}
	return &JobContext{
		request,
		logger,
		ctx,
	}
}

// Report 上报日志
func (ctx *JobContext) Report(format string, args ...interface{}) {
	ctx.Logger.Info(format, args)
}

// Error 任务发生错误
func (ctx *JobContext) Error(args ...interface{}) {
	ctx.Logger.Error(args...)
}

// Done 任务完成返回
func (ctx *JobContext) Done(msg string) string {
	jobDoneList.Store(ctx.LogID, true)
	return ctx.Logger.Done(msg)
}

func IsDone(logID int64) bool {
	val, ok := jobDoneList.Load(logID)
	if ok {
		return val.(bool)
	}
	return ok
}
