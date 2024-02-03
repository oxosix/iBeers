// handler/beer_handler.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d90ares/iBeers/internal/ibeers/domain"
)

type BeerHandler struct {
	useCase domain.UseCase
}

func NewBeerHandler(useCase domain.UseCase) *BeerHandler {
	return &BeerHandler{
		useCase: useCase,
	}
}

func (h *BeerHandler) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	beers, err := h.useCase.GetAllBeers()
	if err != nil {
		HandleHTTPError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, beers)
}

// Função auxiliar para lidar com erros HTTP
func HandleHTTPError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	errorMessage := map[string]interface{}{"error": err.Error()}
	json.NewEncoder(w).Encode(errorMessage)
}

// Função auxiliar para responder com JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// Implementar outros handlers, como GetBeerByID, SearchBeer, etc.
