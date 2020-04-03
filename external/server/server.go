package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
)

// NewServer instance of server
func NewServer(cfg *config.Config) {
	r := gin.New()
	NewRouter(r)
	r.Run(cfg.GetString(constants.EnvPort))
}
