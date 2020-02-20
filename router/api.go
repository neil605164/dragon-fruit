package router

import (
	"dragon-fruit/app/handler/test"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		// 載入測試用API
		if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" {

			// Swagger
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		api.GET("/redis-pub", test.RedisPub)
		api.GET("/redis-sub", test.RedisSub)
	}
}
