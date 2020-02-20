package ws

import (
	"dragon-fruit/app/global"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsConn map[string]*websocket.Conn

func init() {
	wsConn = map[string]*websocket.Conn{}
}

// Wshandler web socket handler
// func Wshandler(c *gin.Context) {
// 	// 賦予用戶唯一 uid
// 	uid := uuid.Must(uuid.NewV4()).String()
// 	fmt.Println("當前用戶 uid ====>", uid)

// // 取 ws 連線
// conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
// if err != nil {
// 	fmt.Println("Failed to set websocket upgrade: ", err)
// 	return
// }

// 	// 連線成功後，存在全域變數作為紀錄
// 	wsConn[uid] = conn

// 	fmt.Println("檢查總用戶連線 ====>", len(wsConn))

// 	for {
// 		// ws 讀取
// 		t, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			// 移除該筆 websocket 連線
// 			delete(wsConn, uid)

// 			fmt.Println("Read Error Msg ===>", err.Error())
// 			fmt.Println("用戶已斷線:", uid)

// 			break
// 		}

// 		fmt.Println(t)
// 		fmt.Println(string(msg))

// 		// ws 回覆
// 		for k := range wsConn {
// 			err = wsConn[k].WriteMessage(t, msg)
// 			if err != nil {
// 				// 移除陣列 websocket 連線
// 				delete(wsConn, uid)

// 				// 移除該筆 websocket 連線
// 				wsConn[k].Close()

// 				fmt.Println("用戶已斷線:", uid)
// 				fmt.Println("Write Error Msg ===>", err.Error())
// 				break
// 			}
// 		}

// 		// err = conn.Close()
// 		// if err != nil {
// 		// 	// 移除該筆 websocket 連線
// 		// 	delete(wsConn, uid)

// 		// 	fmt.Println("Close Error Msg ===>", err.Error())
// 		// 	break
// 		// }

// 		// fmt.Println("連線已中斷")
// 	}
// }

// Wshandler websocket handler
func Wshandler(c *gin.Context) {
	// 取 ws 連線
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}

	uid, _ := uuid.NewV4()
	client := &global.Client{
		ID:     uid.String(),
		Socket: conn,
		Send:   make(chan []byte),
	}

	manager := global.ClientManager{}
	manager.Register <- client

	go client.Read()
	go client.Write()
}
