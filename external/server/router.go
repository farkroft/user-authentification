package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/auth-service/application/controller"
)

// NewRouter router
func NewRouter(r *gin.Engine) {
	r.GET("/ping", controller.Ping)
}
