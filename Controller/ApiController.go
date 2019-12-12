package Controller

type ApiController struct {
}

func (c *ApiController) GetView() string {
	return ""
}

type IController interface {
	GetView() string
	//ViewData interface{}
	//Data interface{}
}
