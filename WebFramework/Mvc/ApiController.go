package Mvc

import "github.com/yoyofx/yoyogo/WebFramework/ActionResult"

type ApiController struct {
	view *ActionResult.HTMLDebug
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

func (c *ApiController) SetViewEngine(viewEngine *ActionResult.HTMLDebug) {
	c.view = viewEngine
}

func (c *ApiController) View(name string, data interface{}) ActionResult.IActionResult {
	return c.view.Instance(name, data)
}

type IController interface {
	GetName() string
	SetViewEngine(viewEngine *ActionResult.HTMLDebug)
	//ViewData interface{}
	//Data interface{}
}
