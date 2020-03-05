package xocdia

import (
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/global/helper"
	"dragon-fruit/app/global/structs"
	"encoding/json"
)

// Xocdia 色疊遊戲玩法
type Xocdia struct{}

// EnterGame 進入色碟遊戲
func (x *Xocdia) EnterGame(ws *structs.WsClient, message []byte) (apiErr errorcode.Error) {
	// 解析 token

	// 判斷預執行的行為
	action := structs.WsAction{}
	if err := json.Unmarshal(message, &action); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1002001, err.Error(), string(message))
	}

	switch action.Action {
	case "init":
		break
	// 下注遊戲
	case "bet":
		x.BetGame(ws, message)
		break
	case "deal":
		break
	case "wait":
		break
	case "clear":
		break
	default:
		// todo 寫log + 移除 ws connect
	}

	return
}
