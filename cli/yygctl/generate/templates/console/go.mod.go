package console

const ProjectItem_go_mod = `
module {{.projectName}}

go 1.16

require github.com/yoyofx/yoyogo {{.version}}
`
