package console

import (
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/projects"
)

var Project = projects.NewEmptyProject("console", "Console Application").With(func(root *projects.ProjectItem) {
	root.AddFileWithContent("main.go", ProjectItem_Main_go)
	root.AddFileWithContent("startup.go", ProjectItem_startup_go)
	root.AddFileWithContent("hostservice.go", ProjectItem_hostservice_go)
	root.AddFileWithContent("config.yml", ProjectItem_conf_yml)
	root.AddFileWithContent("go.mod", ProjectItem_go_mod)
})
