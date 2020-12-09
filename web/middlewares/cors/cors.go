package cors

import (
	"github.com/yoyofx/yoyogo/web/context"
	"net/http"
	"strings"
)

type Cors struct {
	allowAllOrigins  bool
	allowCredentials bool
	allowOriginFunc  func(string) bool
	allowOrigins     []string
	exposeHeaders    []string
	normalHeaders    http.Header
	preflightHeaders http.Header
	wildcardOrigins  [][]string
}

var (
	DefaultSchemas = []string{
		"http://",
		"https://",
	}
	ExtensionSchemas = []string{
		"chrome-extension://",
		"safari-extension://",
		"moz-extension://",
		"ms-browser-extension://",
	}
	FileSchemas = []string{
		"file://",
	}
	WebSocketSchemas = []string{
		"ws://",
		"wss://",
	}
)

func NewCors(config Config) *Cors {
	if err := config.Validate(); err != nil {
		panic(err.Error())
	}

	return &Cors{
		allowOriginFunc:  config.AllowOriginFunc,
		allowAllOrigins:  config.AllowAllOrigins,
		allowCredentials: config.AllowCredentials,
		allowOrigins:     normalize(config.AllowOrigins),
		normalHeaders:    generateNormalHeaders(config),
		preflightHeaders: generatePreflightHeaders(config),
		wildcardOrigins:  config.parseWildcardRules(),
	}
}

func (cors *Cors) ApplyCors(c *context.HttpContext) {
	origin := c.Input.Header("Origin")
	if len(origin) == 0 {
		// request is not a cors request
		return
	}
	host := c.Input.Host()

	if origin == "http://"+host || origin == "https://"+host {
		// request is not a cors request but have origin header.
		// for example, use fetch api
		return
	}

	if !cors.validateOrigin(origin) {
		c.Output.SetStatus(http.StatusForbidden)
		return
	}

	if c.Input.Method() == "OPTIONS" {
		cors.handlePreflight(c)
		defer c.Output.SetStatus(http.StatusNoContent) // Using 204 is better than 200 when the request status is OPTIONS
	} else {
		cors.handleNormal(c)
	}

	if !cors.allowAllOrigins {
		c.Output.Header("Access-Control-Allow-Origin", origin)
	}
}

func (cors *Cors) validateWildcardOrigin(origin string) bool {
	for _, w := range cors.wildcardOrigins {
		if w[0] == "*" && strings.HasSuffix(origin, w[1]) {
			return true
		}
		if w[1] == "*" && strings.HasPrefix(origin, w[0]) {
			return true
		}
		if strings.HasPrefix(origin, w[0]) && strings.HasSuffix(origin, w[1]) {
			return true
		}
	}

	return false
}

func (cors *Cors) validateOrigin(origin string) bool {
	if cors.allowAllOrigins {
		return true
	}
	for _, value := range cors.allowOrigins {
		if value == origin {
			return true
		}
	}
	if len(cors.wildcardOrigins) > 0 && cors.validateWildcardOrigin(origin) {
		return true
	}
	if cors.allowOriginFunc != nil {
		return cors.allowOriginFunc(origin)
	}
	return false
}

func (cors *Cors) handlePreflight(c *context.HttpContext) {
	header := c.Output.Response.Header()
	for key, value := range cors.preflightHeaders {
		header[key] = value
	}
}

func (cors *Cors) handleNormal(c *context.HttpContext) {
	header := c.Output.Response.Header()
	for key, value := range cors.normalHeaders {
		header[key] = value
	}
}
