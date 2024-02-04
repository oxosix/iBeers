// routes.go
package router

import (
	"github.com/d90ares/iBeers/internal/ibeers/http/handler" // Substitua com o caminho real
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, beerHandler *handler.BeerHandler) {
	router.HandleFunc("/beers", beerHandler.GetAllBeers).Methods("GET")
	// Adicione outras rotas conforme necessário
}
