package console

const ProjectItem_hostservice_go = `
package main

import "fmt"

type HostService struct {
}

func NewService() *HostService {
	return &Service1{}
}

func (s *HostService) Run() error {
	fmt.Println("host service Running")
	return nil
}

func (s *HostService) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}
`
