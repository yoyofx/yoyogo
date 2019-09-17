package Middleware

import (
	"bytes"
	Std "github.com/maxzhang1985/yoyogo/Standard"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type LoggerInfo struct {
	StartTime string
	Status    int
	Duration  string
	HostName  string
	Method    string
	Path      string
	Request   *http.Request
}

var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} \t {{.Duration}} | {{.HostName}} | {{.Method}} | {{.Path}} "

var LoggerDefaultDateFormat = time.ANSIC

type ALogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type Logger struct {
	ALogger    ALogger
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
	logger := &Logger{ALogger: log.New(os.Stdout, "[yoyofx] ", 0), dateFormat: LoggerDefaultDateFormat}
	logger.SetFormat(LoggerDefaultFormat)
	return logger
}

func (l *Logger) Inovke(ctx *HttpContext, next func(ctx *HttpContext)) {
	start := time.Now()
	next(ctx)
	res := ctx.Resp

	log := LoggerInfo{
		StartTime: start.Format(l.dateFormat),
		Status:    res.Status(),
		Duration:  Std.PadLeft(time.Since(start).String(), " ", 11),
		HostName:  ctx.Req.Host,
		Method:    ctx.Req.Method,
		Path:      ctx.Req.URL.Path,
		Request:   ctx.Req,
	}

	buff := &bytes.Buffer{}
	_ = l.template.Execute(buff, log)

	l.ALogger.Println(buff.String())

}
