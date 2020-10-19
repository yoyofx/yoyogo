package Middleware

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"net/http"
	"os"
	"strings"
)

type StaticOption struct {
	IsPrefix    bool
	WebRoot     string
	VirtualPath string
}

type Static struct {
	Option *StaticOption
}

func NewStatic(patten string, path string) *Static {
	option := &StaticOption{
		IsPrefix:    false,
		WebRoot:     "",
		VirtualPath: patten,
	}
	if path != "" {
		option.IsPrefix = true
		option.WebRoot = path
	}
	return &Static{option}
}

func NewStaticWithConfig(configuration Abstractions.IConfiguration) *Static {
	if configuration != nil {
		config := configuration.GetSection("yoyogo.application.server.static")
		patten := config.Get("patten").(string)
		path := config.Get("webroot").(string)
		return NewStatic(patten, path)
	} else {
		return NewStatic("/", "./Static")
	}
}

func (s *Static) Invoke(ctx *Context.HttpContext, next func(ctx *Context.HttpContext)) {
	if (ctx.Input.Request.Method != "GET" && ctx.Input.Request.Method != "HEAD") || (s.Option.IsPrefix && !strings.Contains(ctx.Input.Request.URL.Path, s.Option.VirtualPath)) {
		next(ctx)
		return
	}

	webrootPath := s.Option.WebRoot
	prefixPath := s.Option.VirtualPath
	if webrootPath == "" {
		webrootPath = "." + "/" + s.Option.VirtualPath
	}
	virtualPath := s.Option.VirtualPath
	if virtualPath != "" {
		virtualPath += "/"
	}
	requestFilePath := webrootPath + strings.Replace(ctx.Input.Request.URL.Path, virtualPath, "", 20)

	exist, err := pathExists(requestFilePath)
	if !exist || err != nil {
		next(ctx)
		return
	}

	staticHandle := http.FileServer(http.Dir(webrootPath))

	if ctx.Input.Request.URL.Path != "/favicon.ico" {
		if s.Option.IsPrefix {
			staticHandle = http.StripPrefix(prefixPath, staticHandle)
		}
	}

	staticHandle.ServeHTTP(ctx.Output.GetWriter(), ctx.Input.GetReader())
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
