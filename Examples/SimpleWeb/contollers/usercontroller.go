package contollers

import (
	"SimpleWeb/models"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/captcha"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
)

type UserController struct {
	mvc.ApiController
	userAction      models.IUserAction
	discoveryClient servicediscovery.IServiceDiscovery
}

func NewUserController(userAction models.IUserAction, sd servicediscovery.IServiceDiscovery) *UserController {
	return &UserController{userAction: userAction, discoveryClient: sd}
}

type RegisterRequest struct {
	mvc.RequestBody
	UserName string `param:"UserName"`
	Password string `param:"Password"`
}

func (controller UserController) Register(ctx *context.HttpContext, request *RegisterRequest) mvc.ApiResult {

	return mvc.ApiResult{Success: true, Message: "ok", Data: request}
}

func (controller UserController) GetUserName(ctx *context.HttpContext, request *RegisterRequest) actionresult.IActionResult {
	result := mvc.ApiResult{Success: true, Message: "ok", Data: request}

	return actionresult.Json{Data: result}
}

func (controller UserController) PostUserInfo(ctx *context.HttpContext, request *RegisterRequest) actionresult.IActionResult {

	return actionresult.Json{Data: mvc.ApiResult{Success: true, Message: "ok", Data: context.H{
		"user":    ctx.GetUser(),
		"request": request,
	}}}
}

func (controller UserController) GetHtmlHello() actionresult.IActionResult {
	return controller.View("hello", map[string]interface{}{
		"name": "hello world!",
	})
}

func (controller UserController) GetHtmlBody() actionresult.IActionResult {
	return controller.View("raw", map[string]interface{}{
		"body": "raw.htm hello world!",
	})
}

func (controller UserController) GetInfo() mvc.ApiResult {

	return controller.OK(controller.userAction.Login("zhang"))
}

func (controller UserController) GetSD() mvc.ApiResult {
	serviceList := controller.discoveryClient.GetAllInstances("yoyogo_demo_dev")
	return controller.OK(serviceList)
}

func (controller UserController) GetCaptcha(ctx *context.HttpContext) actionresult.IActionResult {
	_, md5, bytes := captcha.CreateImage(6)
	ctx.GetSession().SetValue("cimg_md5", md5)
	return actionresult.Image{Data: bytes}
}

func (controller UserController) GetValidation(ctx *context.HttpContext) mvc.ApiResult {
	text := ctx.Input.Query("val")
	md5 := ctx.GetSession().GetString("cimg_md5")
	ok := captcha.Validation(text, md5)
	return controller.OK(context.H{"validation": ok})
}
