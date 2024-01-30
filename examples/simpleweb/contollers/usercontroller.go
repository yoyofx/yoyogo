package contollers

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/binding"
	"github.com/yoyofx/yoyogo/web/captcha"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"mime/multipart"
	"simpleweb/models"
)

type UserController struct {
	mvc.ApiController `route:"user" doc:"用户接口Controller"`

	userAction      models.IUserAction
	discoveryClient servicediscovery.IServiceDiscovery
	config          abstractions.IConfiguration
}

func NewUserController(userAction models.IUserAction, sd servicediscovery.IServiceDiscovery) *UserController {
	return &UserController{userAction: userAction, discoveryClient: sd}
}

type RegisterRequest struct {
	mvc.RequestBody `route:"/api/users/register" doc:"用户注册"`

	UserName   string `uri:"userName" doc:"用户名"`
	Password   string `uri:"password" doc:"密码"`
	TestNumber uint64 `uri:"num" doc:"数字"`
}

func (controller UserController) Register(ctx *context.HttpContext, request *RegisterRequest) mvc.ApiResult {
	num := context.Query2Number[uint64](ctx, "num", "55")

	fmt.Println(num)
	return mvc.ApiResult{Success: true, Message: "ok", Data: request}
}

type PostUserInfoRequest struct {
	mvc.RequestBody `doc:"用户信息提交"` //`route:"/{id}"`
	UserName        string         `form:"userName" json:"userName" doc:"用户名"`
	Password        string         `form:"password" json:"password" doc:"密码"`
	Token           string         `header:"Authorization" json:"token" doc:"token"`
}

func (controller UserController) PostUserInfo(ctx *context.HttpContext, request *PostUserInfoRequest) actionresult.IActionResult {
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

func (controller UserController) GetDoc() mvc.ApiDocResult[string] {

	return mvc.ApiDocumentResult[string]().Success().Data("ok").Message("hello").Build()
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

func (controller UserController) GetTestApiResult() mvc.ApiResult {
	return controller.ApiResult().
		Success().
		Data("ok").
		Message("hello").
		StatusCode(400).Build()
}

type UserInfo struct {
	UserName string                `form:"user" json:"user" binding:"required" doc:"用户名"`
	Number   int                   `form:"num" json:"num" binding:"gt=0,lt=10" doc:"数字"`
	Id       string                `form:"id" json:"id" binding:"required,gt=0,lt=10" doc:"id"`
	Image    *multipart.FileHeader `form:"file" doc:"图片文件"`
}

// FromBody
func (controller UserController) DefaultBinding(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := &UserInfo{}
	err := ctx.Bind(userInfo)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(userInfo)
}

// FromBody
func (controller UserController) JsonBinding(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := &UserInfo{}
	err := ctx.BindWith(userInfo, binding.JSON)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(userInfo)
}

// FromQuery
func (controller UserController) GetQueryBinding(ctx *context.HttpContext) mvc.ApiResult {
	fmt.Println("进入方法")
	fmt.Println(controller.config.Get("env"))

	userInfo := &UserInfo{}
	err := ctx.BindWith(userInfo, binding.Query)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(userInfo)

}

type UploadForm struct {
	mvc.RequestBody `doc:"文件上传"`
	File            *multipart.FileHeader `form:"file1" doc:"文件"`
	Key             string                `form:"key" doc:"文件ID"`
}

func (controller UserController) Upload(form *UploadForm) mvc.ApiResult {
	return controller.OK(context.H{
		"file": form.File.Filename,
		"size": form.File.Size,
		"key":  form.Key,
	})

}

func (controller UserController) TestFunc(request *struct {
	mvc.RequestGET `route:"/v1/user/:id/test" doc:"测试接口"`
	Name           string `uri:"name" doc:"测试用户名"`
	Id             uint64 `path:"id" doc:"测试ID"`
}) mvc.ApiResult {

	return mvc.Success(request)
}
