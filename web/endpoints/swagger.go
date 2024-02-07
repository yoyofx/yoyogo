package endpoints

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/pkg/swagger"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"github.com/yoyofxteam/reflectx"
	"reflect"
	"regexp"
	"strings"
)

func UseSwaggerDoc(router router.IRouterBuilder, info swagger.Info, configFunc func(openapi *swagger.OpenApi)) {
	xlog.GetXLogger("Endpoint").Debug("loaded swagger ui endpoint.")

	// swagger.json
	router.GET("/resources/swagger.json", func(ctx *context.HttpContext) {
		var env *abstractions.HostEnvironment
		_ = ctx.RequiredServices.GetService(&env)
		baseUrl := fmt.Sprintf("http://localhost:%s", env.Port)
		openapi := swagger.NewOpenApi(baseUrl, info)

		configFunc(openapi)
		GetSwaggerRouteInformation(openapi, router, env)
		ctx.JSON(200, openapi)
	})

	// swagger ui
	router.GET("/resources/swagger", func(ctx *context.HttpContext) {
		var env *abstractions.HostEnvironment
		_ = ctx.RequiredServices.GetService(&env)
		baseUrl := fmt.Sprintf("http://localhost:%s", env.Port)
		serverPath := env.MetaData["server.path"]
		// swagger json address
		swaggerJsonUri := fmt.Sprintf("%s/%s/resources/swagger.json", baseUrl, serverPath)
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
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5.11.2/swagger-ui.css" />
			  </head>
			  <body>
			  <div id="swagger-ui"></div>
			  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5.11.2/swagger-ui-bundle.js" crossorigin></script>
			  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5.11.2/swagger-ui-standalone-preset.js" crossorigin></script>
			  <script>
				window.onload = () => {
				  window.ui = SwaggerUIBundle({
					url: '%s',
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
		swaggerUIHTML = fmt.Sprintf(swaggerUIHTML, swaggerJsonUri)
		ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		_, _ = ctx.Output.Write([]byte(swaggerUIHTML))
		ctx.Output.SetStatus(200)

	})
}

func GetSwaggerRouteInformation(openapi *swagger.OpenApi, router router.IRouterBuilder, env *abstractions.HostEnvironment) {
	// mvc routes
	builder := router.GetMvcBuilder()
	controllerList := builder.GetControllerDescriptorList()
	for _, controller := range controllerList {
		getMvcRouters(controller, openapi, env)
	}

	// default routes
	getEndpointRouters(openapi, router, env)
}

func getEndpointRouters(openapi *swagger.OpenApi, router router.IRouterBuilder, env *abstractions.HostEnvironment) {
	// default normal route ,such as rb.POST("/info/:id", PostInfo)
	routerInfoList := router.GetRouteInfo()
	openapi.Tags = append(openapi.Tags, swagger.Tag{Name: "default", Description: fmt.Sprintf("Endpoints of the default route. (%v)", len(routerInfoList))})
	for idx, _ := range routerInfoList {
		// uri parameters
		pathInfo := swagger.Path{
			Tags:       []string{"default"},
			Responses:  map[string]swagger.ResponsesItem{},
			Parameters: []swagger.Parameters{}}

		actPath := fmt.Sprintf("/%s%s", env.MetaData["server.path"], routerInfoList[idx].Path)
		// used regexp ,replace :id to {id}
		if strings.Contains(actPath, ":") {
			reg := regexp.MustCompile(`:([a-zA-Z0-9]+)`)
			matches := reg.FindAllString(actPath, -1)
			var params []swagger.Parameters
			if len(matches) > 0 {
				for _, match := range matches {
					paramName := strings.Replace(match, ":", "", -1)
					params = append(params, swagger.Parameters{In: "path", Name: paramName})
				}
				pathInfo.Parameters = params
			}
			actPath = reg.ReplaceAllString(actPath, "{$1}")
		}

		pathInfo.Responses["200"] = swagger.ResponsesItem{Description: "OK"}
		openapi.Paths[actPath] = map[string]swagger.Path{strings.ToLower(routerInfoList[idx].Method): pathInfo}
	}
}

func getMvcRouters(controller mvc.ControllerDescriptor, openapi *swagger.OpenApi, env *abstractions.HostEnvironment) {
	serverPath := env.MetaData["server.path"]
	mvcTemplate := env.MetaData["mvc.template"]
	// mvc
	mvcTemplate = strings.ReplaceAll(mvcTemplate, "{controller}", "%s")
	mvcTemplate = strings.ReplaceAll(mvcTemplate, "{action}", "%s")
	mvcTemplate = fmt.Sprintf("/%s/%s", serverPath, mvcTemplate)

	suf := len(controller.ControllerName) - 10
	controllerName := controller.ControllerName[0:suf]
	openapi.Tags = append(openapi.Tags, swagger.Tag{Name: controller.ControllerName, Description: controller.Descriptor})

	for _, act := range controller.GetActionDescriptors() {
		// 遍历 action, 拼接api路径 swagger.Path
		pathInfo := swagger.Path{}
		pathInfo.Tags = []string{controller.ControllerName}
		actionName := strings.ReplaceAll(strings.ToLower(act.ActionName), act.ActionMethod, "")
		//actPath := "/" + controllerName + "/" + actionName
		actPath := fmt.Sprintf(mvcTemplate, controllerName, actionName)
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
		if act.ActionMethod == "any" {
			act.ActionMethod = "get"
		}
		if act.IsAttributeRoute {
			actPath = act.Route.Template
			actPath = fmt.Sprintf("/%s%s", serverPath, actPath)
			// used regexp ,replace :id to {id}
			reg := regexp.MustCompile(`:([a-zA-Z0-9]+)`)
			actPath = reg.ReplaceAllString(actPath, "{$1}")
			pathInfo.Summary = pathInfo.Summary + " ( Route Attribute ) "
		} else {
			pathInfo.Summary = pathInfo.Summary + " ( MVC ) "
		}
		// responses
		pathInfo.Responses = make(map[string]swagger.ResponsesItem)
		responseType := act.MethodInfo.OutType
		// is struct

		if responseType != nil && responseType.Kind() == reflect.Struct {
			// struct , ApiResult , ApiDocResult[?]
			// println(responseType.Name())
			// new struct type to object
			responseObject := reflect.New(responseType).Elem().Interface()

			swaggerResponse := swagger.ConvertToSwaggerResponse(responseObject)
			responseItem := swagger.ResponsesItem{Description: "OK", Content: make(map[string]interface{})}
			responseItem.Content["application/json"] = map[string]interface{}{"schema": swaggerResponse}
			pathInfo.Responses["200"] = responseItem
		} else {
			pathInfo.Responses["200"] = swagger.ResponsesItem{Description: "OK"}
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
		fieldTypeName := strings.ToLower(filed.Type.Name())
		if fieldTypeName == "" {
			fieldTypeName = strings.ToLower(filed.Type.Elem().Name())
		}
		if strings.HasPrefix(fieldTypeName, "Request") {
			continue
		}
		uriField := filed.Tag.Get("uri")
		formField := filed.Tag.Get("form")
		jsonField := filed.Tag.Get("json")
		pathField := filed.Tag.Get("path")
		headerField := filed.Tag.Get("header")

		if uriField != "" || pathField != "" || headerField != "" || formField != "" {
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
			if fieldName == "" {
				fieldName = formField
				params.In = "formData"
			}

			params.Name = fieldName
			params.Schema = struct {
				Type string `json:"type"`
			}(struct{ Type string }{
				Type: swagger.GetSwaggerType(fieldTypeName),
			})

			params.Description = filed.Tag.Get("doc")
			parameterList = append(parameterList, params)
		}

		if jsonField != "" {
			//if contentTypeStr == "" {
			//	contentTypeStr = "application/x-www-form-urlencoded"
			//}
			fieldName := jsonField

			property := swagger.Property{}
			property.Type = strings.ToLower(filed.Type.Name())
			if property.Type == "" {
				property.Type = strings.ToLower(filed.Type.Elem().Name())
			}
			property.Type = swagger.GetSwaggerType(property.Type)
			//if property.Type == "file" {
			//	property.Format = "binary"
			//}

			property.Description = filed.Tag.Get("doc")
			schemaProperties[fieldName] = property
		}
	}
	//application/x-www-form-urlencoded
	//multipart/form-data

	if len(schemaProperties) > 0 {
		//if contentTypeStr == "" {
		contentTypeStr = "application/json"

		content := swagger.ContentType{Schema: schema}
		contentType[contentTypeStr] = content
	} else {
		contentType = nil
	}
	return swagger.RequestBody{Content: contentType}, parameterList
}
