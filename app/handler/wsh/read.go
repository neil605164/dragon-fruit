package wsh

import (
	"dragon-fruit/app/business"
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/helper"
	"dragon-fruit/app/global/ws"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ServeWs 请求ping 返回pong
func ServeWs() gin.HandlerFunc {

	hub := ws.NewHub()
	go hub.Run()

	upGrader := websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		// 取消ws跨域校验
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 啟動 redis Sub 監聽
	bus := business.RedisIns()
	go bus.RedisSub()

	return func(c *gin.Context) {

		// 取 token
		token := strings.TrimSpace(c.Query("token"))
		if token == "" {
			apiErr := helper.ErrorHandle(global.WarnLog, 1004002, "TOKEN NOT EXIST", token)
			c.JSON(http.StatusOK, helper.Fail(apiErr))
			return
		}

		// todo 驗證 + 解析 token
		userID := "use01"
		parentID := "parent01"
		gameID := "1"

		conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			conn.Close()
			helper.ErrorHandle(global.WarnLog, 1004001, err.Error())
			return
		}

		client := &ws.Client{
			UID:    userID,
			PID:    parentID,
			Token:  token,
			GameID: gameID,
			Hub:    hub, Conn: conn,
			Send: make(chan []byte, 256)
		}

		client.Hub.Register <- client

		go client.ReadPump()
		go client.WritePump()
	}
}
