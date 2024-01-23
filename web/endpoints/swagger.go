package endpoints

import (
	"encoding/json"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/pkg/swagger"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"github.com/yoyofxteam/reflectx"
	"strings"
)

func GetAllController(router router.IRouterBuilder) map[string]map[string]swagger.Path {
	builder := router.GetMvcBuilder()
	controllerList := builder.GetControllerDescriptorList()
	pathMap := make(map[string]map[string]swagger.Path)
	for _, controller := range controllerList {
		FilterValidParams(controller, pathMap)
	}
	return pathMap
}

func FilterValidParams(controller mvc.ControllerDescriptor, pathMap map[string]map[string]swagger.Path) map[string]map[string]swagger.Path {
	suf := len(controller.ControllerName) - 10
	basePath := controller.ControllerName[0:suf]
	for _, act := range controller.GetActionDescriptors() {
		if len(act.MethodInfo.Parameters) > 0 {
			for _, param := range act.MethodInfo.Parameters {
				// 跳过HttpContext
				if param.Name == "HttpContext" {
					continue
				}
				// 跳过控制器
				if strings.HasSuffix(param.Name, "Controller") {
					continue
				}
				// 拼接api路径
				actPath := "/" + basePath + "/" + act.ActionName[len(act.ActionMethod):]
				pathInfoMap := make(map[string]swagger.Path)
				pathMap[actPath] = pathInfoMap
				paramSourceData := param.ParameterType.Elem()
				fieldNum := paramSourceData.NumField()
				//根据请求方法分类
				if act.ActionMethod == "post" {
					pathInfo := swagger.Path{}
					pathInfo.RequestBody = RequestBody(param)
					for i := 0; i < fieldNum; i++ {
						// 获取方法注释
						filed := paramSourceData.Field(i)
						if filed.Type.Name() == "RequestBody" {
							pathInfo.Description = filed.Tag.Get("note")
							pathInfo.Summary = filed.Tag.Get("note")

						}
					}
					pathInfoMap[act.ActionMethod] = pathInfo
				}

			}
		}

	}
	return pathMap
}

func RequestBody(param reflectx.MethodParameterInfo) swagger.RequestBody {
	paramSourceData := param.ParameterType.Elem()
	fieldNum := paramSourceData.NumField()
	// 获取BODY参数注释
	requestBody := swagger.RequestBody{}
	contentType := make(map[string]swagger.ContentType)
	requestBody.Content = contentType
	schema := swagger.Schema{}
	schema.Type = "object"
	schemaProperties := make(map[string]swagger.Property)
	schema.Properties = schemaProperties
	for i := 0; i < fieldNum; i++ {
		filed := paramSourceData.Field(i)
		property := swagger.Property{}
		property.Type = filed.Type.Name()
		property.Description = filed.Tag.Get("note")
		schemaProperties[filed.Name] = property
	}
	content := swagger.ContentType{}
	content.Schema = schema
	contentType["application/json"] = content
	return requestBody
}

func UseSwaggerUI(router router.IRouterBuilder, f func() swagger.Info) {
	xlog.GetXLogger("Endpoint").Debug("loaded swagger ui endpoint.")
	openapi := swagger.OpenApi{}
	openapi.Openapi = "3.0.3"
	openapi.Info = f()
	router.GET("/swagger.json", func(ctx *context.HttpContext) {
		pathMap := GetAllController(router)
		openapi.Paths = pathMap
		marshal, _ := json.Marshal(&openapi)
		swaggerJson := marshal
		ctx.Render(200, actionresult.Data{ContentType: "application/json; charset=utf-8", Data: []byte(swaggerJson)})
	})

	router.GET("/swagger", func(ctx *context.HttpContext) {
		swaggerUIHTML := `<!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>Swagger UI</title>
            <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3.52.5/swagger-ui.css">
            <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3.52.5/swagger-ui-bundle.js"></script>
        </head>
        <body>
            <div id="swagger-ui"></div>
            <script>
                window.onload = function() {
                    const ui = SwaggerUIBundle({
                        url: "http://localhost:8080/app/swagger.json",
                        dom_id: '#swagger-ui',
                        presets: [
                            SwaggerUIBundle.presets.apis,
                            SwaggerUIBundle.SwaggerUIStandalonePreset
                        ],
                      
                    })
                }
            </script>
        </body>
        </html>`
		ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		_, _ = ctx.Output.Write([]byte(swaggerUIHTML))
		ctx.Output.SetStatus(200)

	})
}
