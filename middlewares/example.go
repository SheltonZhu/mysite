package middlewares

import (
	"github.com/SheltonZhu/mysite/common"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Stats() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		//c.Set("example", "12345")

		ip := common.GetRequestIP(c)
		log.Printf("ip: %s", ip)

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("cost: %s", latency)
		// access the status we are sending
		//status := c.Writer.Status()
		//log.Printf("status: %d", status)
	}
}
