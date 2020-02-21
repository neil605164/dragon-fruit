package business

import (
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/repository"
	"fmt"
	"log"
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

// RedisPub redis pub
func (a *RedisBus) RedisPub() (apiErr errorcode.Error) {
	repo := repository.RedisIns()

	_, err := repo.Publish("test/foo", "barsss")
	if err != nil {
		log.Fatal(err)
	}

	return
}

// RedisSub redis sub
func (a *RedisBus) RedisSub() {
	repo := repository.RedisIns()

	channel := fmt.Sprintf("test/foo")

	reply := make(chan []byte)
	repo.Subscribe(channel, reply)

	fmt.Println("=============================================###")

	msg := <-reply

	fmt.Println("Msgggggg", string(msg))
	fmt.Printf("++++++++++++++recieved %q", string(msg))

}
