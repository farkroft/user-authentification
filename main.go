package main

import (
	"gitlab.com/auth-service/external/database"
	"gitlab.com/auth-service/external/server"

	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
	"gitlab.com/auth-service/external/log"
)

func main() {
	log.NewLogger()
	v := config.NewConfig(constants.EnvConfigFile)
	server.NewServer(v)
	db := database.NewDatabase(v)

	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()
}
