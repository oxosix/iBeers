package errors

import (
	"encoding/json"
	"net/http"
)

type GenericError struct {
	Status int
	Msg    string
}

func (e *GenericError) StatusCode() int {
	return e.Status
}

func (e *GenericError) Error() string {
	return e.Msg
}

func handleError(w http.ResponseWriter, err error) {
	if e, ok := err.(*GenericError); ok {
		switch e.Status {
		case http.StatusOK:
			respondWithError(w, http.StatusOK, "Sucesso") // Tratar código 200
		case http.StatusCreated:
			// Tratar código 201
		case http.StatusNoContent:
			// Tratar código 204
		case http.StatusFound:
			// Tratar código 302
		case http.StatusBadRequest:
			// Tratar código 400
		case http.StatusUnauthorized:
			// Tratar código 401
		case http.StatusForbidden:
			// Tratar código 403
		case http.StatusNotFound:
			// Tratar código 404
		case http.StatusInternalServerError:
			// Tratar código 500
		default:
			// Tratar outros códigos de status
		}

		// Aqui você pode realizar ações específicas com base no código de status
		respondWithError(w, e.Status, e.Msg)
	} else {
		// Tratar outros tipos de erros
		respondWithError(w, http.StatusInternalServerError, "Erro interno do servidor")
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}
