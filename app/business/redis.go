package business

import (
	"dragon-fruit/app/global"
	"dragon-fruit/app/repository"
	"fmt"
	"sync"
)

// RedisBus 管理者Business
type RedisBus struct {
}

var redisSingleton *RedisBus
var redisOnce sync.Once

// RedisIns 獲得單例對象
func RedisIns() *RedisBus {
	redisOnce.Do(func() {
		redisSingleton = &RedisBus{}
	})
	return redisSingleton
}

// RedisSub redis sub
func (a *RedisBus) RedisSub() {
	repo := repository.RedisIns()

	reply := make(chan []byte)
	go repo.Subscribe(global.Channel, reply)

	for {
		msg := <-reply

		// todo 觸發廣播
		fmt.Println("====>", msg)
	}
}
