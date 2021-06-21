package main

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
