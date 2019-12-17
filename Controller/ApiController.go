package Controller

type ApiController struct {
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

type IController interface {
	GetName() string
	//ViewData interface{}
	//Data interface{}
}
