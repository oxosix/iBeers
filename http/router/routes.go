// routes.go
package router

import (
	"github.com/d90ares/iBeers/http/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, beerHandler *handler.BeerHandler) {
	router.HandleFunc("/v1/beers", beerHandler.GetAllBeers).Methods("GET")
	// router.HandleFunc("/v1/beers", beerHandler.AddBeer).Methods("POST")
	// Adicione outras rotas conforme necessário
}
