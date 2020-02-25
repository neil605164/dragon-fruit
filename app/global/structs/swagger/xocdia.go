package swagger

// 色碟回傳結果
type xocdiaBetting struct {
	ErrorCode int                `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	LogID     string             `json:"log_id"`
	Result    []xocdiaBettingRes `json:"result"`
}

type xocdiaBettingRes struct {
	// 開獎結果(紅白顆數)
	Draw xocdiaBettingResDraw `json:"draw"`
	// 盤面組合(六種組合)
	BetRes []int `json:"bet_res" example:"2,4"`
	// 下注結果
	AllBet []xocdiaBettingResAllBet `json:"bet"`
}

type xocdiaBettingResDraw struct {
	Red   int `json:"red" example:"3"`
	White int `json:"white" example:"1"`
}

type xocdiaBettingResAllBet struct {
	// 用戶ID
	UID string `json:"uid" example:"abc123"`
	// 用戶總贏分
	TotalWin float64 `json:"total_win" example:"1052"`
	// 用戶下注各區塊資訊
	Bet []xocdiaBettingResAllBetBet `json:"bet"`
}

type xocdiaBettingResAllBetBet struct {
	// 下注區塊
	Project int `json:"project" example:"1"`
	// 下注金額
	Mount float64 `json:"mount" example:"100"`
	// 該區塊贏分
	Win float64 `json:"win" example:"200"`
}

// 色碟下注列表
type xocdiaBettingBody struct {
	UID string                 `json:"uid" example:"1"`
	Bet []xocdiaBettingBodyBet `json:"bet"`
}
type xocdiaBettingBodyBet struct {
	Project int     `json:"project" example:"1"`
	Mount   float64 `json:"mount" example:"100"`
}
