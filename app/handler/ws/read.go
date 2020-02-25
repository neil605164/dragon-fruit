// package ws

// import (
// 	"dragon-fruit/app/business/gameb"
// 	"dragon-fruit/app/global"
// 	"dragon-fruit/app/global/helper"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// 	uuid "github.com/satori/go.uuid"
// )

// var wsupgrader = websocket.Upgrader{
// 	ReadBufferSize:   1024,
// 	WriteBufferSize:  1024,
// 	HandshakeTimeout: 5 * time.Second,
// 	// 取消ws跨域校验
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// var wsConn map[string]*websocket.Conn

// // Wshandler web socket handler
// func Wshandler() gin.HandlerFunc {

// 	// 初始化
// 	wsConn = map[string]*websocket.Conn{}

// 	return func(c *gin.Context) {
// 		// 取 ws 連線
// 		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
// 		if err != nil {
// 			helper.ErrorHandle(global.WarnLog, 1004001, err.Error())
// 			return
// 		}

// 		// 賦予用戶唯一 uid + 存在全域變數作為紀錄
// 		uid := uuid.Must(uuid.NewV4()).String()
// 		wsConn[uid] = conn

// 		defer func() {

// 			fmt.Println("Defer here and uid is ", uid)

// 			delete(wsConn, uid)
// 			conn.Close()
// 		}()

// 		// 取 token
// 		token := c.Query("token")

// 		bus := gameb.Instance()
// 		bus.EnterGame(conn, token)

// 		for {
// 			// ws 讀取
// 			t, msg, err := conn.ReadMessage()
// 			if err != nil {
// 				// 移除該筆 websocket 連線
// 				delete(wsConn, uid)

// 				fmt.Println("Read Error Msg ===>", err.Error())
// 				fmt.Println("用戶已斷線:", uid)

// 				break
// 			}

// 			fmt.Println(t)
// 			fmt.Println(string(msg))

// 			// ws 回覆
// 			for k := range wsConn {
// 				err = wsConn[k].WriteMessage(t, msg)
// 				if err != nil {
// 					// 移除陣列 websocket 連線
// 					delete(wsConn, uid)

// 					// 移除該筆 websocket 連線
// 					wsConn[k].Close()

// 					fmt.Println("用戶已斷線:", uid)
// 					fmt.Println("Write Error Msg ===>", err.Error())
// 					break
// 				}
// 			}

// 			// err = conn.Close()
// 			// if err != nil {
// 			// 	// 移除該筆 websocket 連線
// 			// 	delete(wsConn, uid)

// 			// 	fmt.Println("Close Error Msg ===>", err.Error())
// 			// 	break
// 			// }

// 			// fmt.Println("連線已中斷")
// 		}
// 	}
// }
