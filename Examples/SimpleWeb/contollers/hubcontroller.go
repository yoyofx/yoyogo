package contollers

import (
	"SimpleWeb/hubs"
	"github.com/fasthttp/websocket"
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
	redisdb "github.com/yoyofx/yoyogo/pkg/datasources/redis"
	"github.com/yoyofx/yoyogo/web"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
)

// websocket hub
type HubController struct {
	mvc.ApiController

	hub         *hubs.Hub
	redisClient *redisdb.RedisDataSource
}

func NewHubController(redis *redisdb.RedisDataSource, hub *hubs.Hub) *HubController {
	go hub.Run() // run once by async
	return &HubController{hub: hub, redisClient: redis}
}

// url: ws://localhost:8080/app/v1/hub/ws
func (controller HubController) GetWs(ctx *context.HttpContext) {
	web.Upgrade(ctx, func(conn *websocket.Conn) {
		client := &hubs.Client{Hub: controller.hub, Conn: conn, Send: make(chan []byte, 256), MaxMessageSize: 65535}
		client.Hub.Register <- client
		go client.WritePump()
		client.ReadPump()
	})
}

func (controller HubController) GetTodoList() actionresult.IActionResult {
	conn, _, _ := controller.redisClient.Open()
	client := conn.(redis.IClient)
	json, _ := client.GetKVOps().GetString("yoyogo:todolist")
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        []byte(json),
	}
}

func (controller HubController) PostTodoSync(ctx *context.HttpContext) mvc.ApiResult {
	json := string(ctx.Input.FormBody)
	conn, _, _ := controller.redisClient.Open()
	client := conn.(redis.IClient)
	_ = client.GetKVOps().SetString("yoyogo:todolist", json, 0)

	return controller.OK("ok")
}
