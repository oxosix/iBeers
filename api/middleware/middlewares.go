package middleware

import (
	"net/http"

	"github.com/d90ares/iBeers/pkg/errors"
	"github.com/d90ares/iBeers/pkg/logs"
	"go.uber.org/zap"
)

// MiddlewareHandler é o middleware personalizado para tratamento de erros
type MiddlewareHandler struct{}

type statusCapturingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// ServeHTTP implementa a interface http.Handler
func (m *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	capturedResponseWriter := &statusCapturingResponseWriter{ResponseWriter: w}
	logs.Info("Request Started", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	// Tente executar a próxima função
	next(capturedResponseWriter, r)

	// Verifique se ocorreu um erro e se é do tipo GenericError
	if err, ok := recover().(errors.HttpError); ok {
		// Manipule o erro aqui
		w.WriteHeader(err.Code)
		w.Write([]byte(err.Message))
	}
	logs.Info("Request Ended", zap.Int("Status", capturedResponseWriter.statusCode))
}

func (w *statusCapturingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// NewMiddleware cria uma instância do MiddlewareHandler
func NewMiddleware() *MiddlewareHandler {
	return &MiddlewareHandler{}
}
