package XLog

func GetXLogger(class string) ILogger {
	logger := GetClassLogger(class) // NewXLogger()
	//logger.class = class
	return logger
}

func GetXLoggerWithFields(class string, fields map[string]interface{}) ILogger {
	logger := NewXLogger()
	logger.class = class
	return logger
}

func GetXLoggerWith(logger ILogger) ILogger {
	return logger
}
