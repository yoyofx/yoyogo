package middlewares

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/utils"
	"github.com/yoyofx/yoyogo/utils/jwt"
	"github.com/yoyofx/yoyogo/web/context"
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
	Name string `extension:"name"`
}

func NewJwt() *JwtMiddleware {
	return &JwtMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func (jwtmdw *JwtMiddleware) SetConfiguration(config abstractions.IConfiguration) {
	var hasEnable, hasSecret, hasPrefix, hasHeader bool
	if config != nil {
		jwtmdw.Enable, hasEnable = config.Get("yoyogo.application.server.jwt.enable").(bool)
		jwtmdw.SecretKey, hasSecret = config.Get("yoyogo.application.server.jwt.secret").(string)
		jwtmdw.Prefix, hasPrefix = config.Get("yoyogo.application.server.jwt.prefix").(string)
		jwtmdw.Header, hasHeader = config.Get("yoyogo.application.server.jwt.header").(string)
		jwtmdw.SkipPath, _ = config.Get("yoyogo.application.server.jwt.skip_path").([]interface{})
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

	jwtmdw.SkipPath = append(jwtmdw.SkipPath, "/")
	jwtmdw.SkipPath = append(jwtmdw.SkipPath, "/auth/token")
	jwtmdw.SkipPath = append(jwtmdw.SkipPath, "/actuator/health")

}

func (jwtmdw *JwtMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {

	if !jwtmdw.Enable || utils.Contains(ctx.Input.Path(), jwtmdw.SkipPath) {
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
