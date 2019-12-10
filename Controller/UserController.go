package Controller

type UserController struct {
	*ApiController
	Name string
	//
}

func NewUserController() *UserController {
	return &UserController{Name: "www"}
}

type RegiserRequest struct {
	RequestParam
	UserName string
	Password string
}

func (p *UserController) Register(u string, request *RegiserRequest) ApiResult {

	result := ApiResult{Success: true, Message: "ok", Data: request}
	return result
}

func (p *UserController) GetInfo() ApiResult {
	return ApiResult{Success: true, Message: "ok"}
}
