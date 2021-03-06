package server

import (
	"fmt"
	"mos/src/glo"
	annonutil "mos/src/server/route/annotations"
	managerutil "mos/src/server/route/apimanager"
	jenkinsutil "mos/src/server/route/jenkins"
	projectutil "mos/src/server/route/project"
	sysutil "mos/src/server/route/sys"
	ticketutil "mos/src/server/route/ticket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Serve Run
func Serve() {
	engine := gin.Default()
	// 允许使用跨域请求  全局中间件
	engine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	engine.POST("/apiv1/addAnnotation", annonutil.APIAnnotation)
	engine.POST("/InitTable", sysutil.InitTable)
	engine.POST("/UserLogin", sysutil.UserLogin)
	engine.POST("/JenkinsPost", jenkinsutil.JenkinsPost)

	sysutil.InitRoute(engine)
	ticketutil.InitRoute(engine)
	projectutil.InitRoute(engine)
	jenkinsutil.InitRoute(engine)
	annonutil.InitRoute(engine)
	managerutil.InitRoute(engine)
	// apiV1 := engine.Group("/api/v1")
	// 用户组路由
	// apiV1.POST("/AddGroup", AddGroup)
	// apiV1.GET("/ListGroup", ListGroup)
	// sysutil.InitRoute(engine)
	// iaas.InitRoute(engine)

	portVal := fmt.Sprintf("0.0.0.0:%d", glo.Config.MosAPI.ServerPort)
	engine.Run(portVal)
}
