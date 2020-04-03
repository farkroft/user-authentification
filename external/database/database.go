package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
)

// Database struct
type Database struct {
	db *gorm.DB
}

// NewDatabase init new database
func NewDatabase(cfg *config.Config) *Database {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.GetString(constants.EnvDBHost),
		cfg.GetString(constants.EnvDBPort),
		cfg.GetString(constants.EnvDBUser),
		cfg.GetString(constants.EnvDBPass),
		cfg.GetString(constants.EnvDBName))
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	return &Database{
		db: db,
	}
}
