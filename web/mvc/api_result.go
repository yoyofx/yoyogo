package mvc

type ApiResult struct {
	Success bool
	Message string
	Data    interface{}
}
