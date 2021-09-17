package mvc

type ApiResult struct {
	Success bool
	Message string
	Data    interface{}
}

type Builder struct {
	success bool
	message string
	data    interface{}
}

func ApiResultBuilder() *Builder {
	return &Builder{}
}

func (arb *Builder) Success(success bool) *Builder {
	arb.success = success
	return arb
}

func (arb *Builder) Message(msg string) *Builder {
	arb.message = msg
	return arb
}

func (arb *Builder) MessageWithFunc(fc func() string) *Builder {
	arb.message = fc()
	return arb
}

func (arb *Builder) Data(data interface{}) *Builder {
	arb.data = data
	return arb
}

func (arb *Builder) Build() ApiResult {
	return ApiResult{
		Success: arb.success,
		Data:    arb.data,
		Message: arb.message,
	}
}

func SuccessVoid() ApiResult {
	return ApiResult{
		Success: true,
		Message: "操作成功",
	}
}

func Success(data interface{}) ApiResult {
	return ApiResult{
		Success: true,
		Data:    data,
		Message: "操作成功",
	}
}

func SuccessWithMsg(data interface{}, msg string) ApiResult {
	return ApiResult{
		Success: true,
		Data:    data,
		Message: msg,
	}
}
func SuccessWithMsgFunc(data interface{}, fc func() string) ApiResult {
	return ApiResult{
		Success: true,
		Data:    data,
		Message: fc(),
	}
}

func FailVoid() ApiResult {
	return ApiResult{
		Success: false,
		Message: "操作失败",
	}
}

func Fail(data interface{}) ApiResult {
	return ApiResult{
		Success: false,
		Data:    data,
		Message: "操作失败",
	}
}

func FailWithMsg(data interface{}, msg string) ApiResult {
	return ApiResult{
		Success: false,
		Data:    data,
		Message: msg,
	}
}

func FailWithMsgFunc(data interface{}, fc func() string) ApiResult {
	return ApiResult{
		Success: false,
		Data:    data,
		Message: fc(),
	}
}
