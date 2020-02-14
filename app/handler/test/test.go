package test

import (
	"dragon-fruit/app/business"
	"dragon-fruit/app/global/helper"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func ServeWs(c *gin.Context) {
	hub := newHub()
	go hub.run()


	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client


	go client.writePump()
	go client.readPump()
	//升级get请求为webSocket协议
	//ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	return
	//}
	//defer ws.Close()
	//for {
	//	//读取ws中的数据
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//	if string(message) == "ping" {
	//		message = []byte("pong")
	//	}
	//	//写入ws数据
	//	err = ws.WriteMessage(mt, message)
	//	if err != nil {
	//		break
	//	}
	//}
}

// SetRedisValue 測試 Redis 存值
// @Summary 測試 Redis 存值
// @description Redis Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /api/test/set_redis [POST]
func SetRedisValue(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	redisBus := business.RedisIns()
	if apiErr := redisBus.SetRedisKey(); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// GetRedisValue 測試 Redis 取值
// @Summary 測試 Redis 取值
// @description Redis Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /api/test/get_redis [GET]
func GetRedisValue(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	redisBus := business.RedisIns()
	value, err := redisBus.GetRedisValue()
	if err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success(value))

}

// PingDBOnce Ping DB 測試
// @Summary Ping DB 測試
// @description DB Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /api/test/ping_db_once [GET]
func PingDBOnce(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	dbBus := business.DBIns()
	if err := dbBus.PingDBOnce(); err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success("123"))
}

// PingDBSecond Ping DB 測試
// @Summary Ping DB 測試
// @description DB Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /api/test/ping_db_second [GET]
func PingDBSecond(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	dbBus := business.DBIns()
	if err := dbBus.PingDBSecond(); err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success("456"))
}

// ErrorTest 測試錯誤發生時是否可以回傳正確的 logID
// @Summary 測試錯誤發生時是否可以回傳正確的 logID
// @description DB Pool 測試錯誤發生時是否可以回傳正確的 logID
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /api/test/error_task [GET]
func ErrorTest(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	errBus := business.ErrIns()
	if err := errBus.GetErrorLog(); err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success("999"))
}
