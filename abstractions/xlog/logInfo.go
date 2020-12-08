package xlog

type LogInfo struct {
	StartTime string
	Level     string
	Class     string
	Host      string
	Message   string
	Extend    map[string]interface{}
}
