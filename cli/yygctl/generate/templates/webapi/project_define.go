package webapi

import "github.com/yoyofx/yoyogo/cli/yygctl/generate/projects"

var Project = projects.NewEmptyProject("console", "Console Application").With(func(root *projects.ProjectItem) {
	root.AddFileWithContent("main.go", Main_tel)
	root.AddFileWithContent("go.mod", Mod_tel)
})
