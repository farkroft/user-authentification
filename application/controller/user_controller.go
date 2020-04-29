package controller

import (
	"net/http"

	"gitlab.com/auth-service/application/presenter"

	"github.com/gin-gonic/gin"
	"gitlab.com/auth-service/application/request"
	"gitlab.com/auth-service/external/log"
)

// Register new account
func (ctl *Controller) Register(c *gin.Context) {
	req := request.UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind json", err)
		c.JSON(http.StatusInternalServerError, presenter.ErrorPresenter("bind json", err))
		return
	}

	httpCode, strErr, err := ctl.UserUseCase.UserRegister(req)
	if err != nil {
		log.Errorf(strErr, err)
		c.JSON(httpCode, presenter.ErrorPresenter(strErr, err))
		return
	}
}
