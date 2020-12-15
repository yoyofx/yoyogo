package middlewares

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/middlewares/cors"
)

type CORSMiddleware struct {
	*BaseMiddleware

	mCors  *cors.Cors
	Enable bool
}

func NewCORS() *CORSMiddleware {

	return &CORSMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func (corsmw *CORSMiddleware) SetConfiguration(config abstractions.IConfiguration) {
	if config != nil {
		c := config.Get("yoyogo.application.server.cors")
		corsmw.Enable = c != nil
	}
	if corsmw.Enable {
		corsConfig := cors.DefaultConfig()
		allowOrigins, _ := config.Get("yoyogo.application.server.cors.allow_origins").([]interface{})
		if allowOrigins != nil {
			for _, ao := range allowOrigins {
				corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, ao.(string))
			}
			//corsConfig.AllowOrigins = allowOrigins
		}
		var ams []string
		allowMethods, _ := config.Get("yoyogo.application.server.cors.allow_methods").([]interface{})
		if allowMethods != nil {
			for _, am := range allowMethods {
				ams = append(ams, am.(string))
			}
			corsConfig.AllowMethods = ams
		}
		allowCredentials, _ := config.Get("yoyogo.application.server.cors.allow_credentials").(bool)
		if allowMethods != nil {
			corsConfig.AllowCredentials = allowCredentials
		}
		corsmw.mCors = cors.NewCors(corsConfig)
	}
}

func (corsmw *CORSMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	if corsmw.Enable {
		corsmw.mCors.ApplyCors(ctx)
	}
	next(ctx)

}
