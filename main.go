package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
)

func main() {
	v := config.NewConfig(constants.EnvConfigFile)
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(v.GetString(constants.EnvPort))
}
