package game

import (
	"crypto/rand"
	"dragon-fruit/app/global/structs"
	"math/big"
)

func (*Business) Betting(req []structs.Xocdia) (resp structs.XocdiaResp) {
	// 開獎結果
	result := map[string]int{
		"red":   0,
		"white": 0,
	}
	// 生成 4 個 [0, 1] 範圍的真隨機數。
	for i := 0; i < 4; i++ {
		dia, _ := rand.Int(rand.Reader, big.NewInt(2))
		// 將bigInt轉字串
		diaConv := dia.String()
		switch diaConv {
		case "0":
			result["white"]++
		case "1":
			result["red"]++
		}
	}

	// 盤面組合(預設為沒中獎)
	betRes := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
	}
	// 雙 (代號1)
	if result["white"] == 4 || result["red"] == 4 || result["white"] == 2 || result["red"] == 2 {
		betRes[1] = true
	}
	// 單 (代號2)
	if result["white"] == 3 || result["red"] == 3 {
		betRes[2] = true
	}
	// 紅4 (代號3)
	if result["red"] == 4 {
		betRes[3] = true
	}
	// 紅3 (代號4)
	if result["red"] == 3 {
		betRes[4] = true
	}
	// 白4 (代號5)
	if result["white"] == 4 {
		betRes[5] = true
	}
	// 白3 (代號6)
	if result["white"] == 3 {
		betRes[6] = true
	}

	// 賠率假資料
	rate := map[int]float64{
		1: 0.98,
		2: 0.98,
		3: 14.6,
		4: 2.92,
		5: 2.92,
		6: 14.6,
	}

	// 組回傳資訊
	resp.Draw = result
	resp.BetRes = betRes

	// 計算回傳贏分資訊
	for _, v := range req {
		tmpRes := structs.XocdiaRespRes{}
		tmpRes.UID = v.UID
		// 預設總贏分 0
		tmpRes.TotalWin = 0
		// 計算全部下注各贏多少
		for _, bv := range v.Bet {
			tmpBet := structs.XocdiaRespResBet{}
			tmpBet.Project = bv.Project
			tmpBet.Mount = bv.Mount
			// 是否有押中
			if betRes[bv.Project] {
				tmpBet.Win = bv.Mount * rate[bv.Project]
				tmpRes.TotalWin = tmpRes.TotalWin + tmpBet.Win
			}
			tmpRes.Bet = append(tmpRes.Bet, tmpBet)
		}
		resp.AllBet = append(resp.AllBet, tmpRes)
	}
	return
}
