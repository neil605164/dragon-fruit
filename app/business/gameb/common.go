package gameb

import (
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/global/helper"
	"dragon-fruit/app/global/structs"
	"encoding/json"
	"sync"
)

// Business GameList 共用 Business
type Business struct {
	UserID string `json:"user_id"` // User 用戶帳號
	GameID string `json:"game_id"` // 遊戲 ID
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
func (b *Business) EnterGame(token string, message []byte) (res []byte, apiErr errorcode.Error) {
	// 解析 token

	// 判斷預執行的行為
	ws := structs.WsAction{}
	if err := json.Unmarshal(message, &ws); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1002001, err.Error(), string(message))
	}

	switch ws.Action {
	// 下注遊戲
	case "bet":
		b.BetGame(message)
	default:

	}

	return
}
