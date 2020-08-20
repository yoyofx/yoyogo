package contollers

import (
	"github.com/yoyofx/yoyogo/Examples/SimpleWeb/models"
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
)

type UserController struct {
	Mvc.ApiController
	userAction models.IUserAction
}

func NewUserController(userAction models.IUserAction) *UserController {
	return &UserController{userAction: userAction}
}

type RegisterRequest struct {
	Mvc.RequestBody
	UserName string `param:"UserName"`
	Password string `param:"Password"`
}

func (controller UserController) Register(ctx *Context.HttpContext, request *RegisterRequest) ActionResult.IActionResult {
	result := Mvc.ApiResult{Success: true, Message: "ok", Data: request}

	return ActionResult.Json{Data: result}
}

func (controller UserController) GetUserName(ctx *Context.HttpContext, request *RegisterRequest) ActionResult.IActionResult {
	result := Mvc.ApiResult{Success: true, Message: "ok", Data: request}

	return ActionResult.Json{Data: result}
}

func (controller UserController) PostUserInfo(ctx *Context.HttpContext, request *RegisterRequest) ActionResult.IActionResult {

	return ActionResult.Json{Data: Mvc.ApiResult{Success: true, Message: "ok", Data: Context.H{
		"user":    ctx.GetUser(),
		"request": request,
	}}}
}

func (controller UserController) GetHtmlHello() ActionResult.IActionResult {
	return controller.View("hello.tmpl", map[string]interface{}{
		"name": "hello world!",
	})
}

func (controller UserController) GetHtmlBody() ActionResult.IActionResult {
	return controller.View("raw.tmpl", map[string]interface{}{
		"body": "raw.htm hello world!",
	})
}

func (controller UserController) GetInfo() Mvc.ApiResult {

	return controller.OK(controller.userAction.Login("zhang"))
}
