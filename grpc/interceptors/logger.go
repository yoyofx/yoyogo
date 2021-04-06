package interceptors

import (
	"context"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/grpc/interceptors/models"
	"google.golang.org/grpc"
	"strconv"
	"strings"
	"sync"
	"time"
)

var LoggerDefaultDateFormat = "2006/01/02 - 15:04:05.00"

type Logger struct {
	infoPool sync.Pool
	log      xlog.ILogger
}

func NewLogger() *Logger {
	pool := sync.Pool{
		New: func() interface{} {
			return models.LoggerInfo{}
		},
	}
	log := xlog.GetXLogger("")
	log.SetCustomLogFormat(nil)
	return &Logger{infoPool: pool, log: log}

}

func (logger *Logger) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()

		resp, err = handler(ctx, req)

		logInfo := logger.infoPool.Get().(models.LoggerInfo)
		logInfo.StartTime = start.Format(LoggerDefaultDateFormat)
		logInfo.Status = 200
		logInfo.Duration = strconv.FormatInt(time.Since(start).Milliseconds(), 10)
		logInfo.HostName = "grpc"
		logInfo.Method = "unary"
		logInfo.Path = info.FullMethod
		logInfo.Request = req
		logInfo.Response = resp

		statusColor := logInfo.StatusCodeColor()
		methodColor := logInfo.MethodColor()
		resetColor := logInfo.ResetColor()
		outLog := fmt.Sprintf("status: %s %3d %s| %5v ms| %s |%s %5s %s | %s | request: %v | response: %v",
			statusColor, logInfo.Status, resetColor,
			logInfo.Duration,
			logInfo.HostName,
			methodColor, logInfo.Method, resetColor,
			logInfo.Path,
			logInfo.Request,
			logInfo.Response,
		)
		logger.infoPool.Put(logInfo)

		logger.log.Info(outLog)

		return resp, err
	}
}

func (logger *Logger) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if strings.HasPrefix(info.FullMethod, "/grpc.reflection") {
			return handler(srv, ss)
		}
		start := time.Now()
		err := handler(srv, ss)

		logInfo := logger.infoPool.Get().(models.LoggerInfo)
		logInfo.StartTime = start.Format(LoggerDefaultDateFormat)
		logInfo.Status = 200
		logInfo.Duration = strconv.FormatInt(time.Since(start).Milliseconds(), 10)
		logInfo.HostName = "grpc"
		logInfo.Method = "stream"
		logInfo.Path = info.FullMethod

		statusColor := logInfo.StatusCodeColor()
		methodColor := logInfo.MethodColor()
		resetColor := logInfo.ResetColor()
		outLog := fmt.Sprintf("status: %s %3d %s| %5v ms| %s |%s %5s %s | %s ",
			statusColor, logInfo.Status, resetColor,
			logInfo.Duration,
			logInfo.HostName,
			methodColor, logInfo.Method, resetColor,
			logInfo.Path,
		)
		logger.infoPool.Put(logInfo)

		logger.log.Info(outLog)

		return err
	}
}
