package mvc

const DemoController_Tel = `
package controllers

import (
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
)

type DemoController struct {
	mvc.ApiController // 必须继承
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

//-------------------------------------------------------------------------------
type RegisterRequest struct {
	mvc.RequestBody
	UserName string` + " `param:\"UserName\"`" +
	"Password string" + "`param:\"Password\"`" +
	`}

//GET URL  http://localhost:8080/app/v1/demo/register?UserName=max&Password=123
func (controller DemoController) Register(ctx *context.HttpContext, request *RegisterRequest) mvc.ApiResult {
	return mvc.ApiResult{Success: true, Message: "ok", Data: request}
}

//GET URL http://localhost:8080/app/v1/demo/getinfo
func (controller DemoController) GetInfo() mvc.ApiResult {
	return controller.OK("ok")
}

`

const Main_Tel = `
package {{.ModelName}}

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/actionresult/extension"
	"github.com/yoyofx/yoyogo/web/mvc"
    "{{.ModelName}}/controllers"
)

func main() {
	CreateMVCBuilder().Build().Run()
}

//* Create the builder of Web host
func CreateMVCBuilder() *abstractions.HostBuilder {
	configuration := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	return web.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *web.ApplicationBuilder) {
			app.SetJsonSerializer(extension.CamelJson())
			app.UseMvc(func(builder *mvc.ControllerBuilder) {
				builder.AddViewsByConfig()                           //视图
				builder.AddController(controllers.NewDemoController) // 注册mvc controller
			})
		}).
		ConfigureServices(func(serviceCollection *dependencyinjection.ServiceCollection) {
			// ioc
		})
}
`
const Mod_Tel = `

module {{.ModelName}}


go 1.16

require (
     github.com/yoyofx/yoyogo v1.7.2
)
`

const Config_Tel = `
yoyogo:
  application:
    name: yoyogo_demo_dev
    metadata: "develop"
    server:
      type: "fasthttp"
      address: ":8080"
      path: "app"
      max_request_size: 2096157
      session:
        name: "YOYOGO_SESSIONID"
        timeout: 3600
      tls:
        cert: ""
        key: ""
      mvc:
        template: "v1/{controller}/{action}"
        views:
          path: "./static/templates"
          includes: [ "","" ]
      static:
        patten: "/"
        webroot: "./static"
      jwt:
        header: "Authorization"
        secret: "12391JdeOW^%$#@"
        prefix: "Bearer"
        expires: 3
        enable: false
        skip_path: [
            "/info",
            "/v1/user/GetInfo",
            "/v1/user/GetSD"
        ]
      cors:
        allow_origins: ["*"]
        allow_methods: ["POST","GET","PUT", "PATCH"]
        allow_credentials: true
  cloud:
    apm:
      skyworking:
        address: localhost:11800
    discovery:
      type: "nacos"
      metadata:
        url: "120.53.133.30"
        port: 80
        namespace: "public"
        group_name: ""
    #    clusters: [""]
#      type: "consul"
#      metadata:
#        address: "localhost:8500"
#        health_check: "/actuator/health"
#        tags: [""]
#      type: "eureka"
#      metadata:
#        address: "http://localhost:5000/eureka"
  datasource:
      mysql:
        name: db1
        url: tcp(cdb-amqub3mo.bj.tencentcdb.com:10042)/yoyoBlog?charset=utf8&parseTime=True
        username: root
        password: 1234abcd
        debug: true
        pool:
          init_cap: 2
          max_cap: 5
          idle_timeout : 5
      redis:
        name: reids1
        url: 62.234.6.120:31379
        password:
        db: 0
        pool:
          init_cap: 2
          max_cap: 5
          idle_timeout: 5
`
