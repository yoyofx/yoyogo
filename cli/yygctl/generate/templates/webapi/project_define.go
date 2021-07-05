package webapi

import "yygctl/generate/projects"

var Project = projects.NewEmptyProject("console", "Console Application").With(func(root *projects.ProjectItem) {
	root.AddFileWithContent("main.go", Main_tel)
	root.AddFileWithContent("go.mod", Mod_tel)
})
