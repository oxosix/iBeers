// internal/metrics/prometheus.go
package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Registrar métricas personalizadas aqui
var (
	beerCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "beer_count",
			Help: "Total number of beers",
		},
		[]string{},
	)
)

func init() {
	prometheus.MustRegister(beerCount)
}

// RegisterMetricsHandler adiciona um endpoint para as métricas do Prometheus.
func RegisterMetricsHandler() {
	http.Handle("/metrics", promhttp.Handler())
}

// IncrementBeerCount incrementa o contador de cervejas.
func IncrementBeerCount() {
	beerCount.WithLabelValues().Inc()
}
