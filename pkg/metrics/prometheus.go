// metrics/metrics.go
package metrics

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Defina suas métricas aqui
var (
	httpRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "beers_http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "handler", "status"},
	)
)

// Inicialize as métricas
func init() {
	prometheus.MustRegister(httpRequestCount)
}

// Middleware para registrar as métricas para cada solicitação HTTP

func PrometheusHandlerWrapper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		PrometheusHandler().ServeHTTP(w, r)
	}
}

func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}

func IncrementRequestCount(method string, path string, status int) {
	// Incrementa o contador de solicitações para /v1/beers
	httpRequestCount.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
}

// // Rota para expor as métricas para o Prometheus
// func RegisterMetricsHandler() {
// 	http.Handle("/metrics", promhttp.Handler())
// }
