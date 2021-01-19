package web

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"github.com/yoyofx/yoyogo/web/context"
	"log"
	"net/http"
)

var httpUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var fasthttpUpgrader = websocket.FastHTTPUpgrader{
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func UpgradeHandler(upgraderFunc func(*websocket.Conn)) func(ctx *context.HttpContext) {
	return func(ctx *context.HttpContext) {
		Upgrade(ctx, upgraderFunc)
	}
}

// serveWs handles socket requests from the peer.
func Upgrade(ctx *context.HttpContext, upgraderFunc func(*websocket.Conn)) {
	if upgraderFunc == nil {
		panic("Upgrade func is not nil !")
	}
	w := ctx.Output.GetWriter()
	r := ctx.Input.GetReader()
	writer := w.(*context.CResponseWriter)
	responseWriter, ok := writer.ResponseWriter.(*NetHTTPResponseWriter)

	var err error
	if !ok { //http
		conn, e := httpUpgrader.Upgrade(w, r, nil)
		err = e
		upgraderFunc(conn)
	} else { //fasthttp
		responseWriter.IsHijackerConn = true
		err = fasthttpUpgrader.Upgrade(responseWriter.Ctx, func(conn *websocket.Conn) {
			upgraderFunc(conn)
		})
	}
	statusCode := 200
	if err != nil {
		log.Println(err)
		statusCode = 500
	}
	ctx.Output.SetStatus(statusCode)
}
