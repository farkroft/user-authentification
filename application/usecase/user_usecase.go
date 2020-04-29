package usecase

import (
	"fmt"
	"net/http"

	"gitlab.com/auth-service/internal/model"

	"gitlab.com/auth-service/application/request"
)

// UserRegister usecase for register user
func (u *UseCase) UserRegister(req request.UserRequest) (int, string, error) {
	if req.Password == "" {
		err := fmt.Errorf("password is empty")
		return http.StatusBadRequest, "bad request", err
	}

	if req.Username == "" {
		err := fmt.Errorf("username is empty")
		return http.StatusBadRequest, "bad request", err
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := u.UserRepo.RegisterUser(&user)
	if err != nil {
		return http.StatusInternalServerError, "database", err
	}

	return http.StatusOK, "OK", nil
}
