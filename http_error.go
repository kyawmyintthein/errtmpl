package goerror


type HttpError struct {
        HttpStatus      int		 `json:"status"`
        ErrorMessage    interface{}      `json:"error_message"`
}

func newHttpError(status int, message interface{}) HttpError {
    return HttpError{HttpStatus: status, ErrorMessage: message}
}
