package scheduler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yoyofx/yoyogo/utils"
	"os"
	"strconv"
)

type (
	Logger interface {
		Info(format string, args ...interface{})
		Error(args ...interface{})
		Done(msg string) string
	}

	XxlJobLogger struct {
		logger *logrus.Logger
	}
)

func NewXxlJobLogger(logID int64) (*XxlJobLogger, error) {
	logger, err := GetLogger(logID)
	if err != nil {
		return nil, err
	}
	return &XxlJobLogger{logger: logger}, nil
}

func (x *XxlJobLogger) Info(format string, args ...interface{}) {
	x.logger.Infof(format, args...)
	fmt.Println(fmt.Sprintf(format, args...))
}

func (x *XxlJobLogger) Error(args ...interface{}) {
	fmt.Println(args)
}

func (x *XxlJobLogger) Done(msg string) string {
	x.logger.Info("xxljob-done")
	fmt.Println("job done " + msg)
	return msg
}

func GetLogger(logID int64) (*logrus.Logger, error) {
	err := utils.CreateDir("logs")
	if err != nil {
		return nil, err
	}
	f, err := getLogFile(logID, true)
	if err != nil {
		logrus.Error(err)
	}
	logger := logrus.New()
	logger.SetOutput(f)
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	return logger, err
}

func getLogFile(logID int64, appendMod bool) (*os.File, error) {
	fileFlag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	if !appendMod {
		fileFlag = os.O_RDONLY
	}
	return os.OpenFile("logs/"+strconv.FormatInt(logID, 10)+".log", fileFlag, 0755)
}
