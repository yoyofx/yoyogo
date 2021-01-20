package contollers

import (
	"SimpleWeb/hubs"
	"github.com/fasthttp/websocket"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
)

// websocket hub
type HubController struct {
	mvc.ApiController

	hub *hubs.Hub
}

func NewHubController(hub *hubs.Hub) *HubController {
	go hub.Run() // run once by async
	return &HubController{hub: hub}
}

// url: ws://localhost:8080/app/v1/hub/ws
func (controller HubController) GetWs(ctx *context.HttpContext) {
	web.Upgrade(ctx, func(conn *websocket.Conn) {
		client := &hubs.Client{Hub: controller.hub, Conn: conn, Send: make(chan []byte, 256)}
		client.Hub.Register <- client
		go client.WritePump()
		client.ReadPump()
	})
}
