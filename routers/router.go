package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"mysite/controllers"
	"mysite/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Stats())
	router.Use(middlewares.Cors())
	router.Use(middlewares.Session())
	router.Use(middlewares.Csrf())

	test := router.Group("test")
	//test.Use(gin.BasicAuth(gin.Accounts{
	//	"admin": "123456",
	//}))

	{
		test.GET("/protected", func(c *gin.Context) {
			c.String(200, csrf.GetToken(c))
		})
		test.POST("/protected", func(c *gin.Context) {
			c.String(200, "CSRF token is valid")
		})
		test.GET("/hello", func(c *gin.Context) {
			session := sessions.Default(c)
			if session.Get("hello") != "world" {
				session.Set("hello", "world")
				session.Save()
			}
			c.JSON(200, gin.H{"hello": session.Get("hello")})
		})
	}

	v1 := router.Group("v1")

	{
		v1.GET("/ping", controllers.Ping)
		v1.GET("/bibles", controllers.ListBibles)
		v1.GET("/bible/:id", controllers.GetBible)
		v1.POST("/bible/:id", controllers.Ping)
		v1.PUT("/bible", controllers.Ping)
		v1.PATCH("/bible", controllers.Ping)
		v1.DELETE("/bible/:id", controllers.Ping)
	}
	return router

}
