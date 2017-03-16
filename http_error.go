package errtmpl

import (
	"net/http"
)

const (
	BadRequst        = http.StatusBadRequest
	Unauthorized     = http.StatusUnauthorized
	Forbidden        = http.StatusForbidden
	NotFound         = http.StatusNotFound
	MethodNotAllowed = http.StatusMethodNotAllowed
)

type (
	HttpError struct {
		HttpStatus int    `json:"status"`
		Code       string `json:"error_code"`
		Message    string `json:"message"`
	}
)

func ErrorWithHttpStatus(e errorString, status int) HttpError {
	return HttpError{
		HttpStatus: status,
		Code:       e.name,
		Message:    e.message,
	}
}

func DefaultErrorWithHttpStatus(e error, name string, status int) HttpError {
	return HttpError{
		HttpStatus: status,
		Code:       name,
		Message:    e.Error(),
	}
}
