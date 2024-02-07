package middleware

import (
	"net/http"

	"github.com/d90ares/iBeers/config/logs"
	"github.com/d90ares/iBeers/ibeers/errors"
	"go.uber.org/zap"
)

// MiddlewareHandler é o middleware personalizado para tratamento de erros
type MiddlewareHandler struct{}

// ServeHTTP implementa a interface http.Handler
func (m *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logs.Info("Request Started", zap.String("method", r.Method), zap.String("path", r.URL.Path))

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
