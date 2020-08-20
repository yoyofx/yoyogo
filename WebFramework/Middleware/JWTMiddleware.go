package Middleware

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Utils"
	"github.com/yoyofx/yoyogo/Utils/jwt"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net/http"
)

type JwtMiddleware struct {
	*BaseMiddleware

	Enable    bool
	SecretKey string
	Prefix    string
	Header    string
	SkipPath  []interface{}
}

type JwtRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewJwt() *JwtMiddleware {
	return &JwtMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func (jwtmdw *JwtMiddleware) SetConfiguration(config Abstractions.IConfiguration) {
	var hasEnable, hasSecret, hasPrefix, hasHeader bool
	if config != nil {
		jwtmdw.Enable, hasEnable = config.Get("application.server.jwt.enable").(bool)
		jwtmdw.SecretKey, hasSecret = config.Get("application.server.jwt.secret").(string)
		jwtmdw.Prefix, hasPrefix = config.Get("application.server.jwt.prefix").(string)
		jwtmdw.Header, hasHeader = config.Get("application.server.jwt.header").(string)
		jwtmdw.SkipPath, _ = config.Get("application.server.jwt.skip_path").([]interface{})
	}

	if !hasEnable {
		jwtmdw.Enable = false
	}

	if !hasSecret {
		jwtmdw.SecretKey = "12391JdeOW^%$#@"
	}
	if !hasPrefix {
		jwtmdw.Prefix = "Bearer"
	}
	if !hasHeader {
		jwtmdw.Header = "Authorization"
	}

	jwtmdw.SkipPath = append(jwtmdw.SkipPath, "/auth/token")

}

func (jwtmdw *JwtMiddleware) Inovke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {

	if !jwtmdw.Enable || Utils.Contains(ctx.Input.Path(), jwtmdw.SkipPath) {
		next(ctx)
		return
	}
	auth := ctx.Input.Header(jwtmdw.Header)
	if auth == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		return
	}
	token := auth[len(jwtmdw.Prefix)+1:]
	info, err := jwt.ParseToken(token, []byte(jwtmdw.SecretKey))

	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Error(http.StatusUnauthorized, "Unauthorized")
		return
	} else {
		mapClaims := info.(jwt.MapClaims)
		userInfo := make(map[string]interface{})
		for k, v := range mapClaims {
			userInfo[k] = v
		}
		ctx.SetItem("userinfo", userInfo)
		next(ctx)
	}

}
