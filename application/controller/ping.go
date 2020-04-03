package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping controller
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	return
}
