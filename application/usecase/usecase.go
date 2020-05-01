package usecase

import (
	"gitlab.com/auth-service/application/repository"
	"gitlab.com/auth-service/external/config"
)

// UseCase given ...
type UseCase struct {
	UserRepo repository.UserRepository
	Cfg      config.Repository
}

// NewUsecase return data source instance
func NewUsecase(repo repository.UserRepository, cfg config.Repository) *UseCase {
	return &UseCase{
		UserRepo: repo,
		Cfg:      cfg,
	}
}
