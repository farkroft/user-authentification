package model

import (
	"time"

	uuid "github.com/farkroft/go.uuid"
	"github.com/jinzhu/gorm"
)

// Base model
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// BeforeCreate set ID column with uuid
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// User model
type User struct {
	Base
	Username string `json:"username" gorm:"type:varchar(50);unique"`
	Password string `json:"password"`
}
