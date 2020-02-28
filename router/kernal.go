package router

import (
	"dragon-fruit/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	r.Use(middleware.WriteLog)

	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	// api route
	LoadBackendRouter(r)
}
