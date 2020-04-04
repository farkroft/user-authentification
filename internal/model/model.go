package model

import "gitlab.com/auth-service/external/database"

// User model
type User struct {
	database.Base
	Username string `json:"username" gorm:"type:varchar(50);unique"`
	Password string `json:"password"`
}
