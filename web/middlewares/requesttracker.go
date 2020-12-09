package middlewares

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/google/uuid"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"strconv"
)

const headerXRequestID = "X-Request-ID"

type RequestTrackerMiddleware struct {
	*BaseMiddleware

	logger   xlog.ILogger
	reporter go2sky.Reporter
	tracer   *go2sky.Tracer
}

func NewRequestTracker() *RequestTrackerMiddleware {
	return &RequestTrackerMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func (router *RequestTrackerMiddleware) SetConfiguration(config abstractions.IConfiguration) {
	if config != nil {
		serviceName, _ := config.Get("yoyogo.application.name").(string)
		skyworkingAddr, _ := config.Get("yoyogo.cloud.apm.skyworking.address").(string)
		router.logger = xlog.GetXLogger("Skyworking APM")

		router.reporter, _ = reporter.NewGRPCReporter(skyworkingAddr)
		if router.reporter == nil {
			router.logger.Debug("new reporter error")
		}
		router.tracer, _ = go2sky.NewTracer(serviceName, go2sky.WithReporter(router.reporter))
		if router.tracer == nil {
			router.logger.Debug("new tracer error")
		}

	}
}

//const componentIDGINHttpServer = 5006

func (router *RequestTrackerMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	requestId := ctx.Input.Header(headerXRequestID)
	if requestId == "" {
		requestId = uuid.New().String()
	}
	ctx.Output.Header(headerXRequestID, requestId)

	span, ctxc, _ := router.tracer.CreateEntrySpan(ctx.Input.Request.Context(), ctx.Input.Request.URL.Path, func() (string, error) {
		return ctx.Input.Request.Header.Get(headerXRequestID), nil
	})

	//span.SetComponent(componentIDGINHttpServer)
	span.Tag(go2sky.TagHTTPMethod, ctx.Input.Request.Method)

	span.Tag(go2sky.TagURL, ctx.Input.Request.Host+ctx.Input.Request.URL.Path)
	span.SetSpanLayer(3)

	ctx.Input.Request = ctx.Input.Request.WithContext(ctxc)

	next(ctx)

	span.Tag(go2sky.TagStatusCode, strconv.Itoa(ctx.Output.Status()))
	span.End()
}
