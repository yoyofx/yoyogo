package templates

import (
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/templates/console"
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/templates/webapi"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", webapi.Project)
	registerProject("mvc", nil)
	registerProject("grpc", nil)
	registerProject("xxl-job", nil)
}
