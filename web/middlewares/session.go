package middlewares

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/session"
	"github.com/yoyofx/yoyogo/web/session/identity"
	"github.com/yoyofx/yoyogo/web/session/store"
)

type SessionMiddleware struct {
	*BaseMiddleware
	sessionMgr   context.ISessionManager
	sessionStore store.ISessionStore
	identity     identity.IProvider
	sessionName  string
	mMaxLifeTime int64
}

func NewSessionWith(provider identity.IProvider, store store.ISessionStore, config abstractions.IConfiguration) *SessionMiddleware {
	sessionTimeout := config.GetInt("yoyogo.application.server.session.timeout")
	if sessionTimeout == 0 {
		sessionTimeout = 3600
	}
	store.SetMaxLifeTime(int64(sessionTimeout))
	mgr := session.NewSessionWithStore(store)
	return &SessionMiddleware{BaseMiddleware: &BaseMiddleware{}, sessionMgr: mgr, identity: provider, sessionStore: store}
}

func (sessionMid *SessionMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	sessionMid.identity.SetContext(ctx)
	sessionId := sessionMid.sessionMgr.Load(sessionMid.identity)
	ctx.SetItem("sessionId", sessionId)
	ctx.SetItem("sessionMgr", sessionMid.sessionMgr)
	next(ctx)
}
