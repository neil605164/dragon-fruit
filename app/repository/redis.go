package repository

import (
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/global/helper"
	"dragon-fruit/app/model"
	"fmt"
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

// Redis 存取值
type Redis struct{}

var redisSingleton *Redis
var redisOnce sync.Once

// RedisIns 獲得單例對象
func RedisIns() *Redis {
	redisOnce.Do(func() {
		redisSingleton = &Redis{}
	})
	return redisSingleton
}

// RedisPing 檢查Redis是否啟動
func RedisPing() {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("🔔🔔🔔 REDIS CONNECT ERROR: %v 🔔🔔🔔", err.Error())
	}
}

// Exists 檢查key是否存在
func (rConn *Redis) Exists(key string) (ok bool, apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	chkExisits, _ := conn.Do("EXISTS", key)
	ok, err := redis.Bool(chkExisits, nil)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001018, err.Error())

		return
	}

	return
}

// Set 存入redis值
func (rConn *Redis) Set(key string, value interface{}, expiretime int) (apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", key, value, "EX", expiretime); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001013, err.Error())

		return
	}
	return
}

// Get 取出redis值
func (rConn *Redis) Get(key string) (value string, apiErr errorcode.Error) {

	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		helper.ErrorHandle(global.WarnLog, 1001020, err.Error(), key)
	}

	return
}

// Delete 刪除redis值
func (rConn *Redis) Delete(key string) (apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001015, err.Error())

		return
	}

	return
}

// Append 在相同key新增多個值
func (rConn *Redis) Append(key string, value interface{}) (n interface{}, apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	n, err := conn.Do("APPEND", key, value)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001016, err.Error())

		return
	}

	return
}

// HashSet Hash方式存入redis值
func (rConn *Redis) HashSet(hkey string, key interface{}, value interface{}, time int) (apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	// 存值
	if _, err := conn.Do("hset", hkey, key, value); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001014, err.Error())

		return
	}

	// 設置過期時間
	if _, err := conn.Do("EXPIRE", hkey, time); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, 1001017, err.Error())

		return
	}

	return
}

// HashGet Hash方式取出redis值
func (rConn *Redis) HashGet(hkey string, field interface{}) (value string, apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	conn := redisPool.Get()
	defer conn.Close()

	// 取值
	value, err := redis.String(conn.Do("HGET", hkey, field))
	if err != nil {
		helper.ErrorHandle(global.WarnLog, 1001021, err.Error(), hkey, field)
	}

	return
}

// Publish redis pub
func (rConn *Redis) Publish() (apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	fmt.Println("Publish Redis Pool redisPool Access", redisPool)

	return
}

// Subscribe redis sub
func (rConn *Redis) Subscribe() (apiErr errorcode.Error) {
	redisPool := model.RedisPoolConnect()
	fmt.Println("Subscribe Redis Pool redisPool Access", redisPool)

	return
}
