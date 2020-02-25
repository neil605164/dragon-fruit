package gameb

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

// Business GameList 共用 Business
type Business struct {
}

var singleton *Business
var once sync.Once

// Instance 獲得單例對象
func Instance() *Business {
	once.Do(func() {
		singleton = &Business{}
	})
	return singleton
}

// EnterGame 進入遊戲
func (b *Business) EnterGame(conn *websocket.Conn, token string) {
	// Token 驗證

	// 讀取 WebSocket 內容
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {

			break
		}

		fmt.Println(t)
		fmt.Println(string(msg))

		// todo
	}

}
