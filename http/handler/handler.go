// handler/beer_handler.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d90ares/iBeers/config/logs"
	"github.com/d90ares/iBeers/domain"
	"github.com/d90ares/iBeers/errors"
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
	ctx := r.Context()
	beers, err := h.useCase.GetAllBeers(ctx)
	if err != nil {
		// Verifique se o erro implementa a interface HTTPError
		if httpErr, ok := err.(errors.HttpError); ok {
			errors.HandleError(w, httpErr)
			return
		}
		if beers == nil {
			logs.Error("Não foram encontrados valores", err)
			// Se o erro não implementa a interface HTTPError, trate como um erro interno
			genericErr := errors.NewHttpError(http.StatusInternalServerError, "Erro interno do servidooooooooor")
			errors.HandleError(w, genericErr)
			return

		}
	}

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

// func formatJSONError(message string) []byte {
// 	appError := struct {
// 		Message string `json:"message"`
// 	}{
// 		message,
// 	}
// 	response, err := json.Marshal(appError)
// 	if err != nil {
// 		return []byte(err.Error())
// 	}
// 	return response
// }

// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	json.NewEncoder(w).Encode(payload)
// }

// Função auxiliar para lidar com erros HTTP
// func HandleHTTPError(w http.ResponseWriter, err error) {
// 	if e, ok := err.(errors.HTTPError); ok {
// 		// Se o erro implementa a interface, utilize suas propriedades
// 		respondWithError(w, e.StatusCode(), e.Error())
// 	} else {
// 		// Tratar outros tipos de erros
// 		respondWithError(w, http.StatusInternalServerError, "Erro interno do servidor")
// 	}
// }

// // Função auxiliar para responder com JSON
// func respondWithError(w http.ResponseWriter, code int, message string) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	response := map[string]string{"error": message}
// 	json.NewEncoder(w).Encode(response)
// }

// Implementar outros handlers, como GetBeerByID, SearchBeer, etc.
