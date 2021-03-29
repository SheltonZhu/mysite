package controllers

import (
	"github.com/gin-gonic/gin"
	"mysite/services"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, services.Heartbeat())
}