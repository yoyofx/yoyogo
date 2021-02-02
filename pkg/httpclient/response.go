package httpclient

import (
	"net/http"
	"time"
)

type Response struct {
	Body        []byte // response body
	BodyRaw     *http.Response
	RequestTime time.Duration
	CookieData  map[string]*http.Cookie
}

func (c *Response) String() string {
	return string(c.Body)
}

func (c *Response) GetRequestTime() time.Duration {
	return c.RequestTime
}
