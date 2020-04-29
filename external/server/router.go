package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/auth-service/application/controller"
)

// NewRouter router
func NewRouter(r *gin.Engine, ctl *controller.Controller) {
	r.GET("/ping", ctl.Ping)
	r.POST("/register", ctl.Register)
}
