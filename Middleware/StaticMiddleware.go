package Middleware

import (
	"net/http"
	"os"
	"strings"
)

type Static struct {
	IsPrefix   bool
	VirualPath string
}

func NewStatic(path string) *Static {
	return &Static{VirualPath: path, IsPrefix: false}
}

func (s *Static) SetPrefix() {
	s.IsPrefix = true
}

func (s *Static) Inovke(ctx *HttpContext, next func(ctx *HttpContext)) {
	if (ctx.Req.Method != "GET" && ctx.Req.Method != "HEAD") || (s.IsPrefix && !strings.Contains(ctx.Req.URL.Path, s.VirualPath)) {
		next(ctx)
		return
	}

	prefixPath := "/" + s.VirualPath
	webrootPath := "." + "/" + s.VirualPath
	requestFilePath := webrootPath + ctx.Req.URL.Path

	exist, err := pathExists(requestFilePath)
	if !exist || err != nil {
		next(ctx)
		return
	}

	staticHandle := http.FileServer(http.Dir(webrootPath))
	if s.IsPrefix {
		staticHandle = http.StripPrefix(prefixPath, staticHandle)
	}

	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
