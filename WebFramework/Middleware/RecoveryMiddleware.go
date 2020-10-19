package Middleware

import (
	"fmt"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
)

const (
	// NoPrintStackBodyString is the body content returned when HTTP stack printing is suppressed
	NoPrintStackBodyString = "500 Internal Server Error"

	panicText = `{ "Result": -1 , "Data": "error", "Message": "%s" }`

	panicHTML = `<html>
<head><title>PANIC: {{.RecoveredPanic}}</title></head>
<style type="text/css">
html, body {
	font-family: Helvetica, Arial, Sans;
	color: #333333;
	background-color: #ffffff;
	margin: 0px;
}
h1 {
	color: #ffffff;
	background-color: #f14c4c;
	padding: 20px;
	border-bottom: 1px solid #2b3848;
}
.block {
	margin: 2em;
}
.panic-interface {
}

.panic-stack-raw pre {
	padding: 1em;
	background: #f6f8fa;
	border: dashed 1px;
}
.panic-interface-title {
	font-weight: bold;
}
</style>
<body>
<h1>yoyogo - PANIC</h1>

<div class="panic-interface block">
	<h3>{{.RequestDescription}}</h3>
	<span class="panic-interface-title">Runtime error:</span> <span class="panic-interface-element">{{.RecoveredPanic}}</span>
</div>

{{ if .Stack }}
<div class="panic-stack-raw block">
	<h3>Runtime Stack</h3>
	<pre>{{.StackAsString}}</pre>
</div>
{{ end }}

</body>
</html>`
	nilRequestMessage = "Request is nil"
)

var panicHTMLTemplate = template.Must(template.New("PanicPage").Parse(panicHTML))

// PanicInformation contains all
// elements for printing stack informations.
type PanicInformation struct {
	RecoveredPanic interface{}
	Stack          []byte
	Request        *http.Request
}

// StackAsString returns a printable version of the stack
func (p *PanicInformation) StackAsString() string {
	return string(p.Stack)
}

// RequestDescription returns a printable description of the url
func (p *PanicInformation) RequestDescription() string {

	if p.Request == nil {
		return nilRequestMessage
	}

	var queryOutput string
	if p.Request.URL.RawQuery != "" {
		queryOutput = "?" + p.Request.URL.RawQuery
	}
	return fmt.Sprintf("%s %s%s", p.Request.Method, p.Request.URL.Path, queryOutput)
}

// PanicFormatter is an interface on object can implement
// to be able to output the stack trace
type PanicFormatter interface {
	// FormatPanicError output the stack for a given answer/response.
	// In case the the middleware should not output the stack trace,
	// the field `Stack` of the passed `PanicInformation` instance equals `[]byte{}`.
	FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *PanicInformation)
}

// TextPanicFormatter output the stack
// as simple text on os.Stdout. If no `Content-Type` is set,
// it will output the data as `text/plain; charset=utf-8`.
// Otherwise, the origin `Content-Type` is kept.
type TextPanicFormatter struct{}

func (t *TextPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *PanicInformation) {
	if rw.Header().Get("Content-Type") == "" {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	//infos.Stack
	_, _ = fmt.Fprintf(rw, panicText, infos.RecoveredPanic)
}

// HTMLPanicFormatter output the stack inside
// an HTML page. This has been largely inspired by
// https://github.com/go-martini/martini/pull/156/commits.
type HTMLPanicFormatter struct{}

func (t *HTMLPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *PanicInformation) {
	if rw.Header().Get("Content-Type") == "" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	_ = panicHTMLTemplate.Execute(rw, infos)

}

type Recovery struct {
	Logger           ALogger
	PrintStack       bool
	LogStack         bool
	PanicHandlerFunc func(*PanicInformation)
	StackAll         bool
	StackSize        int
	Formatter        PanicFormatter
}

// NewRecovery returns a new instance of Recovery
func NewRecovery() *Recovery {
	return &Recovery{
		Logger:     log.New(os.Stdout, "[yoyogo] ", 0),
		PrintStack: true,
		LogStack:   true,
		StackAll:   false,
		StackSize:  1024 * 8,
		Formatter:  &TextPanicFormatter{},
	}
}

func (rec *Recovery) Invoke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {

	defer func() {
		if err := recover(); err != nil {
			var hostEnv *Context.HostEnvironment
			envErr := ctx.RequiredServices.GetService(&hostEnv)

			if envErr == nil && hostEnv.IsDevelopment() {
				rec.PrintStack = true
				rec.LogStack = true
				rec.StackAll = true
				rec.Formatter = &HTMLPanicFormatter{}
			}
			if ctx.Output.Status() != http.StatusNotFound {
				ctx.Output.SetStatus(http.StatusInternalServerError)
			}
			stack := make([]byte, rec.StackSize)
			if rec.StackAll {
				stack = stack[:runtime.Stack(stack, rec.StackAll)]
			}
			infos := &PanicInformation{RecoveredPanic: err, Request: ctx.Input.Request}

			// PrintStack will write stack trace info to the IResponseWriter if set to true!
			if rec.LogStack {
				infos.Stack = stack
				var msg string
				//print console stack errors
				if !rec.StackAll {
					infos.Stack = nil
					msg = string(stack)
					rec.Logger.Printf(panicText, err)
				} else {
					rec.Logger.Printf(panicText, err, msg)
				}

			}

			if rec.PrintStack {
				rec.Formatter.FormatPanicError(ctx.Output.GetWriter(), ctx.Input.GetReader(), infos)
			} else {
				if ctx.Input.Header("Content-Type") == "" {
					ctx.Output.Header("Content-Type", "text/plain; charset=utf-8")
				}
				_, _ = fmt.Fprint(ctx.Output.GetWriter(), NoPrintStackBodyString)
			}

			if rec.PanicHandlerFunc != nil {
				func() {
					defer func() {
						if err := recover(); err != nil {
							rec.Logger.Printf("provided PanicHandlerFunc panic'd: %s, trace:\n%s", err, debug.Stack())
							rec.Logger.Printf("%s\n", debug.Stack())
						}
					}()
					rec.PanicHandlerFunc(infos)
				}()
			}

		}
	}()

	next(ctx)
}
