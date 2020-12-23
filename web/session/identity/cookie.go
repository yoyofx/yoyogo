package identity

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/yoyofx/yoyogo/web/context"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Cookie cookie provider for identity store
type Cookie struct {
	sessionCookieName string
	httpContext       *context.HttpContext
	mMaxLifeTime      int64
}

var defaultLifeTime int64 = 3600

//NewCookie new cookie provider
func NewCookie(cookieName string) *Cookie {
	return &Cookie{
		sessionCookieName: cookieName,
		httpContext:       nil,
		mMaxLifeTime:      defaultLifeTime,
	}
}

func (c *Cookie) SetContext(cxt interface{}) {
	c.httpContext = cxt.(*context.HttpContext)
}

//SetMaxLifeTime set life time
func (c *Cookie) SetMaxLifeTime(liftTime int64) {
	c.mMaxLifeTime = liftTime
}

//SetID set session id
func (c Cookie) SetID(sessionId string) {
	//让浏览器cookie设置过期时间
	cookie := http.Cookie{Name: c.sessionCookieName, Value: sessionId, Path: "/", HttpOnly: true, MaxAge: int(c.mMaxLifeTime)}
	c.httpContext.Output.Response.Header().Add("Set-Cookie", cookie.String())
}

// GetID get session id
func (c Cookie) GetID() string {
	cookie, err := c.httpContext.Input.Request.Cookie(c.sessionCookieName)
	if err != nil || cookie.Value == "" {
		newSessionId := newId()
		c.SetID(newSessionId)
		return newSessionId
	} else {
		return cookie.Value
	}
}

// Clear clear all session
func (c Cookie) Clear() {
	expiration := time.Now()
	cookie := http.Cookie{Name: c.sessionCookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	c.httpContext.Output.Response.Header().Add("Set-Cookie", cookie.String())
}

func newId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano() //微秒
		return strconv.FormatInt(nano, 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}
