package middleware

import (
	"net/http"
	"time"

	"github.com/d90ares/iBeers/pkg/logs"
	"github.com/d90ares/iBeers/pkg/metrics"
	"go.uber.org/zap"
)

// MiddlewareHandler é o middleware personalizado para tratamento de erros
type MiddlewareHandler struct{}

type statusCapturingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCapturingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// ServeHTTP implementa a interface http.Handler
func (m *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Captura o início da requisição
	startTime := time.Now()

	// Captura a resposta
	capturedResponseWriter := &statusCapturingResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK, // Define um padrão inicial
	}

	logs.Info("Request Started", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	// Executa a próxima função (handler ou outro middleware)
	next(capturedResponseWriter, r)

	// Calcula o tempo total da requisição
	duration := time.Since(startTime).Seconds()
	// durationMs := duration.Milliseconds()

	// Incrementa as métricas
	metrics.IncrementRequestCount(r.Method, r.URL.Path, capturedResponseWriter.statusCode)

	logs.Info("Request Ended", zap.Int("Status", capturedResponseWriter.statusCode),
		zap.Float64("DurationMs", duration),
		zap.String("Method", r.Method),
		zap.String("Path", r.URL.Path))
}

// NewMiddleware cria uma instância do MiddlewareHandler
func NewMiddleware() *MiddlewareHandler {
	return &MiddlewareHandler{}
}
