package errtmpl

import(
	"net/http"
	"strings"
)

const (
	BadRequst = http.StatusBadRequest
	Unauthorized = http.StatusUnauthorized
	Forbidden = http.StatusForbidden
	NotFound = http.StatusNotFound
	MethodNotAllowed = http.StatusMethodNotAllowed
)

type (
	HttpError struct{
		HttpStatus 					int              `json:"status"`
		Code 						int              `json:"error_code"`
		Message					string			 	 `json:"message"`
	}
)

func HttpError(e errorString, status int) Error{
	return HttpError{
   		HttpStatus: status,
   		Code: err.code,
   		Message: err.message,
   	}
}