package controllers

import (
	"github.com/SheltonZhu/mysite/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, services.Heartbeat())
}
