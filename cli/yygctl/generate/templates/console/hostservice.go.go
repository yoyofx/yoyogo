package console

const ProjectItem_hostservice_go = `
package {{.CurrentModelName}}

import "fmt"

type HostService struct {
}

func NewService() *HostService {
	return &HostService{}
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
