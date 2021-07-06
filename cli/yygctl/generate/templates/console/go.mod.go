package console

const ProjectItem_go_mod = `
module {{.ModelName}}

go 1.16

require github.com/yoyofx/yoyogo {{.Version}}
`
