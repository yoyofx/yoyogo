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
	statusCode := 200
	if stater, ok := response.Result.(StatusCoder); ok {
		code := stater.StatusCode()
		if code > 0 {
			statusCode = code
		}
	}

	if actionResult, ok := response.Result.(actionresult.IActionResult); ok {
		ctx.Render(statusCode, actionResult)
	} else {
		contentType := ctx.Input.Request.Header.Get(context.HeaderContentType)
		switch {
		case strings.HasPrefix(contentType, context.MIMEApplicationXML):
			ctx.XML(statusCode, response.Result)
		case strings.HasPrefix(contentType, context.MIMEApplicationYAML):
			ctx.YAML(statusCode, response.Result)
		case strings.HasPrefix(contentType, context.MIMEApplicationProtobuf):
			ctx.ProtoBuf(statusCode, response.Result)
		case strings.HasPrefix(contentType, context.MIMEApplicationJSON):
			fallthrough
		default:
			ctx.JSON(statusCode, response.Result)
		}

	}
}
