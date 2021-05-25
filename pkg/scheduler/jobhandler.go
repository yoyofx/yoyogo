package scheduler

type JobHandler interface {
	// Execute 任务的执行函数
	Execute(cxt *JobContext) (msg string)
	// GetJobName 自定义任务的名字
	GetJobName() string
}
