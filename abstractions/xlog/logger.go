package xlog

type ILogger interface {
	SetClass(className string)

	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})

	SetCustomLogFormat(logFormatterFunc func(logInfo interface{}) string)
	SetDateFormat(format string)
}
