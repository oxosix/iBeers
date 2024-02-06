// routes.go
package router

import (
	"github.com/d90ares/iBeers/internal/ibeers/http/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, beerHandler *handler.BeerHandler) {
	router.HandleFunc("/v1/beers", beerHandler.GetAllBeers).Methods("GET")
	// Adicione outras rotas conforme necessário
}
