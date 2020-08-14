package xlog

type ILogger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})

	SetCustomLogFormat(logFormatterFunc func(logInfo LogInfo) string)
	SetDateFormat(format string)
}
