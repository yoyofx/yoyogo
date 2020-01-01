package contollers

import (
	"github.com/maxzhang1985/yoyogo/ActionResult"
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
	"github.com/maxzhang1985/yoyogo/Examples/SimpleWeb/models"
)

type UserController struct {
	*Controller.ApiController
	userAction models.IUserAction
}

func NewUserController(userAction models.IUserAction) *UserController {
	return &UserController{userAction: userAction}
}

type RegiserRequest struct {
	Controller.RequestBody
	UserName string `param:"username"`
	Password string `param:"password"`
}

func (this *UserController) Register(ctx *Context.HttpContext, request *RegiserRequest) ActionResult.IActionResult {
	result := Controller.ApiResult{Success: true, Message: "ok", Data: request}

	return ActionResult.Json{Data: result}
}

func (this *UserController) GetInfo() Controller.ApiResult {

	return this.OK(this.userAction.Login("zhang"))
}
