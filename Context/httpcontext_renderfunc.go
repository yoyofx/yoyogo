package Context

import (
	"github.com/maxzhang1985/yoyogo/ResponseRender"
	"net/http"
)

func (ctx *HttpContext) HTML(code int, name string, obj interface{}) {
	htmlRender := ResponseRender.HTMLDebug{Files: nil,
		Glob:    "../Static/template/**",
		Delims:  ResponseRender.Delims{Left: "{[{", Right: "}]}"},
		FuncMap: nil,
	}
	instance := htmlRender.Instance(name, obj)
	_ = instance.Render(ctx.Resp)
}

func (ctx *HttpContext) IndentedJSON(code int, obj interface{}) {
	ctx.Render(code, ResponseRender.IndentedJson{Data: obj})
}

func (c *HttpContext) SecureJSON(code int, obj interface{}) {
	c.Render(code, ResponseRender.SecureJson{Prefix: "", Data: obj})
}

func (c *HttpContext) JSONP(code int, obj interface{}) {
	callback := c.QueryStringOrDefault("callback", "")
	if callback == "" {
		c.Render(code, ResponseRender.Json{Data: obj})
		return
	}
	c.Render(code, ResponseRender.Jsonp{Callback: callback, Data: obj})
}

func (c *HttpContext) JSON(code int, obj interface{}) {
	c.Render(code, ResponseRender.Json{Data: obj})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func (c *HttpContext) AsciiJSON(code int, obj interface{}) {
	c.Render(code, ResponseRender.AsciiJson{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func (c *HttpContext) PureJSON(code int, obj interface{}) {
	c.Render(code, ResponseRender.PureJson{Data: obj})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func (c *HttpContext) XML(code int, obj interface{}) {
	c.Render(code, ResponseRender.XML{Data: obj})
}

// YAML serializes the given struct as YAML into the response body.
func (c *HttpContext) YAML(code int, obj interface{}) {
	c.Render(code, ResponseRender.YAML{Data: obj})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (c *HttpContext) ProtoBuf(code int, obj interface{}) {
	c.Render(code, ResponseRender.ProtoBuf{Data: obj})
}

// String writes the given string into the response body.
func (c *HttpContext) Text(code int, format string, values ...interface{}) {
	c.Render(code, ResponseRender.Text{Format: format, Data: values})
}

func (c *HttpContext) File(filepath string) {
	http.ServeFile(c.Resp, c.Req, filepath)
}

func (c *HttpContext) FileStream(code int, bytes []byte) {
	render := ResponseRender.FormFileStream(bytes)
	c.Render(code, render)
}
