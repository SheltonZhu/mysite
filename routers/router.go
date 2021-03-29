package routers

import (
	"github.com/gin-gonic/gin"
	"mysite/controllers"
	"mysite/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(middlewares.Stats())
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	//router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	//router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	v1 := router.Group("v1")

	{
		v1.GET("/ping", controllers.Ping)
		v1.GET("/bibles", controllers.ListBibles)
		v1.GET("/bible/:id", controllers.GetBible)
		v1.POST("/bible", controllers.Ping)
	}
	return router

}
