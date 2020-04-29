package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/auth-service/external/database"
	"gitlab.com/auth-service/internal/model"
)

// UserRepository interface
type UserRepository interface {
	RegisterUser(model *model.User) error
}

// UserRepo struct
type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepository func
func NewUserRepository(db *database.Database) *UserRepo {
	return &UserRepo{DB: db.DB}
}

// RegisterUser record
func (u *UserRepo) RegisterUser(m *model.User) error {
	db := u.DB.Create(m)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
