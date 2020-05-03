package mock

import (
	"gitlab.com/auth-service/application/request"
	"gitlab.com/auth-service/internal/model"
)

type MockRepository struct{}

func (u *MockRepository) RegisterUser(req request.UserRequest) (model.User, error) {
	return model.User{}, nil
}

func (u *MockRepository) GetUser(req request.UserRequest) (model.User, error) {
	return model.User{}, nil
}
