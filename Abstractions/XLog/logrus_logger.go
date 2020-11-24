package XLog

import (
	logrus "github.com/sirupsen/logrus"
	"os"
)

type LogrusLogger struct {
	logger        *logrus.Logger
	dateFormat    string
	fields        map[string]interface{}
	displayFields bool
	class         string
}

func NewLogger() ILogger {
	logger := logrus.New()
	return &LogrusLogger{logger: logger, dateFormat: LoggerDefaultDateFormat}
}

func GetClassLogger(class string) ILogger {
	logger := logrus.New()
	logger.Out = os.Stdout

	logger.SetLevel(logrus.DebugLevel)
	logger.Formatter = &TextFormatter{
		DisableColors:   false,
		ForceColors:     false,
		TimestampFormat: LoggerDefaultDateFormat,
		FullTimestamp:   true,
		ForceFormatting: true,
	}
	return &LogrusLogger{logger: logger, class: class, dateFormat: LoggerDefaultDateFormat, displayFields: true}
}

func (log *LogrusLogger) With(level LogLevel, fiedls map[string]interface{}) *logrus.Entry {

	//start := time.Now()

	fieldsMap := make(map[string]interface{})
	fieldsMap["prefix"] = "YOYOGO"
	if fiedls != nil {
		fieldsMap = fiedls
	}

	if log.displayFields {
		fieldsMap["class"] = log.class
		hostName, _ := os.Hostname()
		fieldsMap["host"] = hostName
	}
	//fieldsMap["message"] = message

	return log.logger.WithFields(fieldsMap)
}

func (log LogrusLogger) Debug(format string, a ...interface{}) {
	log.With(DEBUG, log.fields).Debugf(format, a...)
}

func (log LogrusLogger) Info(format string, a ...interface{}) {
	log.With(INFO, log.fields).Infof(format, a...)
}

func (log LogrusLogger) Warning(format string, a ...interface{}) {
	log.With(WARNING, log.fields).Warnf(format, a...)
}

func (log LogrusLogger) Error(format string, a ...interface{}) {
	log.logger.Out = os.Stderr
	log.With(ERROR, log.fields).Errorf(format, a...)
	log.logger.Out = os.Stdout
}

func (log *LogrusLogger) SetClass(className string) {
	log.class = className
}

func (log *LogrusLogger) SetCustomLogFormat(logFormatterFunc func(logInfo interface{}) string) {
	log.displayFields = false
}

func (log LogrusLogger) SetDateFormat(format string) {
	//log.dateFormat = format
}
