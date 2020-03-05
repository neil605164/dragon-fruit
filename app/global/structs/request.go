package structs

// WsAction websocket 預執行的行為
type WsAction struct {
	Action string `json:"action"`
}

// BetGame 遊戲下注
type BetGame struct {
	Project uint8   `json:"project"`
	Mount   float64 `json:"mount"`
}
