package xlog

import (
	"fmt"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ConsoleColors"
	"log"
	"os"
	"time"
)

type XLogger struct {
	logger       *log.Logger
	dateFormat   string
	class        string
	logFormatter func(LogInfo) string
}

var LoggerDefaultDateFormat = "2006/01/02 15:04:05.00"

type LogLevel int

// MessageLevel
const (
	NOTSET  = iota
	DEBUG   = LogLevel(10 * iota) // DEBUG = 10
	INFO    = LogLevel(10 * iota) // INFO = 20
	WARNING = LogLevel(10 * iota) // WARNING = 30
	ERROR   = LogLevel(10 * iota) // ERROR = 40
)

var LevelString = map[LogLevel]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: ConsoleColors.Yellow("WARNING"),
	ERROR:   ConsoleColors.Red("ERROR"),
}

func defaultLogFormater(logInfo LogInfo) string {
	outLog := fmt.Sprintf(ConsoleColors.Yellow("[yoyogo] ")+"[%s] [%s] [%s] [%s] , %s",
		logInfo.StartTime, logInfo.Level, logInfo.Class, logInfo.Host, logInfo.Message)
	return outLog
}

func (log *XLogger) SetCustomLogFormat(logFormatterFunc func(logInfo LogInfo) string) {
	log.logFormatter = logFormatterFunc
}

func (log *XLogger) SetDateFormat(format string) {
	log.dateFormat = format
}

func NewXLogger() *XLogger {
	logger := NewLoggerWith(log.New(os.Stdout, "", 0))
	return logger
}

func NewLoggerWith(log *log.Logger) *XLogger {
	logger := &XLogger{logger: log, dateFormat: LoggerDefaultDateFormat}
	logger.SetCustomLogFormat(defaultLogFormater)
	return logger
}

func GetXLogger(class string) ILogger {
	logger := NewXLogger()
	logger.class = class
	return logger
}

func (log *XLogger) log(level LogLevel, format string, a ...interface{}) {
	hostName, _ := os.Hostname()
	message := format
	if len(a[0].([]interface{})) > 0 {
		message = fmt.Sprintf(format, a...)
	}

	start := time.Now()
	info := LogInfo{
		StartTime: start.Format(log.dateFormat),
		Level:     LevelString[level],
		Class:     log.class,
		Host:      hostName,
		Message:   message,
	}
	log.logger.Println(log.logFormatter(info))
}

func (log *XLogger) Debug(format string, a ...interface{}) {
	log.log(DEBUG, format, a)
}

func (log *XLogger) Info(format string, a ...interface{}) {
	log.log(INFO, format, a)
}

func (log *XLogger) Warning(format string, a ...interface{}) {
	log.log(WARNING, format, a)
}

func (log *XLogger) Error(format string, a ...interface{}) {
	log.log(ERROR, format, a)
}
