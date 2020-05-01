package response

// ErrorResponse struct
type ErrorResponse struct {
	Message string
}

// SuccessResponse struct
type SuccessResponse struct {
	Success bool
	Message string
	Data    interface{}
}

// LoginResponse struct
type LoginResponse struct {
	Token string
}
