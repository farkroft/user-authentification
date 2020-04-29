package presenter

import (
	"fmt"

	"gitlab.com/auth-service/application/response"
)

// ErrorPresenter presenter of error response
func ErrorPresenter(str string, err error) response.ErrorResponse {
	strErr := fmt.Sprintf(str+": %s", err.Error())
	return response.ErrorResponse{Message: strErr}
}
