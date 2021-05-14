package mvc

import (
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/view"
)

type ApiController struct {
	view view.IViewEngine
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

func (c *ApiController) SetViewEngine(viewEngine view.IViewEngine) {
	c.view = viewEngine
}

func (c *ApiController) View(name string, data interface{}) actionresult.IActionResult {
	html, _ := c.view.ViewHtml(name, data)
	return actionresult.Html{Document: html}
}

type IController interface {
	GetName() string
	SetViewEngine(viewEngine view.IViewEngine)
	//ViewData interface{}
	//Data interface{}
}
