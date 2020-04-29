package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/auth-service/application/controller"
	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
)

// NewServer instance of server
func NewServer(cfg *config.Config, ctl *controller.Controller) {
	r := gin.New()
	NewRouter(r, ctl)
	r.Run(cfg.GetString(constants.EnvPort))
}
