package models

import (
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/yoyofx/yoyogo/pkg/scheduler"
)

func BuildExecutor() scheduler.ExecutorOptionsBuilder  {
	return func() scheduler.ExecutorOptions {
		return scheduler.ExecutorOptions{
			ExecutorIp: "127.0.0.1",
			ExecutorPort: "5000",
		}
	}
}

func BuildJobList()[]scheduler.JobHandler  {
	jobMap:=make([]scheduler.JobHandler,1)
	jobMap=append(jobMap,&DemoJob{})
	jobMap=append(jobMap,&DemoJob2{})
	return jobMap
}


type DemoJob struct {

}

func (*DemoJob)   Execute(cxt context.Context, param *xxl.RunReq) (msg string){

	return "666"
}


//自定义任务的名字
func (*DemoJob) GetJobName() string{
	return "job1"
}

type  DemoJob2 struct {

}

func (*DemoJob2)   Execute(cxt context.Context, param *xxl.RunReq) (msg string){

	return "666"
}


//自定义任务的名字
func (*DemoJob2) GetJobName() string{
	return "job1"
}