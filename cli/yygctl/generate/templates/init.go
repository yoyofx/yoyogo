package templates

import (
	"yygctl/generate/templates/console"
	"yygctl/generate/templates/grpc"
	"yygctl/generate/templates/mvc"
	"yygctl/generate/templates/webapi"
	"yygctl/generate/templates/xxl_job"
)

func init() {
	registerProject("console", console.Project)
	registerProject("webapi", webapi.Project)
	registerProject("mvc", mvc.Project)
	registerProject("grpc", grpc.Project)
	registerProject("xxl-job", xxl_job.Project)
}
