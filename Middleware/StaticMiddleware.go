package Middleware

import (
	"net/http"
	"strings"
)

type Static struct {
	VPath string
}

func NewStatic(path string) *Static {
	return &Static{VPath: path}
}

func (s *Static) Inovke(ctx *HttpContext, next func(ctx *HttpContext)) {
	if (ctx.Req.Method != "GET" && ctx.Req.Method != "HEAD") || !strings.Contains(ctx.Req.URL.Path, s.VPath) {
		next(ctx)
		return
	}

	prefixPath := "/" + s.VPath
	webroot := "." + prefixPath
	staticHandle := http.StripPrefix(prefixPath,
		http.FileServer(http.Dir(webroot)))

	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
}
