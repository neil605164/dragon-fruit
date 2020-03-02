package xocdia

import (
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/service"
	"fmt"
)

// BetGame 下注遊戲
func (x *Xocdia) BetGame(message []byte) (apiErr errorcode.Error) {
	// 取帶入參數

	// 取玩家餘額
	ser := service.WaIns()
	balance, apiErr := ser.GetBalance(x.UserID)
	if apiErr != nil {
		return
	}

	fmt.Println(balance)
	// 檢查是否有足夠金額扣款

	// 扣款

	// 寫下注紀錄

	return
}
