package templates

import (
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/templates/webapi"
	"yygctl/generate/templates/console"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", webapi.Project)
	registerProject("mvc", nil)
	registerProject("grpc", nil)
	registerProject("xxl-job", nil)
}
