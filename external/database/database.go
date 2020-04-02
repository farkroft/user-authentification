package database

import (
	"github.com/jinzhu/gorm"
)

// Database struct
type Database struct {
	db *gorm.DB
}

// NewDatabase init new database
func NewDatabase() *Database {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=auth_service_local sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &Database{
		db: db,
	}
}
