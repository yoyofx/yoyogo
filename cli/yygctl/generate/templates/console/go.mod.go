package console

const ProjectItem_go_mod = `
module {{.ModelName}}

go 1.16

require (
	github.com/yoyofxteam/dependencyinjection v1.0.0
	github.com/yoyofx/yoyogo {{.Version}}
)
`
