package XLog

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogrusLogger struct {
	logger     *logrus.Logger
	dateFormat string
	class      string
}

func NewLogger() ILogger {
	logger := logrus.New()
	return &LogrusLogger{logger: logger, dateFormat: LoggerDefaultDateFormat}
}

func GetClassLogger(class string) ILogger {
	logger := logrus.New()
	return &LogrusLogger{logger: logger, class: class, dateFormat: LoggerDefaultDateFormat}
}

func (log *LogrusLogger) log(level LogLevel, fiedls map[string]interface{}, format string, a ...interface{}) *logrus.Entry {
	hostName, _ := os.Hostname()
	message := format
	message = fmt.Sprintf(format, a...)
	start := time.Now()

	fieldsMap := make(map[string]interface{})
	if fiedls != nil {
		fieldsMap = fiedls
	}
	fieldsMap["StartTime"] = start.Format(log.dateFormat)
	fieldsMap["Level"] = LevelString[level]
	fieldsMap["Class"] = log.class
	fieldsMap["Host"] = hostName
	fieldsMap["Message"] = message

	return logrus.WithFields(fieldsMap)
}

func (l LogrusLogger) Debug(format string, a ...interface{}) {
	panic("implement me")
}

func (l LogrusLogger) Info(format string, a ...interface{}) {
	panic("implement me")
}

func (l LogrusLogger) Warning(format string, a ...interface{}) {
	panic("implement me")
}

func (l LogrusLogger) Error(format string, a ...interface{}) {
	panic("implement me")
}

func (l LogrusLogger) SetCustomLogFormat(logFormatterFunc func(logInfo LogInfo) string) {
	panic("implement me")
}

func (l LogrusLogger) SetDateFormat(format string) {
	panic("implement me")
}
