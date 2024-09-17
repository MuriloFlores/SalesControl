package erro_handler

import "net/http"

type ErrorHandler struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
	Causes  []Cause
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ErrorHandler) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
