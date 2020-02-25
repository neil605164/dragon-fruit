package game

import (
	"dragon-fruit/app/business/game"
	"dragon-fruit/app/global"
	"dragon-fruit/app/global/helper"
	"dragon-fruit/app/global/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Betting 色蝶開獎結果
// @Summary 色蝶開獎結果
// @Tags Betting
// @Produce json
// @Param body  body []swagger.xocdiaBettingBody true "下注清單"
// @Success 200 {object} swagger.xocdiaBetting "語系清單"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /api/game/xocdia/betting [POST]
func Betting(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	// 取傳入值
	raw := []structs.Xocdia{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		apiErr := helper.ErrorHandle(global.WarnLog, 1002001, err.Error())
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	for _, v := range raw {
		// 驗證參數規則
		if err := helper.ValidateStruct(v); err != nil {
			apiErr := helper.ErrorHandle(global.WarnLog, 1002002, err.Error(), raw)
			c.JSON(http.StatusOK, helper.Fail(apiErr))
			return
		}
	}

	bus := game.Instance()
	bet := bus.Betting(raw)

	c.JSON(http.StatusOK, helper.Success(bet))
}
