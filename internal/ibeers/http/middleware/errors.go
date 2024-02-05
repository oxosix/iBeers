package middleware

import (
	"net/http"

	"github.com/d90ares/iBeers/internal/ibeers/errors"
)

// MiddlewareHandler é o middleware personalizado para tratamento de erros
type MiddlewareHandler struct{}

// ServeHTTP implementa a interface http.Handler
func (m *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Tente executar a próxima função
	next(w, r)

	// Verifique se ocorreu um erro e se é do tipo GenericError
	if err, ok := recover().(errors.HttpError); ok {
		// Manipule o erro aqui
		w.WriteHeader(err.Code)
		w.Write([]byte(err.Message))
	}
}

// NewMiddleware cria uma instância do MiddlewareHandler
func NewMiddleware() *MiddlewareHandler {
	return &MiddlewareHandler{}
}
