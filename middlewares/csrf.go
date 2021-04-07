package middlewares

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Csrf() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret:        "cool_csrf_secret",
		//IgnoreMethods: []string{"HEAD", "OPTIONS"},
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
		//default: X-CSRF-TOKEN
		//TokenGetter: func(c *gin.Context) string {
		//	return c.Request.Header.Get("csrftoken")
		//},
	})
}
