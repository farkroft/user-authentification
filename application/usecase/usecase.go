package usecase

import (
	"gitlab.com/auth-service/application/repository"
)

// UseCase given ...
type UseCase struct {
	UserRepo repository.UserRepository
}

// NewUsecase return data source instance
func NewUsecase(repo repository.UserRepository) *UseCase {
	return &UseCase{UserRepo: repo}
}
