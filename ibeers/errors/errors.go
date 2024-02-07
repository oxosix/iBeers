package errors

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	error
}

func HandleError(w http.ResponseWriter, err HttpError) {
	switch err.Code {
	case http.StatusOK:
		RespondWithError(w, http.StatusOK, err.Message, true) // Tratar código 200
	case http.StatusCreated:
		RespondWithError(w, http.StatusCreated, err.Message, true) // Tratar código 201
	case http.StatusNoContent:
		RespondWithError(w, http.StatusNoContent, err.Message, true) // Tratar código 204
	case http.StatusFound:
		RespondWithError(w, http.StatusFound, err.Message, true) // Tratar código 302
	case http.StatusBadRequest:
		RespondWithError(w, http.StatusBadRequest, err.Message, true) // Tratar código 400
	case http.StatusUnauthorized:
		RespondWithError(w, http.StatusUnauthorized, err.Message, true) // Tratar código 401
	case http.StatusForbidden:
		RespondWithError(w, http.StatusForbidden, err.Message, true) // Tratar código 403
	case http.StatusNotFound:
		RespondWithError(w, http.StatusNotFound, err.Message, true) // Tratar código 404
	default:
		RespondWithError(w, http.StatusInternalServerError, err.Message, true) // Tratar código 500
	}
	// Aqui você pode realizar ações específicas com base no código de status
	// RespondWithError(w, e.Status, e.Msg, true)

}

func RespondWithError(w http.ResponseWriter, statusCode int, message string, IncludeStatus bool) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	var Response map[string]interface{}
	if IncludeStatus {
		Response = map[string]interface{}{"error": message, "status": statusCode}
	} else {
		Response = map[string]interface{}{"error": message}
	}
	json.NewEncoder(w).Encode(Response)
}

func NewHttpError(statusCode int, message string) HttpError {
	return HttpError{
		Code:    statusCode,
		Message: message,
	}
}
