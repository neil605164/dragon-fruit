package repository

import (
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/errorcode"
	"dragon-fruit/app/global/helper"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Redis 存取值
type Redis struct{}

var redisSingleton *Redis
var redisOnce sync.Once

// redisPool 存放redis連線池的全域變數
var redisPool *redis.Pool

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     100,              // int 最大可允許的閒置連線數
		MaxActive:   10000,            // int 最大建立的連線數，默認為0不限制(reids 預設最大連線量)
		IdleTimeout: 20 * time.Second, // 連線過期時間，默認為0表示不做過期限制
		Wait:        true,             // 當連線超出限制數量後，是否等待到空閒連線釋放
		Dial: func() (c redis.Conn, err error) {
			// 使用redis封裝的Dial進行tcp連接
			host := global.Config.Redis.RedisHost
			port := global.Config.Redis.RedisPort
			pwd := global.Config.Redis.RedisPwd

			// 組合連接資訊
			var connectionString = fmt.Sprintf("%s:%s", host, port)
			c, err = redis.Dial(
				"tcp",
				connectionString,
				redis.DialPassword(pwd),
			)

			if err != nil {
				helper.ErrorHandle(global.WarnLog, 1001012, err.Error())
				return
			}
			return
		}, // 連接redis的函数
		TestOnBorrow: func(redis redis.Conn, t time.Time) (err error) {
			// 每5秒ping一次redis
			if time.Since(t) < (5 * time.Second) {
				return
			}
			_, err = redis.Do("PING")
			if err != nil {
				helper.ErrorHandle(global.WarnLog, 1001019, err.Error())
				return
			}

			return
		}, // 定期對 redis server 做 ping/pong 測試

	}

}

// RedisIns 獲得單例對象
func RedisIns() *Redis {
	redisOnce.Do(func() {
		redisSingleton = &Redis{}
	})
	return redisSingleton
}

// RedisPing 檢查Redis是否啟動
func RedisPing() {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("🔔🔔🔔 REDIS CONNECT ERROR: %v 🔔🔔🔔", err.Error())
	}
}

// Exists 檢查key是否存在
func (rConn *Redis) Exists(key string) (ok bool, apiErr errorcode.Error) {
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
func (rConn *Redis) Publish(key string, value string) (reply int, apiErr errorcode.Error) {
	var err error

	conn := redisPool.Get()
	defer conn.Close()

	reply, err = redis.Int(conn.Do("PUBLISH", key, value))
	if err != nil {
		helper.ErrorHandle(global.WarnLog, 1003001, err.Error())
		return
	}

	return
}

// Subscribe redis sub
func (rConn *Redis) Subscribe(channel string, msg chan []byte) {

	conn := redisPool.Get()
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	if err := psc.PSubscribe(channel); err != nil {
		fmt.Println("########", err)
	}

	for conn.Err() == nil {
		// fmt.Println("Subscribe Redis Conn ====>", conn)
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("========>%s: message: %s\n", v.Channel, v.Data)
			msg <- v.Data
		case redis.Subscription:
			fmt.Printf("--------->%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println("~~~~~Error", conn.Err())
		}

	}

}
