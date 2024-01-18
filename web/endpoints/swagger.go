package endpoints

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
)

func UseSwaggerUI(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded swagger ui endpoint.")

	routeInfoArr := router.GetRouteInfo()
	builder := router.GetMvcBuilder()
	fmt.Println(builder)
	for _, routeInfo := range routeInfoArr {
		fmt.Println(routeInfo)
	}

	router.GET("/swagger.json", func(ctx *context.HttpContext) {
		swaggerJson := `{
    "swagger": "2.0",
    "info": {
        "title": "My API",
        "version": "1.0.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "schemes": ["http"],
    "paths": {
        "/users": {
            "get": {
                "summary": "List users",
                "description": "Returns a list of users.",
                "responses": {
                    "200": {
                        "description": "A list of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/User"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`
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
