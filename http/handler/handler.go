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
	respondWithJSON(w, http.StatusOK, beers)
}

func (h *BeerHandler) AddBeer(w http.ResponseWriter, r *http.Request) {
	// Decodifica o JSON da requisição para o objeto Beer
	var beer *domain.Beer
	if err := json.NewDecoder(r.Body).Decode(&beer); err != nil {
		http.Error(w, "Erro ao decodificar o JSON da cerveja", http.StatusBadRequest)
		return
	}

	// Adiciona a cerveja usando a camada de caso de uso
	ctx := r.Context()
	beer, err := h.useCase.AddBeer(ctx, beer) // Passando o contexto e o objeto de cerveja como argumentos
	if err != nil {
		// Trata o erro
		if httpErr, ok := err.(errors.HttpError); ok {
			errors.HandleError(w, httpErr)
			return
		}
		// Se o erro não implementa a interface HTTPError, trate como um erro interno
		logs.Error("Erro ao adicionar cerveja:", err)
		genericErr := errors.NewHttpError(http.StatusInternalServerError, "Erro interno do servidor")
		errors.HandleError(w, genericErr)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Cerveja adicionada com sucesso")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

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
