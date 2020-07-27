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
	_ = instance.Render(ctx.Response)
}

func (ctx *HttpContext) IndentedJSON(code int, obj interface{}) {
	ctx.Render(code, ActionResult.IndentedJson{Data: obj})
}

func (c *HttpContext) SecureJSON(code int, obj interface{}) {
	c.Render(code, ActionResult.SecureJson{Prefix: "", Data: obj})
}

func (c *HttpContext) JSONP(code int, obj interface{}) {
	callback := c.QueryStringOrDefault("callback", "")
	if callback == "" {
		c.Render(code, ActionResult.Json{Data: obj})
		return
	}
	c.Render(code, ActionResult.Jsonp{Callback: callback, Data: obj})
}

func (c *HttpContext) JSON(code int, obj interface{}) {
	c.Render(code, ActionResult.Json{Data: obj})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func (c *HttpContext) AsciiJSON(code int, obj interface{}) {
	c.Render(code, ActionResult.AsciiJson{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func (c *HttpContext) PureJSON(code int, obj interface{}) {
	c.Render(code, ActionResult.PureJson{Data: obj})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func (c *HttpContext) XML(code int, obj interface{}) {
	c.Render(code, ActionResult.XML{Data: obj})
}

// YAML serializes the given struct as YAML into the response body.
func (c *HttpContext) YAML(code int, obj interface{}) {
	c.Render(code, ActionResult.YAML{Data: obj})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (c *HttpContext) ProtoBuf(code int, obj interface{}) {
	c.Render(code, ActionResult.ProtoBuf{Data: obj})
}

// String writes the given string into the response body.
func (c *HttpContext) Text(code int, format string, values ...interface{}) {
	c.Render(code, ActionResult.Text{Format: format, Data: values})
}

func (c *HttpContext) File(filepath string) {
	http.ServeFile(c.Response, c.Request, filepath)
}

func (c *HttpContext) FileStream(code int, bytes []byte) {
	render := ActionResult.FormFileStream(bytes)
	c.Render(code, render)
}
