package business

import (
	"dragon-fruit/app/business/xocdia"
	"dragon-fruit/app/global/errorcode"
)

// Game 各款遊戲必須共有的 func
type Game interface {
	EnterGame(token string, message []byte) (res []byte, apiErr errorcode.Error)
}

// NewGame 各遊戲接口
func NewGame(gameID string) (str Game) {
	switch gameID {
	case "1":
		str = &xocdia.Xocdia{}
	}

	return
}
