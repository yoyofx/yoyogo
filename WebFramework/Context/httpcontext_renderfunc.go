package Context

import (
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"net/http"
)

func (ctx *HttpContext) HTML(code int, name string, obj interface{}) {
	htmlRender := ActionResult.HTMLDebug{Files: nil,
		Glob:    "../Static/template/**",
		Delims:  ActionResult.Delims{Left: "{{", Right: "}}"},
		FuncMap: nil,
	}
	instance := htmlRender.Instance(name, obj)
	_ = instance.Render(ctx.Output.GetWriter())
}

func (ctx *HttpContext) IndentedJSON(code int, obj interface{}) {
	ctx.Render(code, ActionResult.IndentedJson{Data: obj})
}

func (ctx *HttpContext) SecureJSON(code int, obj interface{}) {
	ctx.Render(code, ActionResult.SecureJson{Prefix: "", Data: obj})
}

func (ctx *HttpContext) JSONP(code int, obj interface{}) {
	callback := ctx.Input.QueryDefault("callback", "")
	if callback == "" {
		ctx.Render(code, ActionResult.Json{Data: obj})
		return
	}
	ctx.Render(code, ActionResult.Jsonp{Callback: callback, Data: obj})
}

func (ctx *HttpContext) JSON(code int, obj interface{}) {
	ctx.Render(code, ActionResult.Json{Data: obj})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func (ctx *HttpContext) AsciiJSON(code int, obj interface{}) {
	ctx.Render(code, ActionResult.AsciiJson{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func (ctx *HttpContext) PureJSON(code int, obj interface{}) {
	ctx.Render(code, ActionResult.PureJson{Data: obj})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func (ctx *HttpContext) XML(code int, obj interface{}) {
	ctx.Render(code, ActionResult.XML{Data: obj})
}

// YAML serializes the given struct as YAML into the response body.
func (ctx *HttpContext) YAML(code int, obj interface{}) {
	ctx.Render(code, ActionResult.YAML{Data: obj})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (ctx *HttpContext) ProtoBuf(code int, obj interface{}) {
	ctx.Render(code, ActionResult.ProtoBuf{Data: obj})
}

// String writes the given string into the response body.
func (ctx *HttpContext) Text(code int, format string, values ...interface{}) {
	ctx.Render(code, ActionResult.Text{Format: format, Data: values})
}

func (ctx *HttpContext) File(filepath string) {
	http.ServeFile(ctx.Output.GetWriter(), ctx.Input.GetReader(), filepath)
}

func (ctx *HttpContext) FileStream(code int, bytes []byte) {
	render := ActionResult.FormFileStream(bytes)
	ctx.Render(code, render)
}
