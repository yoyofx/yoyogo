package mvc

import "yygctl/generate/projects"

var Project = projects.NewEmptyProject("mvc", "MVC Application").With(func(root *projects.ProjectItem) {
	root.AddDir("controller").AddFileWithContent("democontroller.go", DemoController_Tel)
	root.AddFileWithContent("config_dev.yml", Config_Tel)
	root.AddFileWithContent("go.mod", Mod_Tel)
	root.AddFileWithContent("main.go", Main_Tel)
})
