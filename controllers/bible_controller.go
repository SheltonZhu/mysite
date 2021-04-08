package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mysite/dto"
	"mysite/services"
	"net/http"
	"strconv"
	"strings"
)

func ListBibles(c *gin.Context) {
	data, err := services.ListBibles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBible(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Result{Code: 1, Message: fmt.Sprintf("'%v' is not a number id", idStr)})
		return
	}
	data, err := services.GetBible(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateBible(c *gin.Context) {
	text := c.DefaultPostForm("text", "")
	if len(strings.TrimSpace(text)) == 0 {
		c.JSON(http.StatusBadRequest, dto.Result{Code: 1, Message: fmt.Sprintf("'%v' field is required", "text")})
		return
	}

	data, err := services.CreateBible(text)
	if err != nil {
		c.JSON(http.StatusBadRequest, data)
		return
	}
	c.JSON(http.StatusOK, data)

}

func DeleteBile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Result{Code: 1, Message: fmt.Sprintf("'%v' is not a number id", idStr)})
		return
	}
	data, err := services.DeleteBible(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data)
	}
	c.JSON(http.StatusOK, data)
}
