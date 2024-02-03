package errors

import (
	"fmt"
)

type HttpError struct {
	Code    int    // Código de status HTTP
	Message string // Mensagem de erro
}

func NewHTTPError(code int, message string) HTTPError {
	return &HttpError{
		Code:    code,
		Message: message,
	}
}

// StatusCode retorna o código de status HTTP.
func (e *HttpError) StatusCode() int {
	return e.Code
}

// ErrorMessage retorna a mensagem de erro.
func (e *HttpError) ErrorMessage() string {
	return e.Message
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP error: %d - %s", e.Code, e.Message)
}
