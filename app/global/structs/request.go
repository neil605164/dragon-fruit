package structs

// 色蝶下注清單
type Xocdia struct {
	// 用戶ID
	UID string      `json:"uid" validate:"required"`
	// 用戶下注各區塊資訊
	Bet []xocdiaBet `json:"bet" validate:"required"`
}
type xocdiaBet struct {
	// 下注區塊
	Project int     `json:"project" validate:"required"`
	// 下注金額
	Mount   float64 `json:"mount" validate:"required"`
}
