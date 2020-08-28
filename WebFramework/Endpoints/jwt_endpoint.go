package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/Utils/jwt"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Middleware"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
	"strconv"
	"time"
)

func UseJwt(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded jwt endpoint.")
	config := router.GetConfiguration()
	var secretKey string
	var expires int64
	var hasSecret, hasExpires bool
	if config != nil {
		secretKey, hasSecret = config.Get("yoyogo.application.server.jwt.secret").(string)
		expires, hasExpires = config.Get("yoyogo.application.server.jwt.expires").(int64)
	}
	if !hasSecret {
		secretKey = "12391JdeOW^%$#@"
	}
	if !hasExpires {
		expires = 3
	}
	if config != nil {
		router.POST("/auth/token", func(ctx *Context.HttpContext) {
			name := ctx.Input.Param("name")
			id := ctx.Input.Param("id")
			if name == "" || id == "" {
				request := &Middleware.JwtRequest{}
				err := ctx.Bind(request)
				if err == nil {
					id = request.Id
					name = request.Name
				}
			}
			if name == "" || id == "" {
				XLog.GetXLogger("Jwt Endpoint").Debug("Create Token: name: %s , id: %v , token: %s")
				ctx.JSON(200, Context.H{
					"token":   "",
					"expires": 0,
					"success": false,
				})
				return
			}

			uid, _ := strconv.Atoi(id)
			token, expires := jwt.CreateToken([]byte(secretKey), name, uint(uid), int64(time.Now().Add(time.Hour*time.Duration(expires)).Unix()))
			XLog.GetXLogger("Jwt Endpoint").Debug("Create Token: ( name: %s , id: %s , token: %s )", name, id, token)
			ctx.JSON(200, Context.H{
				"token":   token,
				"expires": expires,
				"success": true,
			})
		})
	} else {
		XLog.GetXLogger("Jwt Endpoint").Error("config load error.")
	}

}
