package main

import (
	"gitlab.com/auth-service/application/controller"
	"gitlab.com/auth-service/application/repository"
	"gitlab.com/auth-service/application/usecase"
	"gitlab.com/auth-service/external/database"
	"gitlab.com/auth-service/external/server"

	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
	"gitlab.com/auth-service/external/log"
)

func main() {
	log.NewLogger()
	v := config.NewConfig(constants.EnvConfigFile)
	db, err := database.NewDatabase(v)
	if err != nil {
		panic(err)
	}
	err = db.Migrate()
	if err != nil {
		log.Errorf("migrate", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewUsecase(userRepo, v)
	ctl := controller.NewController(usecase)
	server.NewServer(v, ctl)
}
