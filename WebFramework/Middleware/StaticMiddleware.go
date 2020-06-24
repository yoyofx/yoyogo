package Middleware

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
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

func (s *Static) Inovke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
	if (ctx.Request.Method != "GET" && ctx.Request.Method != "HEAD") || (s.IsPrefix && !strings.Contains(ctx.Request.URL.Path, s.VirualPath)) {
		next(ctx)
		return
	}

	prefixPath := "/" + s.VirualPath
	webrootPath := "." + "/" + s.VirualPath
	requestFilePath := webrootPath + ctx.Request.URL.Path

	exist, err := pathExists(requestFilePath)
	if !exist || err != nil {
		next(ctx)
		return
	}

	staticHandle := http.FileServer(http.Dir(webrootPath))

	if ctx.Request.URL.Path != "/favicon.ico" {
		if s.IsPrefix {
			staticHandle = http.StripPrefix(prefixPath, staticHandle)
		}
	}

	staticHandle.ServeHTTP(ctx.Response, ctx.Request)
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
