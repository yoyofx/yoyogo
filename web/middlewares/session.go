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

type SessionConfig struct {
	Name    string `mapstructure:"name"`
	TimeOut int64  `mapstructure:"timeout"`
}

func NewSessionWith(provider identity.IProvider, store store.ISessionStore, config abstractions.IConfiguration) *SessionMiddleware {
	var sessionConfig *SessionConfig
	config.GetSection("yoyogo.application.server.session").Unmarshal(&sessionConfig)
	if sessionConfig.TimeOut == 0 {
		sessionConfig.TimeOut = 3600
	}
	if sessionConfig.Name != "" {
		provider.SetName(sessionConfig.Name)
	}
	store.SetMaxLifeTime(sessionConfig.TimeOut)
	mgr := session.NewSessionWithStore(store)
	return &SessionMiddleware{BaseMiddleware: &BaseMiddleware{}, sessionMgr: mgr, identity: provider, sessionStore: store}
}

func (sessionMid *SessionMiddleware) Inovke(ctx *context.HttpContext, next func(ctx *context.HttpContext)) {
	sessionMid.identity.SetContext(ctx)
	sessionId := sessionMid.sessionMgr.Load(sessionMid.identity)
	if sessionId != "" {
		ctx.SetItem("sessionId", sessionId)
		ctx.SetItem("sessionMgr", sessionMid.sessionMgr)
	}
	next(ctx)
}
