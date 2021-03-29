package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mysite/dto"
	"mysite/services"
	"net/http"
	"strconv"
)

func ListBibles(c *gin.Context) {
	c.JSON(http.StatusOK, services.ListBibles())
}

func GetBible(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Result{Code: 1, Message: fmt.Sprintf("'%v' is not a number id", idStr)})
		return
	}
	c.JSON(http.StatusOK, services.GetBible(id))
}
