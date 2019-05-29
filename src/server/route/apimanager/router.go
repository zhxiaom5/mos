package apimanager

import (
	"log"
	"mos/src/pkg/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// 视图映射
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

// InitRoute 初始化路由
func InitRoute(engine *gin.Engine) {
	if engine == nil {
		log.Panicln(`engine is nil!`)
	}
	engine.Use(jwt.JWT())
	// 用户路由
	managerAPI := engine.Group("/manager")
	{
		managerAPI.GET("/APIList", APIList)
		managerAPI.POST("/APIAdd", APIAdd)
		managerAPI.POST("/APIDelete", APIDelete)
		managerAPI.POST("/APIKeyAdd", APIKeyAdd)
		managerAPI.POST("/APIKeyDelete", APIKeyDelete)

	}
}
