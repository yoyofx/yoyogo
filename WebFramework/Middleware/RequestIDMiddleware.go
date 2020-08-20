package Middleware

import (
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

const headerXRequestID = "X-Request-ID"

type RequestIDMiddleware struct {
}

func NewRequestID() *RequestIDMiddleware {
	return &RequestIDMiddleware{}
}

func (router *RequestIDMiddleware) Inovke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
	requestId := ctx.Input.Header(headerXRequestID)
	if requestId == "" {
		requestId = uuid.New().String()
	}
	ctx.Output.Header(headerXRequestID, requestId)
	next(ctx)
}
