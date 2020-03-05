package xocdia

import (
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/global/helper"
	"dragon-fruit/app/global/structs"
	"dragon-fruit/app/service"
	"encoding/json"
)

// BetGame 下注遊戲
func (x *Xocdia) BetGame(ws *structs.WsClient, message []byte) (apiErr errorcode.Error) {
	// 取帶入參數
	bets := []structs.BetGame{}
	if err := json.Unmarshal(message, &bets); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1002002, err.Error(), string(message))
	}

	// 取玩家餘額(呼叫錢包) todo
	ser := service.WaIns()
	balance, apiErr := ser.GetBalance(ws.UID)
	if apiErr != nil {
		return
	}

	// 檢查是否有足夠金額扣款
	total := 0.0
	for _, v := range bets {
		total = total + v.Mount
	}

	if (balance - total) < 0.0 {
		apiErr = helper.ErrorHandle(global.WarnLog, 1002003, "", string(message))
	}

	// 扣款(呼叫錢包) todo
	if apiErr = ser.Deduction(total, ws.UID); apiErr != nil {
		return
	}

	// 寫扣款紀錄 todo(userid,parentid,amount,rounid,gameid,bet_time) mongo DB

	// 寫下注紀錄(先取+檢查，在覆蓋)

	return
}
