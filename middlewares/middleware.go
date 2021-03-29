package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"mysite/common"
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
		// access the status we are sending
		status := c.Writer.Status()
		log.Printf("status: %d, cost: %s", status, latency)
	}
}
