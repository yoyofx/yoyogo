package Controller

type UserController struct {
	//*ApiController
}

func (c *UserController) GetView() string {
	return ""
}

func NewUserController() *UserController {
	return &UserController{}
}

type RegiserRequest struct {
	RequestParam
	UserName string
	Password string
}

func (p *UserController) Register(request *RegiserRequest) ApiResult {

	result := ApiResult{Success: true, Message: "ok", Data: request}
	return result
}

func (p *UserController) GetInfo() ApiResult {
	return ApiResult{Success: true, Message: "ok"}
}
