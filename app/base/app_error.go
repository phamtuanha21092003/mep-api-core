package base

import "net/http"

// AppError is build in error
type AppError struct {
	Err     error
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func BadRequest(err error) error {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "bad_request",
		Err:     err,
	}
}

func InternalServerError(err error) error {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "internal_server_error",
		Err:     err,
	}
}

func Unauthorized(err error) error {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
		Err:     err,
	}
}

func Forbidden(err error) error {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: "forbidden",
		Err:     err,
	}
}

func NotFound(err error) error {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "not_found",
		Err:     err,
	}
}

func Conflict(err error) error {
	return &AppError{
		Code:    http.StatusConflict,
		Message: "Conflict",
		Err:     err,
	}
}

func GatewayTimeout(err error) error {
	return &AppError{
		Code:    http.StatusGatewayTimeout,
		Message: "gateway_timeout",
		Err:     err,
	}
}
