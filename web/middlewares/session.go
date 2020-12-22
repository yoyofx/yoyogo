package middlewares

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/session"
	"github.com/yoyofx/yoyogo/web/session/identity"
)

type SessionMiddleware struct {
	*BaseMiddleware
	sessionMgr   session.IManager
	mMaxLifeTime int64
}

func NewSession() *SessionMiddleware {
	return &SessionMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func (sessionMid *SessionMiddleware) SetConfiguration(config abstractions.IConfiguration) {
	sessionTimeout := config.GetInt("yoyogo.application.server.session_timeout")
	if sessionTimeout == 0 {
		sessionTimeout = 3600
	}
	sessionMid.sessionMgr = session.NewSession(int64(sessionTimeout))
}

func (sessionMid *SessionMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	sessionId := sessionMid.sessionMgr.Load(identity.NewCookie(ctx, "YOYOGOSESSIONID"))
	ctx.SetItem("sessionId", sessionId)
	next(ctx)
}
