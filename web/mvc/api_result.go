package mvc

import "sync"

var (
	apiResultPool = sync.Pool{
		New: func() interface{} {
			return &ApiResult{Status: 200}
		},
	}
)

type ApiResult struct {
	Success bool
	Message string
	Data    interface{}
	Status  int
}

func (api ApiResult) StatusCode() int {
	return api.Status
}

type ApiResultBuilder struct {
	result *ApiResult
}

func NewApiResultBuilder() *ApiResultBuilder {
	return &ApiResultBuilder{result: apiResultPool.Get().(*ApiResult)}
}

func (arb *ApiResultBuilder) Success() *ApiResultBuilder {
	arb.result.Status = 200
	arb.result.Success = true
	return arb
}

func (arb *ApiResultBuilder) Fail() *ApiResultBuilder {
	arb.result.Success = false
	return arb
}

func (arb *ApiResultBuilder) Message(msg string) *ApiResultBuilder {
	arb.result.Message = msg
	return arb
}

func (arb *ApiResultBuilder) MessageWithFunc(fc func() string) *ApiResultBuilder {
	arb.result.Message = fc()
	return arb
}

func (arb *ApiResultBuilder) Data(data interface{}) *ApiResultBuilder {
	arb.result.Data = data
	return arb
}

func (arb *ApiResultBuilder) StatusCode(statusCode int) *ApiResultBuilder {
	arb.result.Status = statusCode
	return arb
}

func (arb *ApiResultBuilder) Build() ApiResult {
	return *arb.result
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

type ApiDocResult[T any] struct {
	Success bool
	Message string
	Data    T
	Status  int
}

type ApiDocResultBuilder[T any] struct {
	result *ApiDocResult[T]
}

func (arb *ApiDocResultBuilder[T]) Success() *ApiDocResultBuilder[T] {
	arb.result.Status = 200
	arb.result.Success = true
	return arb
}

func (arb *ApiDocResultBuilder[T]) Fail() *ApiDocResultBuilder[T] {
	arb.result.Success = false
	return arb
}

func (arb *ApiDocResultBuilder[T]) Message(msg string) *ApiDocResultBuilder[T] {
	arb.result.Message = msg
	return arb
}

func (arb *ApiDocResultBuilder[T]) MessageWithFunc(fc func() string) *ApiDocResultBuilder[T] {
	arb.result.Message = fc()
	return arb
}

func (arb *ApiDocResultBuilder[T]) Data(data T) *ApiDocResultBuilder[T] {
	arb.result.Data = data
	return arb
}

func (arb *ApiDocResultBuilder[T]) StatusCode(statusCode int) *ApiDocResultBuilder[T] {
	arb.result.Status = statusCode
	return arb
}

func ApiDocumentResult[T any]() *ApiDocResultBuilder[T] {
	return &ApiDocResultBuilder[T]{
		result: &ApiDocResult[T]{Status: 200},
	}
}

func (arb *ApiDocResultBuilder[T]) Build() ApiDocResult[T] {
	return *arb.result
}
