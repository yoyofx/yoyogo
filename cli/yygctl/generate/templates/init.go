package templates

import (
	"yygctl/generate/templates/console"
	"yygctl/generate/templates/mvc"
	"yygctl/generate/templates/webapi"
	"yygctl/generate/templates/xxl_job"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", webapi.Project)
	registerProject("mvc", mvc.Project)
	registerProject("grpc", nil)
	registerProject("xxl-job", xxl_job.Project)
}
