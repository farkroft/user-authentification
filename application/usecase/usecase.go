package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/auth-service/application/repository"
	"gitlab.com/auth-service/application/request"
	"gitlab.com/auth-service/external/config"
)

// CaseRepo interface
type CaseRepo interface {
	UserRegister(req request.UserRequest) (int, string, error)
	UserLogin(req request.UserRequest) (int, string, interface{}, error)
	UserAuthVerify(str string) (*jwt.Token, error)
}

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
