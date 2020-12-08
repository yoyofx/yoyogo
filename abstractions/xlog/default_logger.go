package xlog

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"log"
	"os"
	"time"
)

type XDefaultLogger struct {
	logger       *log.Logger
	dateFormat   string
	class        string
	logFormatter func(interface{}) string
}

func NewXLogger() *XDefaultLogger {
	logger := NewLoggerWith(log.New(os.Stdout, "", 0))
	return logger
}

func NewLoggerWith(log *log.Logger) *XDefaultLogger {
	logger := &XDefaultLogger{logger: log, dateFormat: LoggerDefaultDateFormat}
	logger.SetCustomLogFormat(defaultLogFormatter)
	return logger
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
	WARNING: consolecolors.Yellow("WARNING"),
	ERROR:   consolecolors.Red("ERROR"),
}

func defaultLogFormatter(log interface{}) string {
	logInfo := log.(LogInfo)
	outLog := fmt.Sprintf("%s [%s] [%s] [%s] [%s] , %s",
		consolecolors.Yellow("[yoyogo]"), logInfo.StartTime, logInfo.Level, logInfo.Class, logInfo.Host, logInfo.Message)
	return outLog
}

func (log *XDefaultLogger) SetClass(className string) {
	log.class = className
}

func (log *XDefaultLogger) SetCustomLogFormat(logFormatterFunc func(logInfo interface{}) string) {
	log.logFormatter = logFormatterFunc
}

func (log *XDefaultLogger) SetDateFormat(format string) {
	log.dateFormat = format
}

func (log *XDefaultLogger) log(level LogLevel, format string, a ...interface{}) {
	hostName, _ := os.Hostname()
	message := format
	message = fmt.Sprintf(format, a...)

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

func (log *XDefaultLogger) Debug(format string, a ...interface{}) {
	log.log(DEBUG, format, a...)
}

func (log *XDefaultLogger) Info(format string, a ...interface{}) {
	log.log(INFO, format, a...)
}

func (log *XDefaultLogger) Warning(format string, a ...interface{}) {
	log.log(WARNING, format, a...)
}

func (log *XDefaultLogger) Error(format string, a ...interface{}) {
	log.log(ERROR, format, a...)
}
