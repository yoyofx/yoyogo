package mvc

import (
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/context"
	"strings"
)

type RouterHandlerResponse struct {
	Result interface{}
}

func (response *RouterHandlerResponse) Callback(ctx *context.HttpContext) {
	if actionResult, ok := response.Result.(actionresult.IActionResult); ok {
		ctx.Render(200, actionResult)
	} else {
		contentType := ctx.Input.Request.Header.Get(context.HeaderContentType)
		switch {
		case strings.HasPrefix(contentType, context.MIMEApplicationXML):
			ctx.XML(200, response.Result)
		case strings.HasPrefix(contentType, context.MIMEApplicationYAML):
			ctx.YAML(200, response.Result)
		case strings.HasPrefix(contentType, context.MIMEApplicationProtobuf):
			ctx.ProtoBuf(200, response.Result)
		case strings.HasPrefix(contentType, context.MIMEApplicationJSON):
			fallthrough
		default:
			ctx.JSON(200, response.Result)
		}

	}
}
