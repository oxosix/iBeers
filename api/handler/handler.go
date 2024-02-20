// handler/beer_handler.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d90ares/iBeers/internal/app/usecase"
	"github.com/d90ares/iBeers/internal/domain"
	"github.com/d90ares/iBeers/pkg/errors"
	"github.com/d90ares/iBeers/pkg/logs"
	"github.com/d90ares/iBeers/pkg/metrics"
)

type BeerHandler struct {
	useCase usecase.UseCase
}

func NewBeerHandler(useCase usecase.UseCase) *BeerHandler {
	return &BeerHandler{
		useCase: useCase,
	}
}

func (h *BeerHandler) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	ctx := r.Context()
	beers, err := h.useCase.GetAllBeers(ctx)
	if err != nil {
		code := http.StatusInternalServerError
		metrics.IncrementRequestCount(r.Method, "/v1/beers", code)
		if httpErr, ok := err.(errors.HttpError); ok {
			errors.HandleError(w, httpErr)
			return
		}

		// Verifique se o erro é relacionado a "no content"
		if errors.IsNoContentError(err) {
			logs.Error("No Content", err)
			// Se for um erro relacionado a "no content", retorne um erro 204
			genericErr := errors.NewHttpError(http.StatusNoContent, "No content")
			errors.HandleError(w, genericErr)
			return
		}

		// Se o erro não puder ser identificado, trate como um erro interno
		genericErr := errors.NewHttpError(http.StatusInternalServerError, "Internal server error")
		errors.HandleError(w, genericErr)
		return
	}
	metrics.IncrementRequestCount(r.Method, "/v1/beers", code)
	// Se não houver erros, responda com JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(beers)
	// respondWithJSON(w, http.StatusOK, beers)
}

func (h *BeerHandler) Add(w http.ResponseWriter, r *http.Request) {
	// Obtenha o contexto da solicitação
	ctx := r.Context()

	// Decodifique o corpo da solicitação JSON para a estrutura Beer
	var beer domain.Beer
	if err := json.NewDecoder(r.Body).Decode(&beer); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Adicione a cerveja usando os dados fornecidos na solicitação
	addedBeer, err := h.useCase.AddBeer(ctx, &beer)
	if err != nil {
		if httpErr, ok := err.(errors.HttpError); ok {
			// Se o erro implementar a interface HttpError, manipule-o corretamente
			errors.HandleError(w, httpErr)
			return
		}
		// Se ocorrer um erro interno do servidor, retorne um erro 500
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Adicione o cabeçalho Location com a URL do recurso criado
	w.Header().Set("Location", fmt.Sprintf("/beers/%d", addedBeer.ID))

	// Retorne a cerveja adicionada como resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addedBeer)
}
