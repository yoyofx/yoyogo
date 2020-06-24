package contollers

import (
	"github.com/yoyofx/yoyogo/Examples/SimpleWeb/models"
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
)

type UserController struct {
	*Mvc.ApiController
	userAction models.IUserAction
}

func NewUserController(userAction models.IUserAction) *UserController {
	return &UserController{userAction: userAction}
}

type RegiserRequest struct {
	Mvc.RequestBody
	UserName string `param:"username"`
	Password string `param:"password"`
}

func (this *UserController) Register(ctx *Context.HttpContext, request *RegiserRequest) ActionResult.IActionResult {
	result := Mvc.ApiResult{Success: true, Message: "ok", Data: request}

	return ActionResult.Json{Data: result}
}

func (this *UserController) GetInfo() Mvc.ApiResult {

	return this.OK(this.userAction.Login("zhang"))
}
