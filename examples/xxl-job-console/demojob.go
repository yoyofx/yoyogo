package main

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
)

type DemoJob struct {
}

func NewDemoJob() *DemoJob {
	return &DemoJob{}
}

func (*DemoJob) Execute(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("this is job1")
	return "666"
}

//自定义任务的名字
func (*DemoJob) GetJobName() string {
	return "job1"
}
