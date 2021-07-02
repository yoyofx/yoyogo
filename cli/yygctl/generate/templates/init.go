package templates

import (
	"yygctl/generate/templates/console"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", nil)
	registerProject("mvc", nil)
	registerProject("grpc", nil)
	registerProject("xxl-job", nil)
}
