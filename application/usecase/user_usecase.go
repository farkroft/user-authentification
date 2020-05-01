package usecase

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/auth-service/application/response"

	"gitlab.com/auth-service/external/constants"

	"gitlab.com/auth-service/external/util"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/auth-service/application/request"
	"gitlab.com/auth-service/external/log"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("hash", err)
		return http.StatusInternalServerError, "hash", err
	}

	req.Password = string(hashedPassword)

	_, err = u.UserRepo.RegisterUser(req)
	if err != nil {
		return http.StatusInternalServerError, "database", err
	}

	return http.StatusOK, "OK", nil
}

// UserLogin usecase for user login
func (u *UseCase) UserLogin(req request.UserRequest) (int, string, interface{}, error) {
	if req.Password == "" {
		err := fmt.Errorf("password is empty")
		return http.StatusBadRequest, "bad request", nil, err
	}

	if req.Username == "" {
		err := fmt.Errorf("username is empty")
		return http.StatusBadRequest, "bad request", nil, err
	}

	user, err := u.UserRepo.GetUser(req)
	if err != nil && !util.IsErrorRecordNotFound(err) {
		if util.IsErrorRecordNotFound(err) {
			return http.StatusNotFound, "user not found", nil, err
		}

		log.Errorf("get user", err)
		return http.StatusInternalServerError, "get user", nil, err
	}
	now := util.WIBTimezone(util.Now())
	expiredAt := now.Add(time.Minute * constants.TwelveHoursInMinute).Unix()
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Errorf("hash pass", err)
		return http.StatusBadRequest, "Bad Credentials", nil, err
	}

	userClaims := jwt.MapClaims{
		"UserID":   user.ID,
		"Username": user.Username,
		"StandardClaims": &jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	fmt.Println()
	tokenStr, err := token.SignedString([]byte(u.Cfg.GetString(constants.EnvJWTSecret)))
	if err != nil {
		log.Errorf("token decode", err)
		return http.StatusInternalServerError, "token decode", nil, err
	}

	resp := response.LoginResponse{
		Token: tokenStr,
	}

	return http.StatusOK, "OK", resp, nil
}

// UserAuthVerify to verify token valid or not
func (u *UseCase) UserAuthVerify(str string) (*jwt.Token, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.Cfg.GetString(constants.EnvJWTSecret)), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
