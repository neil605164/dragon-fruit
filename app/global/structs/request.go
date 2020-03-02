package structs

// WsAction websocket 預執行的行為
type WsAction struct {
	Action string `json:"action"`
}

// BetGame 遊戲下注
type BetGame struct {
	BetAmount float64 `json:"bet_amount"`
	BetResult string  `json:"bet_result"`
}
