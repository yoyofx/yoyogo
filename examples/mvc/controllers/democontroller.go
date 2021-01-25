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
	UserName string `param:"UserName"`
	Password string `param:"Password"`
}

//GET URL  http://localhost:8080/app/v1/demo/register?UserName=max&Password=123
func (controller DemoController) Register(ctx *context.HttpContext, request *RegisterRequest) mvc.ApiResult {
	return mvc.ApiResult{Success: true, Message: "ok", Data: request}
}

//GET URL http://localhost:8080/app/v1/demo/getinfo
func (controller DemoController) GetInfo() mvc.ApiResult {
	return controller.OK("ok")
}
