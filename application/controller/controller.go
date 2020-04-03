package controller

import (
	"github.com/gin-gonic/gin"
)

// API interface
type API interface {
	Ping(*gin.Context)
}
