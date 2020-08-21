package Mvc

import (
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/View"
)

type ApiController struct {
	view View.IViewEngine
}

func (c *ApiController) GetName() string {
	return "controller"
}

func (c *ApiController) OK(data interface{}) ApiResult {
	return ApiResult{Success: true, Message: "true", Data: data}
}

func (c *ApiController) Fail(msg string) ApiResult {
	return ApiResult{Success: false, Message: msg}
}

func (c *ApiController) SetViewEngine(viewEngine View.IViewEngine) {
	c.view = viewEngine
}

func (c *ApiController) View(name string, data interface{}) ActionResult.IActionResult {
	html, _ := c.view.ViewHtml(name, data)
	return ActionResult.Html{Document: html}
}

type IController interface {
	GetName() string
	SetViewEngine(viewEngine View.IViewEngine)
	//ViewData interface{}
	//Data interface{}
}
