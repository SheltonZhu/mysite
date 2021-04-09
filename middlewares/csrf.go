package middlewares

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"os"
)

func Csrf() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_TOKEN_SECRET"),
		//IgnoreMethods: []string{"HEAD", "OPTIONS"},
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
		//default: X-CSRF-TOKEN
		TokenGetter: func(c *gin.Context) string {
			return c.Request.Header.Get("X-CSRF-TOKEN")
		},
	})
}
