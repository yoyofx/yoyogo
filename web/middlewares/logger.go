package middlewares

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LoggerInfo) StatusCodeColor() string {
	code := p.Status

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (p *LoggerInfo) MethodColor() string {
	method := p.Method

	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (p *LoggerInfo) ResetColor() string {
	return reset
}

type LoggerInfo struct {
	StartTime string
	Status    int
	Duration  string
	HostName  string
	Method    string
	Path      string
	Body      string
	Request   *http.Request
}

var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} \t {{.Duration}} | {{.HostName}} | {{.Method}} | {{.Path}} "

var LoggerDefaultDateFormat = "2006/01/02 - 15:04:05.00"

type Logger struct {
	ALogger    xlog.ILogger
	dateFormat string
	template   *template.Template
}

func (l *Logger) SetFormat(format string) {
	l.template = template.Must(template.New("yoyofx_parser").Parse(format))
}

func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

func NewLogger() *Logger {
	log := xlog.GetXLogger("Request Log")
	log.SetCustomLogFormat(nil)
	logger := &Logger{ALogger: log, dateFormat: LoggerDefaultDateFormat}
	logger.SetFormat(LoggerDefaultFormat)
	return logger
}

func (l *Logger) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	enEscapeUrl, _ := url.QueryUnescape(ctx.Input.Request.URL.RequestURI())
	l.ALogger.Info(fmt.Sprintf("[yoyogo] Method: %s Url: %s  Content-Type: %s , Accept: %s",
		ctx.Input.Method(), enEscapeUrl,
		ctx.Input.Header(context.HeaderContentType),
		ctx.Input.Header(context.HeaderAccept)))

	start := time.Now()
	next(ctx)
	res := ctx.Output.GetWriter()

	strBody := ""
	bodyFormat := "%s"
	if ctx.Input.Request.Method == "POST" {
		body := ctx.Input.GetBody()
		strBody = string(body[:])
		bodyFormat = "\n%s"
	}

	logInfo := LoggerInfo{
		StartTime: start.Format(l.dateFormat),
		Status:    res.Status(),
		Duration:  strconv.FormatInt(time.Since(start).Milliseconds(), 10),
		HostName:  ctx.Input.Request.Host,
		Method:    ctx.Input.Request.Method,
		Path:      enEscapeUrl,
		Body:      strBody,
	}

	statusColor := logInfo.StatusCodeColor()
	methodColor := logInfo.MethodColor()
	resetColor := logInfo.ResetColor()
	outLog := fmt.Sprintf(consolecolors.Yellow("[yoyogo] ")+"%v |%s %3d %s| %7v ms| %15s |%s %5s %s %s "+bodyFormat,
		logInfo.StartTime,
		statusColor, logInfo.Status, resetColor,
		logInfo.Duration,
		logInfo.HostName,
		methodColor, logInfo.Method, resetColor,
		logInfo.Path,
		logInfo.Body,
	)

	l.ALogger.Info(outLog)

}
