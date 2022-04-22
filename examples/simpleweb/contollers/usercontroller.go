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
	mvc.ApiController `route:"user"`

	userAction      models.IUserAction
	discoveryClient servicediscovery.IServiceDiscovery
	config          abstractions.IConfiguration
}

func NewUserController(userAction models.IUserAction, sd servicediscovery.IServiceDiscovery) *UserController {
	return &UserController{userAction: userAction, discoveryClient: sd}
}

type RegisterRequest struct {
	mvc.RequestBody `route:"/v1/users/register"`

	UserName   string `uri:"userName"`
	Password   string `uri:"password"`
	TestNumber uint64 `uri:"num"`
}

func (controller UserController) Register(ctx *context.HttpContext, request *RegisterRequest) mvc.ApiResult {
	num := context.Query2Number[uint64](ctx, "num", "55")

	fmt.Println(num)
	return mvc.ApiResult{Success: true, Message: "ok", Data: request}
}

type PostUserInfoRequest struct {
	mvc.RequestBody //`route:"/{id}"`

	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
	Token    string `header:"Authorization" json:"token"`
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
	UserName string                `form:"user" json:"user" binding:"required"`
	Number   int                   `form:"num" json:"num" binding:"gt=0,lt=10"`
	Id       string                `form:"id" json:"id" binding:"required,gt=0,lt=10"`
	Image    *multipart.FileHeader `form:"file"`
}

//FromBody
func (controller UserController) DefaultBinding(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := &UserInfo{}
	err := ctx.Bind(userInfo)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(userInfo)
}

//FromBody
func (controller UserController) JsonBinding(ctx *context.HttpContext) mvc.ApiResult {
	userInfo := &UserInfo{}
	err := ctx.BindWith(userInfo, binding.JSON)
	if err != nil {
		return controller.Fail(err.Error())
	}
	return controller.OK(userInfo)
}

//FromQuery
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
	mvc.RequestBody
	File *multipart.FileHeader `form:"file1"`
	Key  string                `form:"key"`
}

func (controller UserController) Upload(form *UploadForm) mvc.ApiResult {
	return controller.OK(context.H{
		"file": form.File.Filename,
		"size": form.File.Size,
		"key":  form.Key,
	})

}

func (controller UserController) TestFunc(request *struct {
	mvc.RequestGET `route:"/v1/user/:id/test"`
	Name           string `uri:"name"`
	Id             uint64 `path:"id"`
}) mvc.ApiResult {

	return mvc.Success(request)
}
