package Mvc

import (
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"strings"
)

type RouterHandlerResponse struct {
	Result interface{}
}

func (response *RouterHandlerResponse) Callback(ctx *Context.HttpContext) {
	if actionResult, ok := response.Result.(ActionResult.IActionResult); ok {
		ctx.Render(200, actionResult)
	} else {
		contentType := ctx.Input.Request.Header.Get(Context.HeaderContentType)
		switch {
		case strings.HasPrefix(contentType, Context.MIMEApplicationXML):
			ctx.XML(200, response.Result)
		case strings.HasPrefix(contentType, Context.MIMEApplicationYAML):
			ctx.YAML(200, response.Result)
		case strings.HasPrefix(contentType, Context.MIMEApplicationProtobuf):
			ctx.ProtoBuf(200, response.Result)
		case strings.HasPrefix(contentType, Context.MIMEApplicationJSON):
			fallthrough
		default:
			ctx.JSON(200, response.Result)
		}

	}
}
