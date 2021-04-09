package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func Session() gin.HandlerFunc {
	//"github.com/gin-contrib/sessions/redis"
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	return sessions.Sessions(os.Getenv("SESSION_NAME"), store)
}
