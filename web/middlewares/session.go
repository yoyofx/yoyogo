package middlewares

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/session"
	"github.com/yoyofx/yoyogo/web/session/identity"
)

type SessionMiddleware struct {
	*BaseMiddleware
	sessionMgr   session.IManager
	sessionName  string
	mMaxLifeTime int64
}

func NewSession() *SessionMiddleware {
	return &SessionMiddleware{BaseMiddleware: &BaseMiddleware{}}
}

func RegisterCookieSession() {
	abstractions.RegisterConfigurationProcessor(
		func(config abstractions.IConfiguration, serviceCollection *dependencyinjection.ServiceCollection) {
			serviceCollection.AddSingletonByImplements(NewSession, new(abstractions.IDataSource))
		})
}

func (sessionMid *SessionMiddleware) SetConfiguration(config abstractions.IConfiguration) {
	sessionTimeout := config.GetInt("yoyogo.application.server.session.timeout")
	sessionName := config.GetString("yoyogo.application.server.session.name")
	if sessionTimeout == 0 {
		sessionTimeout = 3600
	}
	if sessionName == "" {
		sessionName = "YOYOGOSESSIONID"
	}
	sessionMid.sessionName = sessionName
	sessionMid.sessionMgr = session.NewSession(int64(sessionTimeout))
}

func (sessionMid *SessionMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	sessionId := sessionMid.sessionMgr.Load(identity.NewCookie(ctx, sessionMid.sessionName))
	ctx.SetItem("sessionId", sessionId)
	next(ctx)
}
