package jenkins

import (
	"log"

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
	// engine.Use(jwt.JWT())
	jenkinsAPI := engine.Group("/jenkins")
	{
		jenkinsAPI.POST("/JenkinsPost", JenkinsPost)
		jenkinsAPI.GET("/JenkinsJobList", JenkinsJobList)

	}
}
