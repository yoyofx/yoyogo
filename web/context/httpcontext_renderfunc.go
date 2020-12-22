package context

import (
	"github.com/yoyofx/yoyogo/web/actionresult"
	"net/http"
	"sync"
)

func (ctx *HttpContext) HTML(code int, name string, obj interface{}) {
	htmlRender := actionresult.HTMLDebug{Files: nil,
		Glob:    "../Static/template/**",
		Delims:  actionresult.Delims{Left: "{{", Right: "}}"},
		FuncMap: nil,
	}
	instance := htmlRender.Instance(name, obj)
	_ = instance.Render(ctx.Output.GetWriter())
}

func (ctx *HttpContext) IndentedJSON(code int, obj interface{}) {
	ctx.Render(code, actionresult.IndentedJson{Data: obj})
}

func (ctx *HttpContext) SecureJSON(code int, obj interface{}) {
	ctx.Render(code, actionresult.SecureJson{Prefix: "", Data: obj})
}

func (ctx *HttpContext) JSONP(code int, obj interface{}) {
	callback := ctx.Input.QueryDefault("callback", "")
	if callback == "" {
		ctx.Render(code, actionresult.Json{Data: obj})
		return
	}
	ctx.Render(code, actionresult.Jsonp{Callback: callback, Data: obj})
}

var (
	jsonPool = sync.Pool{
		New: func() interface{} {
			return actionresult.Json{}
		},
	}
)

func (ctx *HttpContext) JSON(code int, obj interface{}) {
	result := jsonPool.Get().(actionresult.Json)
	defer jsonPool.Put(result)
	result.Data = obj
	ctx.Render(code, result)
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func (ctx *HttpContext) AsciiJSON(code int, obj interface{}) {
	ctx.Render(code, actionresult.AsciiJson{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func (ctx *HttpContext) PureJSON(code int, obj interface{}) {
	ctx.Render(code, actionresult.PureJson{Data: obj})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func (ctx *HttpContext) XML(code int, obj interface{}) {
	ctx.Render(code, actionresult.XML{Data: obj})
}

// YAML serializes the given struct as YAML into the response body.
func (ctx *HttpContext) YAML(code int, obj interface{}) {
	ctx.Render(code, actionresult.YAML{Data: obj})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (ctx *HttpContext) ProtoBuf(code int, obj interface{}) {
	ctx.Render(code, actionresult.ProtoBuf{Data: obj})
}

// String writes the given string into the response body.
func (ctx *HttpContext) Text(code int, format string, values ...interface{}) {
	ctx.Render(code, actionresult.Text{Format: format, Data: values})
}

func (ctx *HttpContext) File(filepath string) {
	http.ServeFile(ctx.Output.GetWriter(), ctx.Input.GetReader(), filepath)
}

func (ctx *HttpContext) FileStream(code int, bytes []byte) {
	render := actionresult.FormFileStream(bytes)
	ctx.Render(code, render)
}
