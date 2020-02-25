package game

import "sync"

// 共用 Business
type Business struct {
}

var singleton *Business
var once sync.Once

// Instance 獲得單例對象
func Instance() *Business {
	once.Do(func() {
		singleton = &Business{}
	})
	return singleton
}
