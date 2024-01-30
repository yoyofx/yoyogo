package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/pkg/swagger"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"github.com/yoyofxteam/reflectx"
	"regexp"
	"strings"
)

func GetSwaggerRouteInfomation(openapi *swagger.OpenApi, router router.IRouterBuilder) {
	builder := router.GetMvcBuilder()
	controllerList := builder.GetControllerDescriptorList()
	for _, controller := range controllerList {
		FilterValidParams(controller, openapi)
	}
}

func FilterValidParams(controller mvc.ControllerDescriptor, openapi *swagger.OpenApi) {
	suf := len(controller.ControllerName) - 10
	controllerName := controller.ControllerName[0:suf]
	openapi.Tags = append(openapi.Tags, swagger.Tag{Name: controller.ControllerName, Description: controller.Descriptor})

	for _, act := range controller.GetActionDescriptors() {
		// 遍历 action, 拼接api路径 swagger.Path
		pathInfo := swagger.Path{}
		pathInfo.Tags = []string{controller.ControllerName}
		actionName := strings.ReplaceAll(strings.ToLower(act.ActionName), act.ActionMethod, "")
		actPath := "/" + controllerName + "/" + actionName

		// action params
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

				paramSourceData := param.ParameterType.Elem()
				fieldNum := paramSourceData.NumField()
				//根据请求方法分类
				if act.ActionMethod == "post" || act.ActionMethod == "any" {
					for i := 0; i < fieldNum; i++ {
						// 获取方法注释
						filed := paramSourceData.Field(i)
						if strings.HasPrefix(filed.Type.Name(), "Request") {
							if filed.Type.Name() == "RequestBody" || filed.Type.Name() == "RequestPOST" {
								act.ActionMethod = "post"
							} else {
								actionMethodDef := filed.Type.Name()
								actionMethodDef = strings.ReplaceAll(actionMethodDef, "Request", "")
								act.ActionMethod = strings.ToLower(actionMethodDef) // get / head / delete / options / patch / put
							}
							// 获取BODY参数注释 RequestBody or RequestGET or RequestPOST
							body, parameters := RequestBody(param)
							if body.Content != nil {
								pathInfo.RequestBody = body
							}
							if len(parameters) > 0 {
								pathInfo.Parameters = parameters
							}

							pathInfo.Description = filed.Tag.Get("doc")
							pathInfo.Summary = filed.Tag.Get("doc")
							break
						}

					}
				}

			}
		}
		if act.ActionMethod == "any" {
			act.ActionMethod = "get"
		}
		if act.IsAttributeRoute {
			actPath = act.Route.Template
			// used regexp ,replace :id to {id}
			reg := regexp.MustCompile(`:[a-zA-Z0-9]+`)
			actPath = reg.ReplaceAllString(actPath, "{$0}")
			actPath = strings.ReplaceAll(actPath, ":", "")
		}

		openapi.Paths[actPath] = map[string]swagger.Path{act.ActionMethod: pathInfo}
	}
}

func RequestBody(param reflectx.MethodParameterInfo) (swagger.RequestBody, []swagger.Parameters) {
	paramSourceData := param.ParameterType.Elem()
	fieldNum := paramSourceData.NumField()
	// 获取参数
	var parameterList []swagger.Parameters
	// 获取BODY参数注释
	contentTypeStr := ""
	contentType := make(map[string]swagger.ContentType)
	schema := swagger.Schema{}
	schema.Type = "object"
	schemaProperties := make(map[string]swagger.Property)
	schema.Properties = schemaProperties
	for i := 0; i < fieldNum; i++ {
		filed := paramSourceData.Field(i)
		if strings.HasPrefix(filed.Type.Name(), "Request") {
			continue
		}
		uriField := filed.Tag.Get("uri")
		formField := filed.Tag.Get("form")
		jsonField := filed.Tag.Get("json")
		pathField := filed.Tag.Get("path")
		headerField := filed.Tag.Get("header")

		if uriField != "" || pathField != "" || headerField != "" {
			// 构建参数
			params := swagger.Parameters{}
			params.In = "query"
			fieldName := uriField
			if fieldName == "" {
				fieldName = pathField
				params.In = "path"
			}
			if fieldName == "" {
				fieldName = headerField
				params.In = "header"
			}
			params.Name = fieldName
			params.Description = filed.Tag.Get("doc")
			parameterList = append(parameterList, params)
		}

		if formField != "" || jsonField != "" {
			fieldName := formField
			if fieldName == "" {
				fieldName = jsonField
			}
			property := swagger.Property{}
			property.Type = strings.ToLower(filed.Type.Name())
			if property.Type == "" {
				property.Type = strings.ToLower(filed.Type.Elem().Name())
			}
			if strings.Contains(property.Type, "int") {
				property.Type = "integer"
			} else if strings.Contains(property.Type, "file") {
				property.Type = "string"
				property.Format = "binary"
			}

			property.Description = filed.Tag.Get("doc")
			schemaProperties[fieldName] = property
		}
	}
	//application/x-www-form-urlencoded
	//multipart/form-data
	if len(schemaProperties) > 0 {
		contentTypeStr = "application/json"
		content := swagger.ContentType{Schema: schema}
		contentType[contentTypeStr] = content
	} else {
		contentType = nil
	}
	return swagger.RequestBody{Content: contentType}, parameterList
}

func UseSwaggerUI(router router.IRouterBuilder, f func() swagger.Info) {
	xlog.GetXLogger("Endpoint").Debug("loaded swagger ui endpoint.")

	router.GET("/swagger.json", func(ctx *context.HttpContext) {
		openapi := &swagger.OpenApi{
			Openapi: "3.1.0",
			Paths:   make(map[string]map[string]swagger.Path)}
		openapi.Info = f()
		GetSwaggerRouteInfomation(openapi, router)
		ctx.JSON(200, openapi)
	})

	router.GET("/swagger", func(ctx *context.HttpContext) {
		swaggerUIHTML := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta
      name="description"
      content="SwaggerUI"
    />
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5.0.0/swagger-ui.css" />
  </head>
  <body>
  <div id="swagger-ui"></div>
  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5.0.0/swagger-ui-bundle.js" crossorigin></script>
  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5.0.0/swagger-ui-standalone-preset.js" crossorigin></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: 'http://localhost:8080/app/swagger.json',
        dom_id: '#swagger-ui',
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        layout: "StandaloneLayout",
      });
    };
  </script>
  </body>
</html>`
		ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		_, _ = ctx.Output.Write([]byte(swaggerUIHTML))
		ctx.Output.SetStatus(200)

	})
}
