package templates

import "github.com/yoyofx/yoyogo/cli/yygctl/generate/projects"

var demoProject = projects.NewEmptyProject("demo").With(func(root *projects.ProjectItem) {
	root.AddFileWithContent("main.go", `
		func main() {
		}
	`)

	root.AddDir("controllers")
	models := root.AddDir("models")
	models.AddFile("userDto.go")
})
