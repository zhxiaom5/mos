package annotations

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
	annAPI := engine.Group("/iaas")
	{
		annAPI.GET("/AnnotationList", AnnotationList)
		annAPI.POST("/AnnotationAdd", AnnotationAdd)
		annAPI.POST("/AnnotationDelete", AnnotationDelete)
		annAPI.GET("/Tst", Tst)
	}
}
