package test

import (
	"dragon-fruit/app/business"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ServeWs 请求ping 返回pong
func ServeWs(hub *Hub, c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	token := c.Query("token")

	client := &Client{id: token, hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}

// RedisPub redis sub
func RedisPub(c *gin.Context) {
	bus := business.RedisIns()
	bus.RedisPub()
}

// RedisSub redis sub
func RedisSub(c *gin.Context) {
	bus := business.RedisIns()
	bus.RedisSub()
}
