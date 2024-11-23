// routes.go
package router

import (
	"database/sql"
	"net/http"

	"github.com/d90ares/iBeers/api/handler"
	"github.com/d90ares/iBeers/pkg/health"
	"github.com/d90ares/iBeers/pkg/metrics"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, beerHandler *handler.BeerHandler) {
	router.HandleFunc("/v1/beers", beerHandler.GetAllBeers).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/beers", beerHandler.Add).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/beers/{id:[0-9]+}", beerHandler.GetByID).Methods("GET")
	// Adicione outras rotas conforme necessário
}

func SetupMetricsRoutes(router *mux.Router) {
	router.HandleFunc("/metrics", metrics.PrometheusHandlerWrapper()).Methods("GET")
}

func SetupHealthRoute(router *mux.Router, db *sql.DB) {
	hc := func(w http.ResponseWriter, r *http.Request) {
		health.HealthCheckHandler(w, r, db)
	}
	router.HandleFunc("/health", hc).Methods("GET")
}
