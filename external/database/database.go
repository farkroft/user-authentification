package database

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
	"gitlab.com/auth-service/external/config"
	"gitlab.com/auth-service/external/constants"
)

// Repository repository
type Repository interface {
	NewDatabase(*config.Config) *Database
}

// Database struct
type Database struct {
	db *gorm.DB
}

// Base model
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (base *Base) beforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Printf("error when generate uuid %s", err.Error())
		return err
	}
	return scope.SetColumn("ID", uuid)
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
