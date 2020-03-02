package service

import (
	"dragon-fruit/app/global/errorcode"
	"sync"
)

// WalletSer Wallet API 專用
type WalletSer struct {
}

var waSingleton *WalletSer
var waOnce sync.Once

// WaIns 獲得Rotate對象
func WaIns() *WalletSer {
	waOnce.Do(func() {
		waSingleton = &WalletSer{}
	})
	return waSingleton
}

// GetBalance 取餘額
func (w *WalletSer) GetBalance(token string) (balance int, apiErr errorcode.Error) {

	balance = 100
	return
}
