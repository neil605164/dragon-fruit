package structs

// 色蝶下注清單
type XocdiaResp struct {
	// 開獎結果(紅白顆數)
	Draw map[string]int `json:"draw"`
	// 盤面組合(六種組合)
	BetRes []int `json:"bet_res"`
	// 下注結果
	AllBet []XocdiaRespRes `json:"all_bet"`
}

type XocdiaRespRes struct {
	// 用戶ID
	UID string `json:"uid"`
	// 用戶總贏分
	TotalWin float64 `json:"total_win"`
	// 用戶下注各區塊資訊
	Bet []XocdiaRespResBet `json:"bet"`
}

type XocdiaRespResBet struct {
	// 下注區塊
	Project int `json:"project"`
	// 下注金額
	Mount float64 `json:"mount"`
	// 該區塊贏分
	Win float64 `json:"win"`
}
