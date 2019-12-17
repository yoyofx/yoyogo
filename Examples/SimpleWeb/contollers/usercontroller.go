package contollers

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
)

type UserController struct {
	*Controller.ApiController
}

func NewUserController() *UserController {
	return &UserController{}
}

type RegiserRequest struct {
	Controller.RequestParam
	UserName string `param:"username"`
	Password string `param:"password"`
}

func (p *UserController) Register(ctx *Context.HttpContext, request *RegiserRequest) Controller.ApiResult {
	result := Controller.ApiResult{Success: true, Message: "ok", Data: request}
	return result
}

func (p *UserController) GetInfo() Controller.ApiResult {
	return Controller.ApiResult{Success: true, Message: "ok"}
}
