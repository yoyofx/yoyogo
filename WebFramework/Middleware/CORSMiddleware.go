package Middleware

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Middleware/CORS"
)

type CORSMiddleware struct {
	*BaseMiddleware

	mCors  *CORS.Cors
	Enable bool
}

func NewCORS() *CORSMiddleware {

	return &CORSMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func (cors *CORSMiddleware) SetConfiguration(config Abstractions.IConfiguration) {
	if config != nil {
		cors.Enable, _ = config.Get("application.server.cors.enable").(bool)
	}
	if cors.Enable {
		corsConfig := CORS.DefaultConfig()
		allowOrigins, _ := config.Get("application.server.cors.allow_origins").([]interface{})
		if allowOrigins != nil {
			for _, ao := range allowOrigins {
				corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, ao.(string))
			}
			//corsConfig.AllowOrigins = allowOrigins
		}
		var ams []string
		allowMethods, _ := config.Get("application.server.cors.allow_methods").([]interface{})
		if allowMethods != nil {
			for _, am := range allowMethods {
				ams = append(ams, am.(string))
			}
			corsConfig.AllowMethods = ams
		}
		allowCredentials, _ := config.Get("application.server.cors.allow_credentials").(bool)
		if allowMethods != nil {
			corsConfig.AllowCredentials = allowCredentials
		}
		cors.mCors = CORS.NewCors(corsConfig)
	}
}

func (cors *CORSMiddleware) Inovke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
	if cors.Enable {
		cors.mCors.ApplyCors(ctx)
	}
	next(ctx)

}
