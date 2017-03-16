package goerror


type HttpError struct {
        HttpStatus      int         `json:"status"`
        ErrorMessage    string      `json:"error_message"`
}

func newHttpError(status int, message string) HttpError {
    return HttpError{HttpStatus: status, ErrorMessage: message}
}