package main

import (
	"gitlab.com/auth-service/external/server"

	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
)

func main() {
	v := config.NewConfig(constants.EnvConfigFile)
	server.NewServer(v)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run(v.GetString(constants.EnvPort))
}
